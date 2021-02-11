package handler

import (
	"github.com/labstack/echo"
	"net/http"
	repModels "github.com/elias506/EchoRestAPI/repository/models"
)

func GetUsers(c echo.Context) error {
	cc := c.(*CustomContext)

	users, err := repModels.IGetUsers(cc.IUserDB)
	if err != nil {
		return err
	}
	return cc.JSON(http.StatusOK, users)
}