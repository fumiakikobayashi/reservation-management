package main

import (
	"backend/src"
	"backend/src/Infrastructures"
	Shared2 "backend/src/Shared"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Errors loading .env file")
		panic(err.Error())
	}

	// DB接続
	db := Infrastructures.Init()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			panic("DB接続の初期化に失敗しました")
		}
	}(db)

	// ログ設定
	logger := Shared2.NewLogger()

	// 依存性の注入したハンドラーを取得
	handlers := src.NewHandlers(db, &logger)

	// echoの初期化
	e := echo.New()

	// カスタムエラーハンドラー
	e.HTTPErrorHandler = customHTTPErrorHandler

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/tasks", handlers.TaskHandler.GetTasks)

	// Start server
	if err := e.Start(":8080"); err != nil {
		panic(err.Error())
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		zap.S().Errorf("Unknown error: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	httpCode := he.Code
	switch err := he.Message.(type) {
	case error:
		switch {
		case httpCode >= 500:
			zap.S().Errorf("Server error: %v", err)
			if me, ok := err.(*Shared2.SampleError); ok {
				fmt.Print(me.StackTrace)
			}
		case httpCode >= 400:
			zap.S().Infof("Clients error: %v", err)
		}
		c.JSON(httpCode, "error")
	case string:
		// 存在しないエンドポイントが叩かれた場合
		zap.S().Errorf("Echo HTTP error: %v", he)
		c.JSON(http.StatusInternalServerError, he)
	default:
		zap.S().Errorf("Unknown HTTP error: %v", he)
		c.JSON(http.StatusInternalServerError, "予期せぬエラーが発生しました")
	}
}
