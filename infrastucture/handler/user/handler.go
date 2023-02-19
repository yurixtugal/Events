package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yurixtugal/Events/domain/user"
	"github.com/yurixtugal/Events/model"
)

type handler struct {
	useCase user.UseCase
}

func new(useCase user.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Create(c echo.Context) error {
	m := model.User{}
	if err := c.Bind(&m); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if err := h.useCase.Create(&m); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, m)
}

func (h handler) GetAll(c echo.Context) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}
