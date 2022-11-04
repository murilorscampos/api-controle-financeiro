package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/murilorscampos/api-controle-/controllers"
)

func HandleRoutes()  {

	r := gin.Default()
	r.GET("/", controllers.Inicia)
	r.Run()
	
}