package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) healthcheck(c echo.Context) error {

	//how to use context :
	//context := c.Request().Context()
	//select {
	//case <-time.After(10 * time.Second):
	//	return c.JSON(http.StatusRequestTimeout, echo.Map{
	//		"error": "server time out",
	//	})
	//case <-context.Done():
	//	err := context.Err()
	//	fmt.Println("error: ", err)
	//}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "everything is ok !",
	})
}
