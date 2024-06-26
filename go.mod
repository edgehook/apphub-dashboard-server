module github.com/edgehook/apphub-dashboard-server

go 1.16

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/google/uuid v1.3.0
	github.com/jwzl/beehive v1.0.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.0
	github.com/ugorji/go v1.1.7 // indirect
	gorm.io/driver/mysql v1.3.2
	gorm.io/driver/postgres v1.3.1
	gorm.io/driver/sqlite v1.3.1 // indirect
	gorm.io/gorm v1.23.1
	k8s.io/component-base v0.22.4
	k8s.io/klog v1.0.0
	k8s.io/klog/v2 v2.9.0
)
