package handler

import (
	. "github.com/elias506/EchoRestAPI/models"
	"github.com/labstack/echo"
	"net/http"
)

func AddUser(c echo.Context) error {
	cc := c.(*CustomContext)

	reqUser := &RequestUser{}
	if err := cc.Bind(reqUser); err != nil {
		return cc.JSON(http.StatusBadRequest, err)
	}
	newUserID, err := IAddUser(cc.UserDB, reqUser)
	if err != nil {
		return cc.JSON(http.StatusInternalServerError, err)
	}
	return cc.JSON(http.StatusOK, newUserID)
}
