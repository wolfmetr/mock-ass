package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/wolfmetr/mock-ass/random_data"

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

func newSessionResponse(sessionUuid string) *SessionResponse {
	return &SessionResponse{
		Session: sessionUuid,
		Url:     fmt.Sprintf(sessionUrl, sessionUuid),
	}
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

func responseRedirect(w http.ResponseWriter, r *http.Request, sessionUuid, hash string) int {
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

func generateRespGetMethod(w http.ResponseWriter, r *http.Request, collection *random_data.RandomDataCollection) int {
	hash := r.FormValue("h")
	sessionUuid := r.FormValue("s")
	if sessionUuid == "" {
		log.Println("session argument empty or not found")
		w.WriteHeader(http.StatusBadRequest)
		return http.StatusBadRequest
	}

	var found bool
	var userTplC interface{}
	if userTplC, found = LocalCache.Get(sessionUuid); !found {
		w.WriteHeader(http.StatusBadRequest)
		return http.StatusUnauthorized
	}
	if hash != "" {
		if result, found := LocalCache.Get(getCacheHashKey(hash)); found {
			return responseFromCache(w, hash, result.(string), sessionUuid)
		}
		// TODO: invalid hash response?
		w.WriteHeader(http.StatusBadRequest)
		return http.StatusUnauthorized

	}

	// generate resp from template
	hash = getHash()

	userTpl := userTplC.(string)
	out, err := random_data.Render(userTpl, hash, collection)
	if err != nil {
		return respInternalServerError(w, err)
	}
	// set resp to cache
	LocalCache.Set(getCacheHashKey(hash), out, defaultDataTtlMinutes)

	// and redirect to stable url
	return responseRedirect(w, r, sessionUuid, hash)
}

func generateRespPostMethod(w http.ResponseWriter, r *http.Request, collection *random_data.RandomDataCollection) int {
	userTpl := r.FormValue(formKeyTemplate)
	contentType := parseContentType(r)

	hash := getHash()
	out, err := random_data.Render(userTpl, hash, collection)
	if err != nil {
		return respInternalServerError(w, err)
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

func initSession(w http.ResponseWriter, r *http.Request, _ *random_data.RandomDataCollection) int {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return http.StatusMethodNotAllowed
	}

	r.ParseForm()

	contentType := parseContentType(r)
	userTpl := r.FormValue(formKeyTemplate)
	ttl, err := parseTtlMin(r)
	if err != nil {
		return respInternalServerError(w, err)
	}
	sessionUuid := getHash()

	LocalCache.Set(sessionUuid, userTpl, ttl)
	LocalCache.Set(getCacheCtKey(sessionUuid), contentType, ttl)

	sessionResp := newSessionResponse(sessionUuid)
	sessionRespJson, err := json.Marshal(sessionResp)
	if err != nil {
		return respInternalServerError(w, err)
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

func respInternalServerError(w http.ResponseWriter, err error) int {
	if err != nil {
		log.Println(err.Error())
	}
	w.WriteHeader(http.StatusInternalServerError)
	return http.StatusInternalServerError
}
