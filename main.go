package main

import (
	"fmt"
	"runtime"
	"flag"
	"github.com/leegoway/go-demo/config"
	"github.com/gin-gonic/gin"
    "github.com/fvbock/endless"
	"github.com/leegoway/go-demo/routers"
	_ "github.com/leegoway/go-demo/model"
)

var (
	configFile = flag.String("config", "", "config file path");
	logLevel = flag.String("log level", "info", "log level");
	port = flag.Int("port", 8850, "http port");
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	var err error
	if len(*configFile) == 0 {
		fmt.Println("use default config")
		config.Cfg = config.NewConfigDefault()
	} else {
		if config.Cfg, err = config.NewConfigWithFile(*configFile); err != nil {
			fmt.Println(err)
		}
	}

	if config.Cfg.AppMode == "product" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	fmt.Print(config.Cfg)
	r := routers.InitRouters()
	err = endless.ListenAndServe(config.Cfg.HttpAddr, r)
	if err != nil {
		fmt.Println(err)
	}
}
