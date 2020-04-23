package main

import (
	"After-End-Of-API-Manager/apiproject_new/app"
	"After-End-Of-API-Manager/apiproject_new/controllers"
	_ "After-End-Of-API-Manager/apiproject_new/docs"
	"After-End-Of-API-Manager/apiproject_new/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io"
	"net/http"
	"os"
	"strings"
)

// @title Swagger Example API
// @version 1.0
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host 127.0.0.1:8000
// @BasePath /
func main() {
	models.Init()
	InitRouter()
	Log()
}

/**
	Author:charlie
	Description:日志文件
	Time:2019-1-10
*/
func Log()  {
	f,_ := os.Create("apiProject.log")
	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)
}
/**
	Author:charlie
	Description:初始化路由器
	Time:2019-1-10
*/
func InitRouter() {
	router := gin.New()
	router.Use(Cors()) //跨域
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	user := router.Group("/api/user/")
	user.Use(app.LimitByCache(models.ClientIP()),app.LimitRoleAuthentication())
	{
		user.POST("/new",controllers.NewUser)
		user.POST("/login",controllers.LoginUser)
		user.POST("/getToken",controllers.GetNewToken)
	}
	manager := router.Group("/api/manager/")
	manager.Use(app.LimitByCache(models.ClientIP()))
	{
		manager.POST("/login",controllers.ManagerLogin)
		manager.POST("/register",controllers.ManagerRegister)
		manager.GET("/selectManagerInfo",controllers.SelectApiUserInfo)
		manager.POST("/UpdateUsers",controllers.UpdateUsers)
		manager.POST("/RecorverPassword",controllers.RecoverPassword)
		manager.GET("/Delete",controllers.DeleteUser)
		manager.POST("/AddApiUrlLimit",controllers.AddApiURLLimit)
		manager.GET("/UpdateApiLimit",controllers.UpdateApiLimit)
	}
	godotenv.Load(".env")
	port := os.Getenv("port")
	port = ":" + port
	router.Run(port)
}
/**
	Author:charlie
	Description:跨域
	Time:2019-1-10
*/
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method      //请求方法
		origin := c.Request.Header.Get("Origin")        //请求头部
		var headerKeys []string                             // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")       // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()        //  处理请求
	}
}
