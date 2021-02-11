package handler

import (
	. "github.com/elias506/EchoRestAPI/models"
	"github.com/labstack/echo"
	"net/http"
)

func GetUsers(c echo.Context) error {
	cc := c.(*CustomContext)

	users, err := IGetUsers(cc.UserDB)
	if err != nil {
		return cc.JSON(http.StatusBadRequest, err)
	}
	return cc.JSON(http.StatusOK, users)
}
