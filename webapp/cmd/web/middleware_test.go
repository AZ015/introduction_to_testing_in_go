package main

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_addIPToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{
			headerName:  "",
			headerValue: "",
			addr:        "",
			emptyAddr:   false,
		},
		{
			headerName:  "",
			headerValue: "",
			addr:        "",
			emptyAddr:   true,
		},
		{
			headerName:  "X-Forwarded-For",
			headerValue: "192.3.2.1",
			addr:        "",
			emptyAddr:   false,
		},
		{
			headerName:  "",
			headerValue: "",
			addr:        "hello:world",
			emptyAddr:   false,
		},
	}

	var app application

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(contextUserKey)
		require.NotNil(t, val, "not present")

		ip, ok := val.(string)
		require.True(t, ok, "not string")
		t.Log(ip)
	})

	for _, e := range tests {
		handlerToTest := app.addIPToContext(nextHandler)

		req := httptest.NewRequest(http.MethodGet, "http://testing", nil)
		if e.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {
	var app application

	ctx := context.Background()

	ctx = context.WithValue(ctx, contextUserKey, "whatever")

	ip := app.ipFromContext(ctx)

	require.Equal(t, ip, "whatever")
}
