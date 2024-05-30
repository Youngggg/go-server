package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"apple/common/cfg"
	"apple/common/log"
	"apple/server/boluo/web"

	_ "go.uber.org/automaxprocs"
)

func main() {
	// 初始化log
	log.InitLog(cfg.GetDefaultLogConfig("boluo"))

	// 启动profile监控
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// 初始化web
	web.StartWeb()
}
