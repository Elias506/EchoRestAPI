package server

import (
	"context"
	. "github.com/elias506/EchoRestAPI/models"
	r "github.com/elias506/EchoRestAPI/repository"
	h "github.com/elias506/EchoRestAPI/server/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"os/signal"
	"time"
)

type App struct {
	server *echo.Echo
	UserDB IUserDB
}

func NewApp() *App {
	return &App{
		server: echo.New(),
	}
}

func (a *App) Run(port string) error {
	// Init middleware
	a.server.Use(middleware.Logger())
	a.server.Use(middleware.Recover())

	// Init routes and custom context
	u := a.server.Group("/users")
	u.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{
				Context: c,
				UserDB:  a.UserDB,
			}
			return next(cc)
		}
	})
	u.GET("", h.GetUsers)
	u.GET("/:id", h.GetUser)
	u.DELETE("/:id", h.DeleteUser)
	u.POST("/", h.AddUser)
	u.PUT("/:id", h.UpdateUser)

	// Start listening
	a.server.Logger.Fatal(a.server.Start(":" + port))

	go func() {
		if err := a.server.Start(":" + port); err != nil {
			a.server.Logger.Info("Failed to listen and serve:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return a.Shutdown(ctx)
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

func (a *App) InitFileDB(path string) {
	a.UserDB = &r.FileDB{
		Path: path,
	}
}
