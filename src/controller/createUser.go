package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	rest_err "github.com/larissa-oliv/meu-primeiro-crud-emgo/src/configuration"
	"github.com/larissa-oliv/meu-primeiro-crud-emgo/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds, error=%s\n", err.Error))

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
