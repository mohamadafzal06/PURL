package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health struct {
}

func NewHealth() Health {
	return Health{}
}

func (h Health) HealthCheck(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Bad Request"})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "everything is good!",
	})
}

func (h Health) Register(g *echo.Group) {
	g.GET("/health", h.HealthCheck)
}
