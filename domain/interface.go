package domain

import (
	"github.com/ky-hy/ulabhome-backend/domain/model"
	"github.com/ky-hy/ulabhome-backend/repository"
	"github.com/labstack/echo/v4"
)

// Userに対するインターフェース
type UserRepo interface {
	GetAll(ctx echo.Context, db repository.Queryer) (model.Users, error)
}
