package config

import (
	"k8s.io/klog/v2"
)

// webserver config
type WebServerConfig struct {
	BindAddress string
	SSL         bool
	SSLCert     string
	SSLKey      string
}

func GetWebServerConfig() *WebServerConfig {
	cfg := &WebServerConfig{}

	cfg.BindAddress = DASHBOARD_CONFIG.GetString("webserver.bind_address")
	if cfg.BindAddress == "" {
		klog.Warningf("webserver.bind_address is empty, we use the default :8082")
		cfg.BindAddress = ":8082"
	}

	cfg.SSL = DASHBOARD_CONFIG.GetBool("webserver.ssl")
	if cfg.SSL {
		cfg.SSLCert = DASHBOARD_CONFIG.GetString("webserver.ssl_cert_file")
		cfg.SSLKey = DASHBOARD_CONFIG.GetString("webserver.ssl_key_file")
		if cfg.SSLCert == "" || cfg.SSLKey == "" {
			cfg.SSL = false
			klog.Warningf("Disable SSL since the cfg.SSLCert or cfg.SSLKey is empty.")
		}
	}

	return cfg
}
