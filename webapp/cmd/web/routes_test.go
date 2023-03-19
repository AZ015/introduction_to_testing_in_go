package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"net/http"
	"strings"
	"testing"
)

func Test_application_routes(t *testing.T) {
	var registered = []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/login", "POST"},
		{"/static/*", "GET"},
	}

	var app application
	mux := app.routes()

	chiRoutes := mux.(chi.Routes)

	for _, route := range registered {
		require.True(t, routeExists(route.route, route.method, chiRoutes))
	}
}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}

		return nil
	})

	return found
}
