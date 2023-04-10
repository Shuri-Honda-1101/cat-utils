package main

// NOTE:https://entgo.io/ja/docs/getting-started
// NOTE:https://zenn.dev/a_ichi1/articles/c9f3870350c5e2
// NOTE:https://github.com/mattn/echo-ent-example/blob/main/main.go

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Shuri-Honda-1101/cat-utils/ent"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// .envファイルを読み込み
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// MySQLデータベースに接続
	// ※generateしてからじゃないと"ent. Open"はundefined
	client, err := ent.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	// リソースを解放
	defer client.Close()
	// Entの自動マイグレーションツールを実行して、データベーススキーマを作成
	// NOTE:https://entgo.io/ja/docs/migrate/
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
