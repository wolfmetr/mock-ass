package main

import (
	"net/http"
	"strconv"
	"time"
)

func parseTtlMin(r *http.Request) (time.Duration, error) {
	var err error
	var ttlRaw string
	if ttlRaw = r.FormValue(formKeySessionTtlMin); ttlRaw == "" {
		return defaultSessionTtlMinutes, nil
	}

	var ttlParsedInt int64
	ttlParsedInt, err = strconv.ParseInt(ttlRaw, 10, 64)
	if err != nil {
		return 0, err
	}

	ttl := time.Duration(ttlParsedInt) * time.Minute
	return ttl, err
}

func parseContentType(r *http.Request) string {
	if contentTypeRaw := r.FormValue(formKeyContentType); contentTypeRaw != "" {
		return contentTypeRaw
	}
	return defaultContentType

}
