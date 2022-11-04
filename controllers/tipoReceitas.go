package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Inicia(c *gin.Context)  {

	c.SecureJSON(http.StatusOK, gin.H{"data": "Iniciando servidor"})
	
}