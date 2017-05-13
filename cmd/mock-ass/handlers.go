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

const SESSION_URL string = "/session/?s=%s"
const DEFAULT_CONTENT_TYPE string = "application/json"

const (
	DEFAULT_SESSION_TTL_MINUTES int64 = 60
	DEFAULT_DATA_TTL_MINUTES    int64 = 15
)
const (
	FORM_KEY_TEMPLATE        = "template"
	FORM_KEY_CONTENT_TYPE    = "content_type"
	FORM_KEY_SESSION_TTL_MIN = "session_ttl_min"
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

func hello(w http.ResponseWriter, r *http.Request, collection *random_data.RandomDataCollection) int {
	r.ParseForm()
	if r.Method == http.MethodGet {
		need_redirect := false
		hash := r.FormValue("h")
		session_uuid := r.FormValue("s")
		if session_uuid == "" {
			w.WriteHeader(http.StatusBadRequest)
			return http.StatusBadRequest
		}
		data_ttl_min := DEFAULT_DATA_TTL_MINUTES
		cache_data_ttl_td := time.Duration(data_ttl_min * int64(time.Minute))
		if hash == "" {
			hash = getHash()
			need_redirect = true
		} else {
			if result, found := LocalCache.Get(getCacheHashKey(hash)); found {
				LocalCache.Set(getCacheHashKey(hash), result, cache_data_ttl_td)
				content_type := DEFAULT_CONTENT_TYPE
				content_type_raw, found := LocalCache.Get(getCacheCtKey(session_uuid))
				if found {
					content_type = content_type_raw.(string)
				}
				LocalCache.Set(getCacheCtKey(session_uuid), content_type, cache_data_ttl_td)
				w.Header().Set("Content-Type", content_type)
				w.Header().Set("Access-Control-Allow-Headers", "X-Jquery-Json, Content-Type, Accept, Content-Length, Origin")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				io.WriteString(w, result.(string))
				return http.StatusOK
			}
		}

		if user_tpl_c, found := LocalCache.Get(session_uuid); found {
			user_tpl := user_tpl_c.(string)
			out, err := gen.GenerateByTemplate(user_tpl, hash, collection)
			if err != nil {
				log.Println(color.RedString(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return http.StatusInternalServerError
			}
			LocalCache.Set(getCacheHashKey(hash), out, cache_data_ttl_td)

			if need_redirect {
				url_redirect := r.URL
				q := url_redirect.Query()
				q.Set("s", session_uuid)
				q.Set("h", hash)
				url_redirect.RawQuery = q.Encode()
				w.Header().Set("Location", url_redirect.String())
				w.Header().Set("Access-Control-Allow-Headers", "X-Jquery-Json, Content-Type, Accept, Content-Length, Origin")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(http.StatusTemporaryRedirect)
				return http.StatusTemporaryRedirect
			}
			content_type, found := LocalCache.Get(getCacheCtKey(session_uuid))
			if !found {
				content_type = DEFAULT_CONTENT_TYPE
			}
			LocalCache.Set(getCacheCtKey(session_uuid), content_type, cache_data_ttl_td)
			w.Header().Set("Content-Type", content_type.(string))
			w.Header().Set("Access-Control-Allow-Headers", "X-Jquery-Json, Content-Type, Accept, Content-Length, Origin")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			io.WriteString(w, out)
			return http.StatusOK
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return http.StatusUnauthorized
		}
	} else if r.Method == http.MethodPost {
		user_tpl := r.FormValue(FORM_KEY_TEMPLATE)

		content_type := DEFAULT_CONTENT_TYPE
		if content_type_raw := r.FormValue(FORM_KEY_CONTENT_TYPE); content_type_raw != "" {
			content_type = content_type_raw
		}
		hash := getHash()
		out, err := gen.GenerateByTemplate(user_tpl, hash, collection)
		if err != nil {
			log.Println(color.RedString(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return http.StatusInternalServerError
		}
		w.Header().Set("Content-Type", content_type)
		w.Header().Set("Access-Control-Allow-Headers", "X-Jquery-Json, Content-Type, Accept, Content-Length, Origin")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		io.WriteString(w, out)
		return http.StatusOK
	} else if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Headers", "X-Jquery-Json, Content-Type, Accept, Content-Length, Origin")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		io.WriteString(w, "")
		return http.StatusOK
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return http.StatusMethodNotAllowed
}

func get_session(w http.ResponseWriter, r *http.Request, _ *random_data.RandomDataCollection) int {
	if r.Method == http.MethodPost {
		r.ParseForm()
		user_tpl := r.FormValue(FORM_KEY_TEMPLATE)

		content_type := DEFAULT_CONTENT_TYPE
		if content_type_raw := r.FormValue(FORM_KEY_CONTENT_TYPE); content_type_raw != "" {
			content_type = content_type_raw
		}

		ttl := DEFAULT_SESSION_TTL_MINUTES
		if ttl_raw := r.FormValue(FORM_KEY_SESSION_TTL_MIN); ttl_raw != "" {
			var err error
			ttl, err = strconv.ParseInt(ttl_raw, 10, 32)
			if err != nil {
				log.Println(color.RedString(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return http.StatusInternalServerError
			}
		}

		session_uuid, err := newUUID()
		for err != nil {
			session_uuid, err = newUUID()
		}

		cache_data_ttl_td := time.Minute * time.Duration(ttl)
		LocalCache.Set(session_uuid, user_tpl, cache_data_ttl_td)
		LocalCache.Set(getCacheCtKey(session_uuid), content_type, cache_data_ttl_td)

		session_resp := &SessionResponse{
			Session: session_uuid,
			Url:     fmt.Sprintf(SESSION_URL, session_uuid),
		}

		session_resp_json, err := json.Marshal(session_resp)
		if err != nil {
			log.Println(color.RedString(err.Error()))
			return http.StatusInternalServerError
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(session_resp_json)
		return http.StatusOK
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	return http.StatusMethodNotAllowed
}
