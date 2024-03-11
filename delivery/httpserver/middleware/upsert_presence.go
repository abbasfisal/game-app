package middleware

import (
	"github.com/abbasfisal/game-app/delivery/dto"
	"github.com/abbasfisal/game-app/pkg/errmsg"
	"github.com/abbasfisal/game-app/pkg/timestamp"
	"github.com/abbasfisal/game-app/service/presenceservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpsertPresence(service presenceservice.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, err := service.Upsert(c.Request().Context(), dto.UpsertPresenceRequest{
				UserID:    1,
				Timestamp: timestamp.Now(),
			})
			if err != nil {
				//we can just log the internal error and just continue
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": errmsg.ErrorMsgSomethingWentWrong,
				})
			}
			return nil
		}
	}
}
