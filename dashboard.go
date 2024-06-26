package main

import (
	"flag"
	"os"

	"github.com/edgehook/apphub-dashboard-server/cmd"
	"k8s.io/component-base/logs"
)

func main() {
	//initial log.
	logFile := os.Args[0] + ".log"
	flag.Set("log_file", logFile)
	flag.Set("log_file_max_size", "50") //in MB, default as 50MB
	flag.Set("logtostderr", "false")
	flag.Set("alsologtostderr", "true")

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
