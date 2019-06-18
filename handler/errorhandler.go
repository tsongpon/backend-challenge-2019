package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/tsongpon/backend-challenge-2019/bserror"
	"gopkg.in/go-playground/validator.v9"
)

type errorTransport struct {
	Message string `json:"message"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	log.Debug("handlering error")
	code := http.StatusInternalServerError
	et := errorTransport{}
	switch e := err.(type) {
	case *bserror.NotFoundError:
		code = http.StatusNotFound
		et.Message = e.Error()
		break
	case *bserror.InsufficientStockError:
		code = http.StatusBadRequest
		et.Message = e.Error()
		break
	case *bserror.BadParameterError:
		code = http.StatusBadRequest
		et.Message = e.Error()
		break
	case *bserror.DataVersionError:
		code = http.StatusConflict
		et.Message = e.Error()
		break
	case validator.ValidationErrors:
		code = http.StatusBadRequest
		et.Message = e.Error()
		break
	case *echo.HTTPError:
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			et.Message = http.StatusText(code)
		}
		break
	}
	c.Logger().Error(err)
	c.JSON(code, et)
}
