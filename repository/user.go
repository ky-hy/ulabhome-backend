package repository

import (
	"github.com/ky-hy/ulabhome-backend/domain/model"
	"github.com/ky-hy/ulabhome-backend/uerror"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

// ユーザ一覧取得
func (r *Repository) GetAll(ctx echo.Context, db Queryer) (model.Users, error) {
	sql := `
		SELECT *
		FROM users;
	`
	var users model.Users
	if err := db.SelectContext(ctx.Request().Context(), &users, sql); err != nil {
		return users, xerrors.Errorf(": %w: %s", uerror.ErrDBException.Err, err.Error())
	}
	return users, nil
}
