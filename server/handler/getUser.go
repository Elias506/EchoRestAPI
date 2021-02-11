package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	repModels "github.com/elias506/EchoRestAPI/repository/models"
)

func GetUser(c echo.Context) error {
	cc := c.(*CustomContext)

	userID, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return err
	}
	user, err := repModels.IGetUser(cc.IUserDB, userID)
	if err != nil {
		return err
	}
	return cc.JSON(http.StatusOK, user)
}
