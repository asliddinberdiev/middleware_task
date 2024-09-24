package app

import (
	"log"

	"github.com/asliddinberdiev/middleware_task/internal/app/endpoint"
	"github.com/asliddinberdiev/middleware_task/internal/app/middleware"
	"github.com/asliddinberdiev/middleware_task/internal/app/service"
	"github.com/labstack/echo"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(middleware.RoleCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	log.Println("server running")

	if err := a.echo.Start(":9090"); err != nil {
		log.Fatal("server error: ", err)
	}

	return nil
}
