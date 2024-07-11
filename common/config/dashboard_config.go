package config

import (
	"sync"

	_ "k8s.io/klog/v2"
)

const ()

var DASHBOARD_CONFIG *Config
var once = sync.Once{}

type DashboardConfig struct {
	LogLevel   string
	FileLogger bool
}

// get the db config from config file.
func GetDashboardConfig() *DashboardConfig {
	dashboardConfig := &DashboardConfig{}
	logLevel := DASHBOARD_CONFIG.GetString("logger.log_level")
	if logLevel == "" {
		logLevel = "info"
	}
	switch logLevel {
	case "debug":
		dashboardConfig.LogLevel = "4"
	default:
		dashboardConfig.LogLevel = "3"
	}
	dashboardConfig.FileLogger = DASHBOARD_CONFIG.GetBool("logger.file_logger")
	return dashboardConfig
}

func ForceInit() {
	DASHBOARD_CONFIG = NewYamlConfig("config.yaml")
}

func init() {
	once.Do(func() {
		//load the config.yaml from conf/
		DASHBOARD_CONFIG = NewYamlConfig("config.yaml")
	})
}
