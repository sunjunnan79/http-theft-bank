package router

import (
	"http-theft-bank/handler/checkpoint3"
	"http-theft-bank/handler/checkpoint1"
	"http-theft-bank/handler/checkpoint5"
	"net/http"

	"http-theft-bank/handler/checkpoint2"
	"http-theft-bank/handler/sd"
	"http-theft-bank/router/middleware"

	"github.com/gin-gonic/gin"
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
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

  cp1 := g.Group("/origanization")
	{
		cp1.GET("/code", checkpoint1.CheckCode)
	}
  
	cp3 := g.Group("/bank/gate")
	{
		cp3.GET("", checkpoint3.GetMethod)
		cp3.POST("", checkpoint3.PostMethod)
		cp3.PUT("", checkpoint3.PutMethod)
		cp3.DELETE("", checkpoint3.DelMethod)
		cp3.PATCH("", checkpoint3.PatchMethod)
  }

	cp5 := g.Group("/muxi/backend/computer/examination")
	{
		cp5.GET("", checkpoint5.GetText)
		cp5.POST("", checkpoint5.UploadFile)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	organization := g.Group("/organization")
	organization.Use(middleware.AuthMiddleware())
	{
		organization.GET("/secret_key", checkpoint2.GetSercetKey)
	}

	return g
}
