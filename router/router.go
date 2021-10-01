package router

import (
	"http-theft-bank/handler/checkpoint1"
	"http-theft-bank/handler/checkpoint5"
	"net/http"

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

	return g
}
