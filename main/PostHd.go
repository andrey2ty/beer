package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func PostHandler(e echo.Context) error {
	var NameEn Name

	if err := e.Bind(&NameEn); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, Response{
			Status:  "error",
			Message: "error binding name",
		})
	}

	if err := db.Create(&NameEn).Error; err != nil {
		return e.JSON(http.StatusUnprocessableEntity, Response{
			Status:  "error",
			Message: "error create name",
		})
	}

	return e.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "create name success",
	})

}
