package models

import (
	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
	UserDB IUserDB
}
