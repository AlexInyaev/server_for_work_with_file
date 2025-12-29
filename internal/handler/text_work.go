package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/AlexInyaev/server_for_work_with_file/internal/handler/formatter"
)

func (h *Handler) TextWrite(c echo.Context) error {
	body := new(formatter.TextWriteRequest) // создает экземпляр структуры и возвращает указатель на него

	err := c.Bind(body) // заполняет экземпляр структуры на которую указывает body данными из json
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}

	// Удаляем пробелы
	body.Email = strings.TrimSpace(body.Email)
	if body.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "email is required",
		})
	}

	body.Name = strings.TrimSpace(body.Name)
	if body.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "name is required",
		})
	}

	err = h.exWriteText.Execute(c.Request().Context(), body.Name, body.Email, body.Text, body.Directory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"error": "text is required",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success": "ok",
	})
}
