package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Get(e echo.Context) error {
	var NameEm []Name

	if err := db.Find(&NameEm).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: "message no found",
		})
	}

	return e.JSON(http.StatusOK, &NameEm)

}
