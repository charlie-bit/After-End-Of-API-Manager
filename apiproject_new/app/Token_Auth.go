package app

import (
	"apiproject_new/models"
	"apiproject_new/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

/**
	Author:charlie
	Description:token_auth验证
	Time:2019-12-31
*/

func JwtAuthentication(ctx *gin.Context) bool {
		response := make(map[string]interface{})
		godotenv.Load("../.env")
		tokenHeader := ctx.Request.Header.Get("Authorization") //Grab the token from the header
		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			response = utils.Message("缺失 auth token")
			utils.Respond(500,response ,ctx)
			return false
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = utils.Message("auth token请以Bearer token格式输入")
			utils.Respond(500,response ,ctx)
			return false
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("Token_password")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			response = utils.Message("authentication token 输入错误")
			utils.Respond(500, response,ctx)
			return false
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = utils.Message("Token 无效.")
			utils.Respond(500, response ,ctx)
			return false
		}
		return true
}
