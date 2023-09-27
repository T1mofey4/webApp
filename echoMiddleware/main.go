package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Server running")

	s := echo.New()
	s.Use(MW)
	s.GET("/status", Handler)

	err := s.Start(":3000")
	if err != nil {
		log.Fatal(err)
	}

}

func Handler(ctx echo.Context) error {
	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d)
	s := fmt.Sprintf("Количество дней: %d", int64(dur.Hours())/24)
	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}

func MW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")
		if val == "admin" {
			log.Println("user detected")
		}
		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
