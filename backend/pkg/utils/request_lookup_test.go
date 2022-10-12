package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRequestLookUpQuery(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/?token=secret", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token, err := RequestLookUp(c, "token", "query")
	assert.NoError(t, err)
	assert.Equal(t, "secret", token)
}

func TestRequestLookUpParam(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("token")
	c.SetParamValues("secret")

	token, err := RequestLookUp(c, "token", "param")
	assert.NoError(t, err)
	assert.Equal(t, "secret", token)
}

func TestRequestLookUpCookie(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: "secret"})
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token, err := RequestLookUp(c, "token", "cookie")
	assert.NoError(t, err)
	assert.Equal(t, "secret", token)
}

func TestRequestLookUpHeader(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("token", "secret")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token, err := RequestLookUp(c, "token", "header")
	assert.NoError(t, err)
	assert.Equal(t, "secret", token)
}

func TestRequestLookUpForm(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Form = make(url.Values)
	req.Form.Set("token", "secret")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token, err := RequestLookUp(c, "token", "form")
	assert.NoError(t, err)
	assert.Equal(t, "secret", token)
}

func TestRequestLookUpNotFound(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_, err := RequestLookUp(c, "token", "query", "param", "form", "cookie", "header")
	assert.Error(t, err)
	assert.Equal(t, "not found", err.Error())
}

func TestRequestLookUpUndefined(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_, err := RequestLookUp(c, "token", "area")
	assert.Error(t, err)
	assert.Equal(t, "undefined field: area", err.Error())
}
