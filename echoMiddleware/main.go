package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Server running")

	s := echo.New()

	s.GET("/status", Handler)

	err := s.Start(":3000")
	if err != nil {
		log.Fatal(err)
	}

}

func Handler(s echo.Context) error {
	err := s.String(http.StatusOK, "test")
	if err != nil {
		return err
	}
	return nil
}
