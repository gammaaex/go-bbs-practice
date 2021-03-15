package main

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/proullon/ramsql/driver"

	"myapp/pkg/handler"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	e := echo.New()

	db, err := sql.Open("ramsql", "bbs")
	if err != nil {
		e.Logger.Fatalf("sql.Open : Error : %s\n", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		e.Logger.Fatalf("failed to ping by error '%#v'", err)
		return
	}

	batch := []string{
		`CREATE TABLE messages (id VARCHAR(255) PRIMARY KEY, content VARCHAR(5000));`,
	}
	for _, b := range batch {
		if _, err = db.Exec(b); err != nil {
			e.Logger.Fatalf("sql.Exec: Error: %s\n", err)
		}
	}

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Static("/", "static")

	messageHandler := handler.NewMessage(db)
	messagesAPI := e.Group("/messages")
	messagesAPI.GET("", messageHandler.HandleGet)
	messagesAPI.POST("", messageHandler.HandlePost)
	messagesAPI.PUT("", messageHandler.HandlePut)
	messagesAPI.DELETE("", messageHandler.HandleDelete)

	e.Logger.Fatal(e.Start(":1323"))
}
