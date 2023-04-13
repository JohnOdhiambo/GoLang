// Mostly handle the list of routes
package main

import (
	"api-test/cmd/api/handlers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostIndexHandler)     //Gets the list of all posts
	e.GET("/post/:id", handlers.PostSingleHandler) //Gets a specific post

	e.Logger.Fatal(e.Start(":1323"))
}
