package main

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	form := NewForm(nil)

	has := form.Has("whatever")
	require.False(t, has, "form shows has field when it should not")

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = NewForm(postedData)

	has = form.Has("a")
	require.True(t, has)
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/whatever", nil)
	form := NewForm(r.PostForm)
	form.Required("a", "b", "c")
	require.False(t, form.Valid())

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")
	r, _ = http.NewRequest(http.MethodPost, "/whatever", nil)
	r.PostForm = postedData

	form = NewForm(r.PostForm)
	form.Required("a", "b", "c")
	require.True(t, form.Valid())
}

func TestForm_Check(t *testing.T) {
	form := NewForm(nil)

	form.Check(false, "password", "password is required")
	require.False(t, form.Valid())
}

func TestForm_ErrorGet(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")
	s := form.Errors.Get("password")
	require.True(t, len(s) > 0)

	s = form.Errors.Get("whatever")
	require.Equal(t, 0, len(s))
}
