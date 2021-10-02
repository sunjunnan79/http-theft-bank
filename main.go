package main

import (
	"errors"
	"fmt"
	"http-theft-bank/pkg/cache"
	"http-theft-bank/pkg/text"
	"net/http"
	"time"

	"go.uber.org/zap"

	"http-theft-bank/config"
	"http-theft-bank/log"
	"http-theft-bank/router"
	"http-theft-bank/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

// @title http-theft-bank
// @version 1.0
// @description a game to study http with go
// @host localhost
// @BasePath /api/v1

// @tag.name organization
// @tag.description 组织服务
// @tag.name bank
// @tag.description 银行服务
// @tag.name end
// @tag.description 终点服务
func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// logger sync
	defer log.SyncLogger()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Init Cache
	cache.LocalStorage.Init()

	text.InitText()

	// Create the Gin engine.
	g := gin.New()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	g.MaxMultipartMemory = 8 << 20

	// Routes.
	router.Load(
		// Cores.
		g,

		// MiddleWares.
		middleware.Logging(),
		middleware.RequestId(),
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.",
				zap.String("reason", err.Error()))
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Info(
		fmt.Sprintf("Start to listening the incoming requests on http address: %s", viper.GetString("addr")))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
