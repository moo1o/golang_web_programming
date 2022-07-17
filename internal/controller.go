package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}

func (controller *Controller) GetByID(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}
