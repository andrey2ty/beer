package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func DeleteHandler(e echo.Context) error {
	isParam := e.Param("id")
	id, err := strconv.Atoi(isParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "not strconv id",
		})
	}

	if err := db.Model(&Name{}).Where("id = ?", id).Delete(&Name{}).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, Response{
			Status:  "error",
			Message: "id not found",
		})
	}

	return e.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: "id is deleted",
	})

}
