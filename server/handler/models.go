package handler

import (
	repModels "github.com/elias506/EchoRestAPI/repository/models"
	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
	IUserDB repModels.IUserDB
}
