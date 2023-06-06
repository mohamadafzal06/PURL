package httpserver

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mohamadafzal06/purl/param"
)

type PurlService interface {
	Short(ctx context.Context, sReq param.ShortRequest) (param.ShortResponse, error)
	GetLong(ctx context.Context, surl param.LongRequest) (param.LongResponse, error)
	GetLongInfo(ctx context.Context, surl param.LongRequest) (param.LongResponse, error)
}

type Handler struct {
	schema  string
	host    string
	service PurlService
}

func NewServer(schema, host string, srv PurlService) Handler {
	return Handler{
		schema:  schema,
		host:    host,
		service: srv,
	}
}

func (h Handler) Short(c echo.Context) error {
	if c.Request().Method != http.MethodPost {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"message": "Only except POST method"})
	}

	reqPram := new(param.ShortRequest)

	err := c.Bind(reqPram)
	if err != nil {
		log.Errorf("cannot bind the Body request into request param: %v\n", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "the requsted url for shortning has problem"})
	}

	sResp, err := h.service.Short(c.Request().Context(), *reqPram)
	if err != nil {
		// TODO: check other possible error
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot short the url"})
	}

	// TODO: check better response format
	return c.JSON(http.StatusOK, echo.Map{"message": sResp})
}

func (h Handler) Long(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Only except GET method"})
	}

	key := c.Param("key")
	resPram := param.LongRequest{
		Key: key,
	}

	sResp, err := h.service.GetLong(c.Request().Context(), resPram)
	if err != nil {
		// TODO: check other possible error
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "cannot get the original url"})
	}
	// TODO: check better response format
	return c.JSON(http.StatusOK, echo.Map{"message": sResp})

}

func (h Handler) Info(c echo.Context) error {
	return nil
}

func (h Handler) Register(g *echo.Group) {
	g.POST("/short", h.Short)
	g.GET("/long/:key", h.Long)
	g.GET("/long/:key/info", h.Info)
}
