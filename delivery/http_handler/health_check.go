package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohamadafzal06/purl/repository/redis"
)

type Health struct {
	r redis.Redis
}

func NewHealth(r redis.Redis) Health {
	return Health{r: r}
}

func (h Health) HealthCheck(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Bad Request"})
	}
	err := h.r.Ping(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{"message": "Pinging the database failed"})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "everything is good!",
	})
}

func (h Health) Register(g *echo.Group) {
	g.GET("/health", h.HealthCheck)
}
