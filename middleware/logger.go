package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/ky-hy/ulabhome-backend/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"golang.org/x/xerrors"
)

// コンテキストに保存するエラー情報のキー定義
const (
	err     = "err"
	errCode = "errCodo"
	// ユーザに返却するレスポンス
	rps = "rps"
)

type Log struct {
	errCode int
	rps     string
	err     error
}

// ログに保存する構造体の作成
// 最後に登録するにはLoggingを呼ぶ必要あり
func NewLog() *Log {
	return &Log{}
}
func (l *Log) ErrCode(code int) *Log {
	l.errCode = code
	return l
}
func (l *Log) Rsp(r string) *Log {
	l.rps = r
	return l
}
func (l *Log) Err(e error) *Log {
	l.err = e
	return l
}

// ログ情報を登録
func (l *Log) Logging(c echo.Context) {
	c.Set(rps, l.rps)
	c.Set(errCode, l.errCode)
	c.Set(err, l.err)
}

// ロギング処理
func Logger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogMethod:    true,
		LogError:     true,
		HandleError:  true,
		LogRemoteIP:  true,
		LogUserAgent: true,
		LogReferer:   true,
		LogLatency:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			// エラーデータ取得
			ec, _ := c.Get(errCode).(int)
			rps, _ := c.Get(rps).(string)
			e, _ := c.Get(err).(error)

			logger := zerolog.New(os.Stdout).
				With().
				Timestamp().
				Logger().
				Output(zerolog.ConsoleWriter{Out: os.Stderr, FormatTimestamp: func(i any) string {
					clocker := utils.RealClocker{}
					// 日本時間にする
					jst := time.FixedZone("Asia/Tokyo", 9*60*60)
					return clocker.Now().In(jst).Format(time.RFC3339)
				}})

			if v.Status >= 400 && v.Status < 500 {
				logger.Warn().
					// リクエスト
					Str("URI", v.URI).
					Str("method", v.Method).
					Str("user_agent", v.UserAgent).
					Str("IP", v.RemoteIP).
					// レスポンス
					Int("status", v.Status).
					Str("response", rps).
					// エラー
					Err(e).
					Int("ErrCode", ec).
					Str("StackTrace", fmt.Sprintf("%+v", xerrors.Errorf("%w", e))).
					// パフォーマンス
					Dur("latency(ms)", v.Latency).
					Msg("")
				return nil
			}
			if v.Status >= 500 {
				logger.Error().
					// リクエスト
					Str("URI", v.URI).
					Str("method", v.Method).
					Str("user_agent", v.UserAgent).
					Str("IP", v.RemoteIP).
					// レスポンス
					Str("response", rps).
					Int("status", v.Status).
					// エラー
					Err(e).
					Int("ErrCode", ec).
					Str("StackTrace", fmt.Sprintf("%+v", xerrors.Errorf("%w", e))).
					// パフォーマンス
					Dur("latency(ms)", v.Latency).
					Msg("")
				return nil
			}

			// 成功時
			logger.Info().
				// リクエスト
				Str("method", v.Method).
				Str("URI", v.URI).
				Str("user_agent", v.UserAgent).
				Str("IP", v.RemoteIP).
				// レスポンス
				Int("status", v.Status).
				Str("response", rps).
				// パフォーマンス
				// TODO: runtime, runtime/pprofパッケージを利用してメモリ利用率など出す
				Dur("latency(ms)", v.Latency).
				Msg("")
			return nil
		},
	})
}
