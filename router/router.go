package router

import (
	"http-theft-bank/handler/checkpoint1"
	"http-theft-bank/handler/checkpoint3"
	"http-theft-bank/handler/checkpoint4"
	"http-theft-bank/handler/checkpoint5"
	"net/http"

	"http-theft-bank/handler/checkpoint2"

	"http-theft-bank/handler/sd"
	"http-theft-bank/router/middleware"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "去哪呢？没路了哦！")
	})

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	organization := g.Group("/api/v1/organization")
	{
		organization.GET("/code", checkpoint1.CheckCode)
		organization.GET("/secret_key", middleware.AuthMiddleware(), checkpoint2.GetSecretKey)
		organization.GET("/iris_sample", middleware.AuthMiddleware(), checkpoint4.UserGetImage)
	}

	bank := g.Group("/api/v1/bank")
	bank.Use(middleware.AuthMiddleware())
	{
		bank.GET("/gate", checkpoint3.GetMethod)
		bank.POST("/gate", checkpoint3.PostMethod)
		bank.PUT("/gate", checkpoint3.PutMethod)
		bank.DELETE("/gate", checkpoint3.DelMethod)
		bank.PATCH("/gate", checkpoint3.PatchMethod)
		bank.GET("/iris_recognition_gate", checkpoint4.BackTips)
		bank.POST("/iris_recognition_gate", checkpoint4.VerifyParameter)

	}

	end := g.Group("/api/v1/muxi/backend/computer/examination")
	end.Use(middleware.AuthMiddleware())
	{
		end.GET("", checkpoint5.GetText)
		end.POST("", checkpoint5.UploadFile)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
