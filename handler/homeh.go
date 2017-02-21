package handler

import (
	"github.com/gin-gonic/gin"
)

func Homeh(c *gin.Context){
	c.String(200,"hello %s","world")
}