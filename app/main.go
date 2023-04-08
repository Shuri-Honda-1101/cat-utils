package main

import (
	"context"
	"log"
	"net/http"

	"entgo.io/ent/entc/integration/ent"
	"github.com/labstack/echo"
)

func main() {
	// MySQLデータベースに接続し、Ent ORMライブラリの自動マイグレーションツールを実行して、データベーススキーマを作成
    client, err := ent.Open("mysql", "<user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True")
    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }
	// リソースを解放
    defer client.Close()
    // Entの自動マイグレーションツールを実行して、データベーススキーマを作成
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

	// Echoのインスタンスを作成
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}