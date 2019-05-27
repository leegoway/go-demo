package main

import (
	"fmt"
	"runtime"
	"flag"
	"demo/config"
	"github.com/gin-gonic/gin"
    "github.com/fvbock/endless"
	"demo/routers"
	_ "demo/models"
	)

var (
	configFile = flag.String("config", "", "config file path");
	logLevel = flag.String("log level", "info", "log level");
	port = flag.Int("port", 8850, "http port");
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	var cfg *config.Config
	var err error

	if len(*configFile) == 0 {
		fmt.Println("use default config")
		cfg = config.NewConfigDefault()
	} else {
		if cfg, err = config.NewConfigWithFile(*configFile); err != nil {
			fmt.Println(err)
		}
	}

	if cfg.AppMode == "product" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	routers.InitRouters(r)
	err = endless.ListenAndServe(cfg.HttpAddr, r)
	if err != nil {
		fmt.Println(err)
	}
}
