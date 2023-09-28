package main

import (
	"context"
	"fmt"

	"github.com/ky-hy/ulabhome-backend/config"
	"github.com/ky-hy/ulabhome-backend/middleware"
	"github.com/ky-hy/ulabhome-backend/repository"
	"github.com/ky-hy/ulabhome-backend/router"
	"github.com/labstack/echo/v4"
)

// @title ulabhome API
// @version 1.0
// @description ulabhomeのAPIです
// @termsOfService

// @contact.name ulabhome
// @contact.url https://github.com/ky-hy/ulabhome-backend/

// @host localhost:8081
// @BasePath /v1
func main() {
	run(context.Background())
}

func run(c context.Context) {
	e := echo.New()

	cfg, err := config.New()
	if err != nil {
		e.Logger.Fatal(err)
	}
	// ミドルウェアの設定（順番注意）
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	// DB関係初期化
	db, cleanup, err := repository.NewDB(c, cfg)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer cleanup()

	router.SetRouting(c, e, db, cfg)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", cfg.Port)))
}
