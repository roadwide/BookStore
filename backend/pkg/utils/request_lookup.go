package utils

import (
	"errors"
	"github.com/labstack/echo/v4"
)

func RequestLookUp(c echo.Context, key string, fields ...string) (string, error) {
	s := ""
	for _, field := range fields {
		switch field {
		case "query":
			s = c.QueryParam(key)
		case "form":
			s = c.FormValue(key)
		case "param":
			s = c.Param(key)
		case "cookie":
			cookie, err := c.Cookie(key)
			if err == nil {
				s = cookie.Value
			}
		case "header":
			s = c.Request().Header.Get(key)
		default:
			return "", errors.New("undefined field: " + field)
		}
		if s != "" {
			return s, nil
		}
	}
	return "", errors.New("not found")
}
