package httpserver

import "github.com/labstack/echo/v4"

type Server struct {
	Router *echo.Echo
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
