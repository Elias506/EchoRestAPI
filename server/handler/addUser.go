package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"github.com/elias506/EchoRestAPI/models"
	repModels "github.com/elias506/EchoRestAPI/repository/models"
)

func AddUser(c echo.Context) error {
	cc := c.(*CustomContext)

	reqUser := &RequestUser{}
	if err := cc.Bind(reqUser); err != nil {
		fmt.Println(err)
		return err
	}
	newUserID, err := repModels.IAddUser(cc.IUserDB, reqUser)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return cc.JSON(http.StatusOK, newUserID)
}
