package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/wolfmetr/mock-ass/gen"
	"github.com/wolfmetr/mock-ass/random_data"

	"github.com/fatih/color"
	"github.com/pmylund/go-cache"
)

const sessionUrl string = "/session/?s=%s"
const defaultContentType string = "application/json"

const (
	defaultSessionTtlMinutes time.Duration = 60 * time.Minute
	defaultDataTtlMinutes    time.Duration = 15 * time.Minute
)
const (
	formKeyTemplate      = "template"
	formKeyContentType   = "content_type"
	formKeySessionTtlMin = "session_ttl_min"
)

type SessionResponse struct {
	Session string `json:"session"`
	Url     string `json:"url"`
}

type ErrorResponse struct {
	ErrorMsg string `json:"error_message"`
}

var LocalCache *cache.Cache

func init() {
	LocalCache = cache.New(60*time.Minute, 30*time.Second)
}

func getHash() string {
	hash, err := newUUID()
	for err != nil {
		hash, err = newUUID()
	}
	return hash
}

func getCacheHashKey(hash string) string {
	return fmt.Sprintf("hash_%s", hash)
}
func getCacheCtKey(sessionUuid string) string {
	return fmt.Sprintf("ct_%s", sessionUuid)
}

func generateRespGetMethod(w http.ResponseWriter, r *http.Request, collection *random_data.RandomDataCollection) int {
	needRedirect := false
	hash := r.FormValue("h")
	sessionUuid := r.FormValue("s")
	if sessionUuid == "" {
		log.Println(color.RedString("session argument empty or not found"))
		w.WriteHeader(http.StatusBadRequest)
		return http.StatusBadRequest
	}

	if hash == "" {
		hash = getHash()
		needRedirect = true
	} else {
		if result, found := LocalCache.Get(getCacheHashKey(hash)); found {
			LocalCache.Set(getCacheHashKey(hash), result, defaultDataTtlMinutes)
			contentType := defaultContentType
			if contentTypeRaw, found := LocalCache.Get(getCacheCtKey(sessionUuid)); found {
				contentType = contentTypeRaw.(string)
			}
			LocalCache.Set(getCacheCtKey(sessionUuid), contentType, defaultDataTtlMinutes)
			w.Header().Set("Content-Type", contentType)
			w.Header().Set("Access-Control-Allow-Headers", "X-Jquery-Json, Content-Type, Accept, Content-Length, Origin")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			io.WriteString(w, result.(string))
			return http.StatusOK
		}
	}

	if userTplC, found := LocalCache.Get(sessionUuid); found {
		userTpl := userTplC.(string)
		out, err := gen.GenerateByTemplate(userTpl, hash, collection)
		if err != nil {
			log.Println(color.RedString(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return http.StatusInternalServerError
		}
		LocalCache.Set(getCacheHashKey(hash), out, defaultDataTtlMinutes)

		if needRedirect {
			urlRedirect := r.URL
			q := urlRedirect.Query()
			q.Set("s", sessionUuid)
			q.Set("h", hash)
			urlRedirect.RawQuery = q.Encode()
			w.Header().Set("Location", urlRedirect.String())
			w.Header().Set("Access-Control-Allow-Headers", "X-Jquery-Json, Content-Type, Accept, Content-Length, Origin")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusTemporaryRedirect)
			return http.StatusTemporaryRedirect
		}
		contentType, found := LocalCache.Get(getCacheCtKey(sessionUuid))
		if !found {
			contentType = defaultContentType
		}
		LocalCache.Set(getCacheCtKey(sessionUuid), contentType, defaultDataTtlMinutes)
		w.Header().Set("Content-Type", contentType.(string))
		w.Header().Set("Access-Control-Allow-Headers", "X-Jquery-Json, Content-Type, Accept, Content-Length, Origin")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		io.WriteString(w, out)
		return http.StatusOK
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return http.StatusUnauthorized
	}
}

func generateRespPostMethod(w http.ResponseWriter, r *http.Request, collection *random_data.RandomDataCollection) int {
	userTpl := r.FormValue(formKeyTemplate)

	contentType := defaultContentType
	if contentTypeRaw := r.FormValue(formKeyContentType); contentTypeRaw != "" {
		contentType = contentTypeRaw
	}
	hash := getHash()
	out, err := gen.GenerateByTemplate(userTpl, hash, collection)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return http.StatusInternalServerError
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Access-Control-Allow-Headers", "X-Jquery-Json, Content-Type, Accept, Content-Length, Origin")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, out)
	return http.StatusOK
}

func generateResp(w http.ResponseWriter, r *http.Request, collection *random_data.RandomDataCollection) int {
	switch r.Method {
	case http.MethodGet:
		r.ParseForm()
		return generateRespGetMethod(w, r, collection)
	case http.MethodPost:
		r.ParseForm()
		return generateRespPostMethod(w, r, collection)
	case http.MethodOptions:
		w.Header().Set("Access-Control-Allow-Headers", "X-Jquery-Json, Content-Type, Accept, Content-Length, Origin")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		io.WriteString(w, "")
		return http.StatusOK
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return http.StatusMethodNotAllowed
	}
}

func get_session(w http.ResponseWriter, r *http.Request, _ *random_data.RandomDataCollection) int {
	if r.Method == http.MethodPost {
		r.ParseForm()
		userTpl := r.FormValue(formKeyTemplate)

		contentType := defaultContentType
		if contentTypeRaw := r.FormValue(formKeyContentType); contentTypeRaw != "" {
			contentType = contentTypeRaw
		}

		ttl := defaultSessionTtlMinutes
		if ttlRaw := r.FormValue(formKeySessionTtlMin); ttlRaw != "" {
			var err error
			var ttlParsedInt int64
			ttlParsedInt, err = strconv.ParseInt(ttlRaw, 10, 32)
			if err != nil {
				log.Println(color.RedString(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return http.StatusInternalServerError
			}
			ttl = time.Duration(ttlParsedInt) * time.Minute
		}

		sessionUuid := getHash()
		cacheDataTtlTd := time.Minute * time.Duration(ttl)
		LocalCache.Set(sessionUuid, userTpl, cacheDataTtlTd)
		LocalCache.Set(getCacheCtKey(sessionUuid), contentType, cacheDataTtlTd)

		sessionResp := &SessionResponse{
			Session: sessionUuid,
			Url:     fmt.Sprintf(sessionUrl, sessionUuid),
		}

		sessionRespJson, err := json.Marshal(sessionResp)
		if err != nil {
			log.Println(color.RedString(err.Error()))
			return http.StatusInternalServerError
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(sessionRespJson)
		return http.StatusOK
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	return http.StatusMethodNotAllowed
}
