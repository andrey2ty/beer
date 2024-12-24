package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func PatchHandler(e echo.Context) error {
	isParam := e.Param("id")
	id, err := strconv.Atoi(isParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "id must be a number",
		})
	}

	var NameMl Name

	if err := e.Bind(&NameMl); err != nil {
		return e.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "name not binding",
		})
	}

	if err := db.Model(&Name{}).Where("id=?", id).Updates(Name{
		Name:  NameMl.Name,
		Email: NameMl.Email,
	}).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: " id not found",
		})
	}

	return e.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "name changed",
	})
}
