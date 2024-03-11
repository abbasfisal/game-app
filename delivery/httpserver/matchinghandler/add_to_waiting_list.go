package matchinghandler

import (
	"github.com/abbasfisal/game-app/delivery/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) addToWaitingList(c echo.Context) error {

	//baind request
	var req dto.AddToWaitingListRequest
	res, err := h.matchingSvc.AddToWaitingList(req)

	req.UserID = 3

	if err != nil {
		return echo.NewHTTPError(422, err)
	}

	return c.JSON(http.StatusOK, res)
}
