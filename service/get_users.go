package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/ky-hy/ulabhome-backend/domain"
	"github.com/ky-hy/ulabhome-backend/domain/model"
	"github.com/ky-hy/ulabhome-backend/repository"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

type GetUsers struct {
	DB       repository.Queryer
	UserRepo domain.UserRepo
}

func NewGetUsers(db *sqlx.DB, repo domain.UserRepo) *GetUsers {
	return &GetUsers{DB: db, UserRepo: repo}
}

func (gu *GetUsers) GetUsers(ctx echo.Context) (model.Users, error) {
	// ユーザ一覧を取得する
	users, err := gu.UserRepo.GetAll(ctx, gu.DB)
	if err != nil {
		return users, xerrors.Errorf(": %w", err)
	}
	return users, nil
}
