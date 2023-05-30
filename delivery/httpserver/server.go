package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohamadafzal06/purl/param"
	"github.com/mohamadafzal06/purl/service"
)

type Server struct {
	Router  *echo.Echo
	Service service.Service
}

func New(router *echo.Echo) *Server {
	return &Server{
		Router: router,
	}
}

func (s *Server) Serve() {
	// Routes
	s.Router.GET("/health-check", s.healthCheck)
}

func (s *Server) Short(c echo.Context) error {
	url := c.Request().URL.Query().Get("url")
	resPram := param.ShortRequest{
		LongURL: url,
	}

	sResp, err := s.Service.Short(c.Request().Context(), resPram)
	if err != nil {
		// TODO: check other possible error
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot short the url"})
	}
	// TODO: check better response format
	return c.JSON(http.StatusOK, echo.Map{"message": sResp})

}

func (s *Server) Long(c echo.Context) error {
	key := c.Request().URL.Query().Get("key")
	resPram := param.LongRequest{
		ShortURL: key,
	}

	sResp, err := s.Service.GetLong(c.Request().Context(), resPram)
	if err != nil {
		// TODO: check other possible error
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot get the original url"})
	}
	// TODO: check better response format
	return c.JSON(http.StatusOK, echo.Map{"message": sResp})

}
