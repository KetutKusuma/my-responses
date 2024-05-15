package myresponses

import (
	"net/http"
	"strings"

	"github.com/ketutkusuma/my-responses/responsegraph"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// return error bad request 400 with many message
// but will return 404 if record not found as message
func HandleBadrequest(c echo.Context, message ...string) error {
	status := http.StatusBadRequest
	if message[0] == "record not found" {
		status = 404
	}
	res := responsegraph.ResponseGeneric{
		Code:    status,
		Message: strings.Join(message, " "),
	}
	ErrorShowErrorLogrus(c, message...)
	return c.JSON(status, res)
}

// return error with status custom
// but will return 404 if record not found as message
func HandleError(c echo.Context, status int, message ...string) error {
	if len(message) > 1 {
		if message[1] == "record not found" {
			status = 404
		}
	}
	if message[0] == "record not found" {
		status = 404
	}
	res := responsegraph.ResponseGeneric{
		Code:    status,
		Message: strings.Join(message, " "),
	}
	ErrorShowErrorLogrus(c, message...)
	return c.JSON(status, res)
}

// this handle error not found use this like whatNotFound : "user"
// will return 404 and message : Not found user
func HandleErrorNotFound(c echo.Context, whatNotfound string) error {
	status := http.StatusNotFound
	res := responsegraph.ResponseGeneric{
		Code:    status,
		Message: "Not found " + whatNotfound,
	}
	return c.JSON(status, res)
}

// this will handle error internal 500
// but if message error is record not found will return 404
func HandleErrorInternal(c echo.Context, message ...string) error {
	status := http.StatusInternalServerError
	if len(message) > 1 {
		if message[1] == "record not found" {
			status = 404
		}
	}
	if message[0] == "record not found" {
		status = 404
	}
	res := responsegraph.ResponseGeneric{
		Code:    status,
		Message: strings.Join(message, " "),
	}
	ErrorShowErrorLogrus(c, message...)

	return c.JSON(status, res)
}

func ErrorShowErrorLogrus(c echo.Context, message ...string) {
	l := logrus.Logger{}

	l.WithFields(
		logrus.Fields{
			"Method":     c.Request().Method,
			"Remote IP":  c.Request().RemoteAddr,
			"Uri Path":   c.Request().URL.Path,
			"Route Path": c.Path(),
			"Request ID": c.Response().Header().Get(echo.HeaderXRequestID),
			"User-Agent": c.Request().UserAgent(),
			"Error":      message,
			"Real IP":    c.RealIP(),
		},
	).Errorln("Request error : ")

}
