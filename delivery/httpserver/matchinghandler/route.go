package matchinghandler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(e *echo.Echo) {
	grp := e.Group("/matching")

	grp.POST("/add-to-waiting-list", h.addToWaitingList)
}
