package handler

import (
	repModels "github.com/elias506/EchoRestAPI/repository/models"
	"github.com/labstack/echo"
	"net/http"
)

func GetUsers(c echo.Context) error {
	cc := c.(*CustomContext)

	users, err := repModels.IGetUsers(cc.IUserDB)
	if err != nil {
		return err
	}
	return cc.JSON(http.StatusOK, users)
}
