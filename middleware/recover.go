package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ky-hy/ulabhome-backend/uerror"
	"github.com/ky-hy/ulabhome-backend/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"golang.org/x/xerrors"
)

// パニック時の回復処理
func Recover() echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			// ログ設定
			logger := zerolog.
				New(os.Stdout).
				With().
				Timestamp().
				Logger().
				Output(zerolog.ConsoleWriter{Out: os.Stderr, FormatTimestamp: func(i any) string {
					clocker := utils.RealClocker{}
					// 日本直にする
					jst := time.FixedZone("Asia/Tokyo", 9*60*60)
					return clocker.Now().In(jst).Format(time.RFC3339)
				}})

			r := struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			}{
				Code:    uerror.ErrPanic.Code,
				Message: uerror.ErrPanic.RpsMsg,
			}
			js, merr := json.Marshal(r)
			if merr != nil {
				panic("can not marshal")
			}

			// ログの出力
			req := c.Request()
			res := c.Response()
			logger.Error().
				Str("URI", req.URL.Path).
				Str("method", req.Method).
				Str("user_agent", req.UserAgent()).
				Str("IP", req.RemoteAddr).
				// レスポンス
				Str("response", string(js)).
				Int("status", res.Status).
				// エラー
				Err(err).
				Int("ErrCode", uerror.ErrPanic.Code).
				Str("StackTrace", fmt.Sprintf("%+v", xerrors.Errorf("%w", err))).
				// パフォーマンス
				Msg("panic")

			// レスポンス
			return echo.NewHTTPError(http.StatusInternalServerError, r)
		},
	})
}
