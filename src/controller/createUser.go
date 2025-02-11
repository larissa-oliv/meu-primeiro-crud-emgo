package controller

import (
	"github.com/gin-gonic/gin"
	rest_err "github.com/larissa-oliv/meu-primeiro-crud-emgo/src/configuration"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rota de forma errada")
	c.JSON(err.Code,err)
}