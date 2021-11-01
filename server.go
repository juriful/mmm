package main

import (
	"github.com/labstack/echo"
	"mmm/handlers"
)

func main() {
	e := echo.New()

	e.GET("/", handlers.Home)
	e.POST("/git-hub-discussions", handlers.CreateGitHubDiscussions)

	e.Logger.Fatal(e.Start(":8888"))
}
