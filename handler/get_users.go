package handler

import (
	"errors"
	"net/http"

	"github.com/ky-hy/ulabhome-backend/uerror"
	"github.com/labstack/echo/v4"
)

type GetUsers struct {
	Service GetUsersService
}

func NewGetUsers(s GetUsersService) *GetUsers {
	return &GetUsers{Service: s}
}

// ユーザ一覧取得ハンドラー
// @Summary      ユーザ一蘭取得
// @Description  ユーザの一覧を取得する
// @Tags ユーザ
// @Router /v1/users [get]
func (gu *GetUsers) ServeHTTP(ctx echo.Context) error {
	// サービスにユーザ仮登録処理を依頼
	users, err := gu.Service.GetUsers(ctx)

	// エラーレスポンスを返す
	if err != nil {
		if errors.Is(err, uerror.ErrDBException.Err) {
			input := ErrResponseInput{
				ErrCode: uerror.ErrDBException.Code,
				ErrMsg:  uerror.ErrDBException.RpsMsg,
				Status:  http.StatusInternalServerError,
				Err:     err,
			}
			return ErrResponse(ctx, input)
		}
		input := ErrResponseInput{
			ErrCode: uerror.ErrException.Code,
			ErrMsg:  uerror.ErrException.RpsMsg,
			Status:  http.StatusInternalServerError,
			Err:     err,
		}
		return ErrResponse(ctx, input)
	}

	type user struct {
		ID int `json:"id"`
	}

	usersResponse := []user{}
	for _, u := range users {
		usersResponse = append(usersResponse, user{
			ID: int(u.ID),
		})
	}
	rsp := struct {
		Users []user `json:"users"`
	}{Users: usersResponse}
	return SuccessResponse(ctx, http.StatusOK, rsp)
}
