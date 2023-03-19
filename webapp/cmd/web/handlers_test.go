package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var theTests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/fish", http.StatusNotFound},
	}

	var app application
	routes := app.routes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	pathToTemplates = "./../../templates/"

	for _, e := range theTests {
		res, err := ts.Client().Get(fmt.Sprintf("%s%s", ts.URL, e.url))
		if err != nil {
			t.Log(e)
			t.Fatal(err)
		}

		require.Equal(t, e.expectedStatusCode, res.StatusCode)
	}
}
