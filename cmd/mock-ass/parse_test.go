package main

import (
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestParseTtlMin(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/?session_ttl_min=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	ttl, err := parseTtlMin(req)
	if err != nil {
		t.Errorf("expected err is nil, but %+v", err)
	}

	if ttl != 10*time.Minute {
		t.Errorf("ttl expected %v; actual %v", 10*time.Minute, ttl)
	}
}

func TestParseTtlMinDefault(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/?nothing=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	ttl, err := parseTtlMin(req)
	if err != nil {
		t.Errorf("expected err is nil, but %+v", err)
	}

	if ttl != defaultSessionTtlMinutes {
		t.Errorf("ttl expected %v (defaultSessionTtlMinutes); actual %v", defaultSessionTtlMinutes, ttl)
	}
}

func TestParseTtlMinError(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/?session_ttl_min=invalid_ttl", nil)
	if err != nil {
		t.Fatal(err)
	}

	ttl, err := parseTtlMin(req)
	if err == nil {
		t.Error("expected err is parse error, but is nil")
	} else if err.(*strconv.NumError).Err != strconv.ErrSyntax {
		t.Errorf("expected err is strconv.ErrSyntax, but %+v", err)
	}

	if ttl != 0 {
		t.Errorf("ttl expected %v; actual %v", 0, ttl)
	}
}
