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

	"github.com/fatih/color"
	"github.com/pmylund/go-cache"
	"github.com/wolfmetr/mock-ass/random_data"
)

const SESSION_URL string = "/session/?s=%s"

const DEFAULT_SESSION_TTL_MINUTES int64 = 60
const DEFAULT_DATA_TTL_MINUTES int64 = 15

const DEFAULT_CONTENT_TYPE string = "application/json"
const FORM_KEY_TEMPLATE = "template"
const FORM_KEY_CONTENT_TYPE = "content_type"
const FORM_KEY_SESSION_TTL_MIN = "session_ttl_min"

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

func hello(w http.ResponseWriter, r *http.Request, collection *random_data.RandomDataCollection) int {
	r.ParseForm()
	if r.Method == "GET" {
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
			if result, found := LocalCache.Get("hash_" + hash); found {
				LocalCache.Set("hash_"+hash, result, cache_data_ttl_td)
				content_type, found := LocalCache.Get("ct_" + session_uuid)
				if !found {
					content_type = DEFAULT_CONTENT_TYPE
				}
				LocalCache.Set("ct_"+session_uuid, content_type, cache_data_ttl_td)
				w.Header().Set("Content-Type", content_type.(string))
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
			LocalCache.Set("hash_"+hash, out, cache_data_ttl_td)

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
			content_type, found := LocalCache.Get("ct_" + session_uuid)
			if !found {
				content_type = DEFAULT_CONTENT_TYPE
			}
			LocalCache.Set("ct_"+session_uuid, content_type, cache_data_ttl_td)
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
	} else if r.Method == "POST" {
		user_tpl := r.FormValue(FORM_KEY_TEMPLATE)
		content_type := DEFAULT_CONTENT_TYPE
		content_type_raw := r.FormValue(FORM_KEY_CONTENT_TYPE)
		if content_type_raw != "" {
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
	} else if r.Method == "OPTIONS" {
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
	if r.Method == "POST" {
		r.ParseForm()
		user_tpl := r.FormValue(FORM_KEY_TEMPLATE)
		content_type := DEFAULT_CONTENT_TYPE
		content_type_raw := r.FormValue(FORM_KEY_CONTENT_TYPE)
		if content_type_raw != "" {
			content_type = content_type_raw
		}
		ttl := DEFAULT_SESSION_TTL_MINUTES
		ttl_raw := r.FormValue(FORM_KEY_SESSION_TTL_MIN)
		if ttl_raw != "" {
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
		session_resp := &SessionResponse{
			Session: session_uuid,
			Url:     fmt.Sprintf(SESSION_URL, session_uuid),
		}
		cache_data_ttl_td := time.Duration(ttl * int64(time.Minute))
		LocalCache.Set(session_uuid, user_tpl, cache_data_ttl_td)
		LocalCache.Set("ct_"+session_uuid, content_type, cache_data_ttl_td)

		session_resp_json, _ := json.Marshal(session_resp)
		w.Header().Set("Content-Type", "application/json")
		w.Write(session_resp_json)
		return http.StatusOK
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	return http.StatusMethodNotAllowed
}
