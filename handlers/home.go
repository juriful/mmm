package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

func Home(c echo.Context) error {
	status := map[string]string{"status": "success"}

	if err := c.JSON(http.StatusOK, status); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
