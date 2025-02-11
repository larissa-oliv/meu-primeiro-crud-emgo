package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/larissa-oliv/meu-primeiro-crud-emgo/src/controller"
)

func InitRoutes(r *gin.RouterGroup){
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)

}