package handler

import (
	repModels "github.com/elias506/EchoRestAPI/repository/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
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
