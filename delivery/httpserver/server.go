package httpserver

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mohamadafzal06/purl/param"
)

type Service interface {
	Short(ctx context.Context, sReq param.ShortRequest) (param.ShortResponse, error)
	GetLong(ctx context.Context, surl param.LongRequest) (param.LongResponse, error)
}

type Server struct {
	service Service
}

func New(srv Service) Server {
	return Server{
		service: srv,
	}
}

func (s Server) Short(c echo.Context) error {
	if c.Request().Method != http.MethodPost {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Only except POST method"})

	}

	var resPram param.ShortRequest

	err := c.Bind(&resPram)
	if err != nil {
		log.Errorf("cannot bind the Body request into requeste param: %v\n", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Only except POST method"})
	}

	sResp, err := s.service.Short(c.Request().Context(), resPram)
	if err != nil {
		// TODO: check other possible error
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot short the url"})
	}
	// TODO: check better response format
	return c.JSON(http.StatusOK, echo.Map{"message": sResp})

}

func (s Server) Long(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Only except GET method"})
	}

	key := c.QueryParam("key")
	resPram := param.LongRequest{
		ShortURL: key,
	}

	sResp, err := s.service.GetLong(c.Request().Context(), resPram)
	if err != nil {
		// TODO: check other possible error
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot get the original url"})
	}
	// TODO: check better response format
	return c.JSON(http.StatusOK, echo.Map{"message": sResp})

}

func (s Server) Register(g *echo.Group) {
	g.GET("/long", s.Long)
	g.POST("/short", s.Short)
}
