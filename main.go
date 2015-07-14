package main

import (
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/wolfmetr/mock-ass/generator"
	"github.com/wolfmetr/mock-ass/random_data"

	"github.com/fatih/color"
	"github.com/golang/glog"
	"github.com/pmylund/go-cache"
)

type myHandler struct{}

var mux map[string]func(http.ResponseWriter, *http.Request) int

const SESSION_URL string = "/session/?s=%s"

const DEFAULT_SESSION_TTL_MINUTES int64 = 60
const DEFAULT_DATA_TTL_MINUTES int64 = 15

const DEFAULT_CONTENT_TYPE string = "application/json"
const FORM_KEY_TEMPLATE = "template"
const FORM_KEY_CONTENT_TYPE = "content_type"
const FORM_KEY_SESSION_TTL_MIN = "session_ttl_min"

var LocalCache *cache.Cache

func init() {
	LocalCache = cache.New(60*time.Minute, 30*time.Second)
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if hand, ok := mux[r.URL.Path]; ok {
		status_code := hand(w, r)
		if status_code >= 200 && status_code < 300 {
			glog.Info(color.GreenString("[%s] %s — %d", r.Method, r.URL.String(), status_code))
		} else {
			glog.Info(color.RedString("[%s] %s — %d", r.Method, r.URL.String(), status_code))
		}
		return
	} else {
		glog.Info(color.RedString("[%s] %s — %d", r.Method, r.URL.String(), 404))
		io.WriteString(w, "404 not found mthrfckr!")
	}

}

func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

type SessionResponse struct {
	Session string `json:"session"`
	Url     string `json:"url"`
}

type ErrorResponse struct {
	ErrorMsg string `json:"error_message"`
}

func GetHash() string {
	hash, err := newUUID()
	for err != nil {
		hash, err = newUUID()
	}
	return hash
}

func hello(w http.ResponseWriter, r *http.Request) int {
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
			hash = GetHash()
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
			out, err := generator.GenerateByTemplate(user_tpl, hash)
			if err != nil {
				glog.Error(color.RedString(err.Error()))
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
		hash := GetHash()
		out, err := generator.GenerateByTemplate(user_tpl, hash)
		if err != nil {
			glog.Error(color.RedString(err.Error()))
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

func get_session(w http.ResponseWriter, r *http.Request) int {
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
				glog.Error(color.RedString(err.Error()))
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

func main() {
	flagNoColor := flag.Bool("no-color", false, "Disable color output")
	port := flag.Int("port", 8000, "Server start port")
	flag.Lookup("logtostderr").Value.Set("true")
	flag.Parse()
	if *flagNoColor {
		color.NoColor = true
	}

	random_data.InitWithDefaults()

	port_str := strconv.FormatInt(int64(*port), 10)
	server := http.Server{
		Addr:    ":" + port_str,
		Handler: &myHandler{},
	}
	glog.Info(color.BlueString("Start server port %s", port_str))
	mux = make(map[string]func(http.ResponseWriter, *http.Request) int)
	mux["/session/"] = hello
	mux["/init"] = get_session
	server.ListenAndServe()
}
