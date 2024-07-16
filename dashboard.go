package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/edgehook/apphub-dashboard-server/cmd"
	"github.com/edgehook/apphub-dashboard-server/common/config"
	"k8s.io/component-base/logs"
)

func main() {
	//initial log.
	dashboardConfig := config.GetDashboardConfig()
	logLevel := dashboardConfig.LogLevel
	fLogger := dashboardConfig.FileLogger

	if fLogger {
		logFile := os.Args[0] + ".log"
		flag.Set("log_file", logFile)
		flag.Set("log_file_max_size", "5") //in MB, default as 1800MB
		flag.Set("logtostderr", "false")
		flag.Set("alsologtostderr", "false")
	} else {
		flag.Set("logtostderr", "true")
	}

	flag.Set("v", logLevel)
	logs.InitLogs()
	defer logs.FlushLogs()
	sys := runtime.GOOS

	if sys == "windows" {
		cmd.NewAppService()
	} else {
		if err := cmd.Execute(); err != nil {
			os.Exit(1)
		}
	}

}
