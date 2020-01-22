package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Message(message string) map[string]interface{} {
	return map[string]interface{}{"message":message}
}

func Respond(code int,message map[string]interface{},ctx *gin.Context)  {
	ctx.Set("content-type", "application/json")
	message["code"] = code
	ctx.JSON(http.StatusOK,message)
}
