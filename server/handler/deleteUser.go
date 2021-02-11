package handler

import (
	repModels "github.com/elias506/EchoRestAPI/repository/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func DeleteUser(c echo.Context) error {
	cc := c.(*CustomContext)

	userID, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		return err
	}
	remoteUserID, err := repModels.IDeleteUser(cc.IUserDB, userID)
	if err != nil {
		return err
	}
	return cc.JSON(http.StatusOK, remoteUserID)
}
