package handler

import (
	"fmt"
	. "github.com/elias506/EchoRestAPI/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func UpdateUser(c echo.Context) error {
	cc := c.(*CustomContext)
	userID, err := strconv.Atoi(cc.Param("id"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	reqUser := &RequestUser{}
	if err := cc.Bind(reqUser); err != nil {
		fmt.Println(err)
		return err
	}
	updatedUserID, err := IUpdateUser(cc.UserDB, userID, reqUser)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if updatedUserID == -1 {
		return cc.JSON(http.StatusNotFound, userID)
	}

	return cc.JSON(http.StatusOK, updatedUserID)
}
