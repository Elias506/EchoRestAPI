package handler

import (
	"github.com/labstack/echo"
	repModels "github.com/elias506/EchoRestAPI/repository/models"
)

type CustomContext struct {
	echo.Context
	IUserDB repModels.IUserDB
}