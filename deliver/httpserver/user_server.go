package httpserver

import (
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

	//authSrv := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)
	//mysqlRepo := mysql.New()
	//userSvc := userservice.New(authSrv, mysqlRepo)

	response, err := s.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, response)
}
