package router

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/ky-hy/ulabhome-backend/config"
	_ "github.com/ky-hy/ulabhome-backend/docs"
	"github.com/ky-hy/ulabhome-backend/handler"
	"github.com/ky-hy/ulabhome-backend/repository"
	"github.com/ky-hy/ulabhome-backend/service"
	"github.com/ky-hy/ulabhome-backend/utils"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// ルーティング
// 各層のDIを行う
func SetRouting(c context.Context, e *echo.Echo, db *sqlx.DB, cfg *config.Config) {
	// レポジトリ
	clocker := utils.RealClocker{}
	rep := repository.NewRepository(clocker)
	// トランザクション
	// tx := repository.NewTransaction(db)

	api := e.Group("v1")
	if cfg.Env == "development" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	hhc := handler.NewHealthCheckHandler()
	api.GET("/healthcheck", hhc.ServeHTTP)

	sgu := service.NewGetUsers(db, rep)
	hgu := handler.NewGetUsers(sgu)
	api.GET("/users", hgu.ServeHTTP)
}
