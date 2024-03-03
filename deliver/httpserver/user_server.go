package httpserver

import (
	"github.com/abbasfisal/game-app/pkg/richerror"
	"github.com/abbasfisal/game-app/service/userservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) registerHandler(c echo.Context) error {

	var req userservice.RegisterRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	response, err := s.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, response)
}

func codeAndMessage(err error) (message string, code int) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)
		return re.Message(), MapKindToHttpStatusCode(re.Kind())
	default:
		return err.Error(), http.StatusBadRequest
	}
}

func MapKindToHttpStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindInvalid:
		return http.StatusUnprocessableEntity
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	default:
		//must be logged
		return http.StatusBadRequest
	}
}
