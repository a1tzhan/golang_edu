package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"

	"hw3/internal/db"
	"hw3/internal/handler"
	"hw3/internal/service"
)

func main(){
	ctxt := context.Background()
	db, err := db.NewPostgresDB(ctxt)
	if err != nil {
		fmt.Printf("Failed to connect to DB: %v\n", err)
		return
	}
	defer db.Close()

	my_srvc := service.NewService(db)
	my_hdlr := handler.NewHandler(my_srvc)

	e := echo.New()
	e.GET("/students/:id", my_hdlr.GetStudents)
	e.GET("/all_schedule", my_hdlr.GetAllSchedule)
	e.GET("/schedule/group/:id", my_hdlr.GetGroupSchedule)
	e.Logger.Fatal(e.Start(":8080"))
}