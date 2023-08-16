package main

import (
	"dghire.com/libs/go-monitor"
	"go-uc/vconfig"
	"go-uc/web"
)

func main() {
	monitor.Start(vconfig.AppName(), vconfig.MonitorPort())
	web.Run()
}
