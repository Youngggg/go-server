package web

import (
	"fmt"
	"net/http"

	"apple/common/env"
	"apple/common/ginx"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// gin的Engine指针实例变量
var Engine *gin.Engine

func init() {

	if env.IsProd() {
		gin.SetMode("release")
	} else {
		gin.SetMode("debug")
	}

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true

	Engine = gin.New()
	Engine.NoRoute(ginx.NoRoute)
	Engine.NoMethod(ginx.NoMethod)
	Engine.Use(gzip.Gzip(gzip.DefaultCompression))
	Engine.Use(cors.New(corsCfg))
	Engine.Use(ginx.Recovery())
	Engine.Use(ginx.AccessHandlerFunc)
}

func StartWeb() {
	listen()
}

func listen() {
	addr := ":8001"
	if err := gracehttp.Serve(&http.Server{Addr: addr, Handler: Engine}); err != nil {
		fmt.Println(err)
	}
}
