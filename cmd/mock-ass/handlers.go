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

func getContentTypeFromCache(sessionUuid, fallback string) string {
	if contentTypeRaw, found := LocalCache.Get(getCacheCtKey(sessionUuid)); found {
		return contentTypeRaw.(string)
	}
	return fallback
}

func responseFromCache(w http.ResponseWriter, hash, result, sessionUuid string) int {
	LocalCache.Set(getCacheHashKey(hash), result, defaultDataTtlMinutes)

	contentType := getContentTypeFromCache(sessionUuid, defaultContentType)
	LocalCache.Set(getCacheCtKey(sessionUuid), contentType, defaultDataTtlMinutes)

	w.Header().Set("Content-Type", contentType)
	setCorsHeaders(w)
	io.WriteString(w, result)
	return http.StatusOK
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
			return responseFromCache(w, hash, result.(string), sessionUuid)
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
			setCorsHeaders(w)
			w.WriteHeader(http.StatusTemporaryRedirect)
			return http.StatusTemporaryRedirect
		}

		contentType := getContentTypeFromCache(sessionUuid, defaultContentType)
		LocalCache.Set(getCacheCtKey(sessionUuid), contentType, defaultDataTtlMinutes)

		w.Header().Set("Content-Type", contentType)
		setCorsHeaders(w)
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
	setCorsHeaders(w)
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
		setCorsHeaders(w)
		io.WriteString(w, "")
		return http.StatusOK
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return http.StatusMethodNotAllowed
	}
}

func parseTtlMin(r *http.Request, fallback time.Duration) (ttl time.Duration, err error) {
	ttl = fallback

	var ttlRaw string
	if ttlRaw = r.FormValue(formKeySessionTtlMin); ttlRaw == "" {
		return
	}

	var ttlParsedInt int64
	ttlParsedInt, err = strconv.ParseInt(ttlRaw, 10, 64)
	if err != nil {
		return
	}

	ttl = time.Duration(ttlParsedInt) * time.Minute
	return
}

func get_session(w http.ResponseWriter, r *http.Request, _ *random_data.RandomDataCollection) int {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return http.StatusMethodNotAllowed
	}

	r.ParseForm()

	contentType := defaultContentType
	if contentTypeRaw := r.FormValue(formKeyContentType); contentTypeRaw != "" {
		contentType = contentTypeRaw
	}

	ttl, err := parseTtlMin(r, defaultSessionTtlMinutes)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return http.StatusInternalServerError
	}

	userTpl := r.FormValue(formKeyTemplate)
	sessionUuid := getHash()

	LocalCache.Set(sessionUuid, userTpl, ttl)
	LocalCache.Set(getCacheCtKey(sessionUuid), contentType, ttl)

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

func setCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "X-Jquery-Json, Content-Type, Accept, Content-Length, Origin")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
