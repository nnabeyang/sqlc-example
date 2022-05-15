package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/mattn/go-slim"
	"github.com/nnabeyang/sqlc-example/database"
)

func main() {
	db, err := sql.Open("mysql", os.Getenv("DSN")+"?tls=false&parseTime=true")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	t, err := slim.ParseFile("template/index.slim")
	if err != nil {
		log.Fatal(err)
	}
	client := database.New(db)
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		entries, err := client.GetEntries(context.Background(), 10)
		if err != nil {
			return err
		}
		c.Request().Header.Set("content-type", "text/html")
		return t.Execute(c.Response(), map[string]interface{}{
			"entries": entries,
		})
	})

	e.POST("/add", func(c echo.Context) error {
		if _, err := client.CreateEntry(context.Background(), c.FormValue("content")); err != nil {
			log.Println(err.Error())
			return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
		return c.Redirect(http.StatusFound, "/")
	})
	e.Logger.Fatal(e.Start(":8989"))
}
