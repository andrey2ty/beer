package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Name struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var db *gorm.DB

func main() {
	InitDB()

	e := echo.New()

	e.GET("/users", Get)
	e.POST("/users", PostHandler)
	e.PATCH("/users/:id", PatchHandler)
	e.DELETE("/users/:id", DeleteHandler)
	e.Start(":8080")

}
