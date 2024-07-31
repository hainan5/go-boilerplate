package controllers

import (
	"github.com/gin-gonic/gin"
)

type CommonController struct {
	BaseController
}

func (ctrl *CommonController) SayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func (ctrl *CommonController) Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
