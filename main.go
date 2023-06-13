package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sugarshop/asgard-gateway/db"
	"github.com/sugarshop/asgard-gateway/handler"
	"github.com/sugarshop/env"
)

func main() {
	// start config
	var conf string
	flag.StringVar(&conf, "conf", "conf/test.json", "specify the load config file")
	flag.Parse()

	// load env configuration
	env.LoadGlobalEnv(conf)

	engine := gin.New()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// init db
	db.Init()
	fmt.Println("[main]: db init success")

	// register other api
	handler.Register(engine)
	fmt.Println("[main]: handler register success")

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	go func() {
		engine.Run()
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGIN'T
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
