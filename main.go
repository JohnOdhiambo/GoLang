package main

import (
	"os"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	"basic-api/controller"
	//"basic-api/config"
)
func  main()  {
	//Create Log
		// Open the log file in append mode. Change the file path and name as per your requirement.
	file, err := os.OpenFile("/usr/local/go/api-test/log/mylog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Set the log output to the file
	log.SetOutput(file)	


	e := echo.New()
	e.POST("/api/insertTicket", controller.InsertTicket)
	e.Logger.Fatal(e.Start(":8080"))
	// e.GET("/", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, map[string]interface{}{
	// 		"Get started": "Jay",
	// 	})
	// })
	

	// // ORM Connection To Database
	// config.DatabaseInit()
	// gorm := config.DB()

	// dbGorm, err := gorm.DB()
	// if err != nil {
	// 	panic(err)
	// }

	// dbGorm.Ping()

	// //SQL Connection to DB
	//dbConn := config.Connect()
	

	//ticketRoute := e.Group("/ticket")
	//ticketRoute.POST("/", controller.CreateTicket)
	// ticketRoute.GET("/:id", controller.GetBook)
	// ticketRoute.PUT("/:id", controller.UpdateBook)
	// ticketRoute.DELETE("/:id", controller.DeleteBook)
	
	//e.Logger.Fatal(e.Start(":8080"))
}