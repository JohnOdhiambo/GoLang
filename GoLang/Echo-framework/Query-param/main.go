package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func test(c echo.Context) error {
	return c.String(http.StatusOK, "Server is connected successfully")
}

func getCars(c echo.Context) error {
	carName := c.QueryParam("name")
	carMake := c.QueryParam("make")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Your car is: %s\nand Make is: %s\n", carName, carMake))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": carName,
			"make": carMake,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "Select String or JSON format",
	})

}

func main() {
	fmt.Println("Welcome to the server")

	e := echo.New()
	e.GET("/", test)
	e.GET("/Cars/:data", getCars)

	e.Start(":1323")
}
