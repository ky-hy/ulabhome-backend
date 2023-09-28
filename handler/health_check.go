package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheck struct{}

func NewHealthCheckHandler() *HealthCheck {
	return &HealthCheck{}
}

// ヘルスチェック
// @Summary      ヘルスチェック
// @Description  ヘルスチェックを行うエンドポイント
// @Tags 共通
// @Success 200
// @Router /healthcheck [get]
func (hc *HealthCheck) ServeHTTP(c echo.Context) error {
	return SuccessResponse(c, http.StatusOK, nil)
}
