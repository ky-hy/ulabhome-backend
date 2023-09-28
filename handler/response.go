package handler

import (
	"encoding/json"

	"github.com/ky-hy/ulabhome-backend/middleware"
	"github.com/labstack/echo/v4"
)

// 成功時のレスポンス
func SuccessResponse(c echo.Context, status int, data any) error {
	js, err := json.Marshal(data)
	if err != nil {
		panic("can not marshal")
	}
	// ログに使う情報登録
	middleware.NewLog().
		Rsp(string(js)).
		Logging(c)

	return c.JSON(status, data)
}

// エラー時の入力構造体
type ErrResponseInput struct {
	ErrCode int
	ErrMsg  string
	Status  int
	Err     error
}

// エラーレスポンス作成
func ErrResponse(c echo.Context, input ErrResponseInput) error {
	// エラーのレスポンスの型
	res := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    input.ErrCode,
		Message: input.ErrMsg,
	}
	js, err := json.Marshal(res)
	if err != nil {
		panic("can not marshal")
	}
	// ログに使う情報登録
	middleware.NewLog().
		ErrCode(input.ErrCode).
		Err(input.Err).
		Rsp(string(js)).
		Logging(c)

	return echo.NewHTTPError(input.Status, res)
}
