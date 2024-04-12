package myresponses

import (
	"log"
	"net/http"
	"strings"

	"github.com/ketutkusuma/my-responses/responsegraph"
	"github.com/labstack/echo/v4"
)

func HandleSuccess(c echo.Context, data interface{}) error {
	res := responsegraph.ResponseGeneric{
		Code:    200,
		Message: "Success",
		Data:    data,
	}
	return c.JSON(http.StatusOK, res)
}
func HandleSuccessMsg(c echo.Context, msg ...string) error {
	status := 201
	res := responsegraph.ResponseGeneric{
		Code:    status,
		Message: strings.Join(msg, " "),
		Data:    nil,
	}

	return c.JSON(status, res)
}
func HandleSuccessCreate(c echo.Context, data interface{}) error {
	res := responsegraph.ResponseGeneric{
		Code:    201,
		Message: "Success",
		Data:    data,
	}
	return c.JSON(http.StatusCreated, res)
}
func HandleSuccessCustom(c echo.Context, data interface{}, status int, message *string) error {
	log.Println("message : ", &message)
	if message != nil {
		res := responsegraph.ResponseGeneric{
			Code:    status,
			Message: "Success " + *message,
			Data:    data,
		}
		return c.JSON(status, res)
	}
	res := responsegraph.ResponseGeneric{
		Code:    status,
		Message: "Success",
		Data:    data,
	}
	return c.JSON(status, res)
}

func HandleBadrequest(c echo.Context, message ...string) error {
	status := http.StatusBadRequest
	if message[0] == "record not found" {
		status = 404
	}
	res := responsegraph.ResponseGeneric{
		Code:    status,
		Message: strings.Join(message, " "),
	}
	return c.JSON(status, res)
}

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
	return c.JSON(status, res)
}

func HandleErrorNotFound(c echo.Context, whatNotfound string) error {
	status := http.StatusNotFound
	res := responsegraph.ResponseGeneric{
		Code:    status,
		Message: "Not found " + whatNotfound,
	}
	return c.JSON(status, res)
}

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
	return c.JSON(status, res)
}

func HandleSuccessPaginate(c echo.Context, data interface{}, paginate responsegraph.Paginate) error {
	status := 200
	res := responsegraph.ResponseGenericPaginate{
		Code:     status,
		Message:  "Success",
		Data:     data,
		Paginate: paginate,
	}
	return c.JSON(status, res)
}
