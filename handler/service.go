package handler

import (
	"github.com/ky-hy/ulabhome-backend/domain/model"
	"github.com/labstack/echo/v4"
)

type GetUsersService interface {
	GetUsers(ctx echo.Context) (model.Users, error)
}
