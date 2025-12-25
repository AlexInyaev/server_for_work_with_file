package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/AlexInyaev/server_for_work_with_file/internal/executor"
)

type Handler struct {
	// common
	e           *echo.Echo
	exWriteText *executor.WriteText
}

func NewHandler(
	e *echo.Echo,
	exWriteText *executor.WriteText,
) *Handler {
	return &Handler{
		e:           e,
		exWriteText: exWriteText,
	}
}

func (h *Handler) InitRoutes() {
	h.e.POST("/user", h.TextWrite)
}

/*func(c echo.Context) error {
	var user User
	// Преобразуем JSON в структуру User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	// Возвращаем полученные данные в формате JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Сообщение получено",
		"user":    user,
	})
})
*/
