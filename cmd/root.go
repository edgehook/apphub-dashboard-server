package cmd

import (
	"os"

	"github.com/edgehook/apphub-dashboard-server/common/dbm"
	"github.com/edgehook/apphub-dashboard-server/common/dbm/model"
	"github.com/edgehook/apphub-dashboard-server/common/global"
	"github.com/edgehook/apphub-dashboard-server/webserver"
	"github.com/jwzl/beehive/pkg/core"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"github.com/kardianos/service"
	"k8s.io/component-base/logs"
)

var rootCmd = &cobra.Command{
	Use:     "ithings",
	Long:    `iot things manager.. `,
	Version: "0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: To help debugging, immediately log version
		klog.Infof("###########  Start the dashboard...! ###########")
		registerModules()
		// start all modules
		core.Run()
	},
}

func init() {
}

// register all module into beehive.
func registerModules() {
	webserver.Register()

	//init db
	pwd, _ := os.Getwd()
	dbPath := pwd + string(os.PathSeparator) + "dashboard.db"
	global.DBAccess = dbm.GormSQLite(dbPath)
	// global.DBAccess = edgeDbm.GormPostgreSQL()
	err := model.RegisterTables(global.DBAccess)
	if err != nil {
		panic(err)
	}
}

var logger service.Logger
var serviceConfig = &service.Config{
	Name:             "AppHubDashboardService",
	DisplayName:      "AppHub-Dashboard Service",
	Description:      "Advantech AppHub-Dashboard",
	UserName:         "",
	Arguments:        []string{},
	Executable:       "C:\\Program Files\\AppHub-Dashboard\\dashboard.exe",
	Dependencies:     []string{},
	WorkingDirectory: "",
	ChRoot:           "",
	Option:           map[string]interface{}{},
}

func NewAppService() {
	options := make(service.KeyValue)
	options["DelayedAutoStart"] = true
	options["StartType"] = "automatic"
	options["OnFailure"] = "restart"
	options["OnFailureDelayDuration"] = "1s"
	options["OnFailureResetPeriod"] = 10

	serviceConfig.Option = options

	prog := &Program{}
	s, err := service.New(prog, serviceConfig)
	if err != nil {
		klog.Errorf("create windows service with error: %s", err.Error())
	}

	errs := make(chan error, 5)
	// logger, err = s.Logger(errs)
	logger, err = s.SystemLogger(errs)
	if err != nil {
		klog.Errorf("windows service logger with err: %v", err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				klog.Errorf("windows service with err: %v", err)
			}
		}
	}()

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			klog.V(4).Infof("os.Args[2]: %v", os.Args[2])
			if os.Args[2] == "" {
				klog.Errorf("Executable Path is NULL, PLease check")
				return
			}
			serviceConfig.Executable = os.Args[2]
			s.Install()
			klog.V(4).Infof("Install Service Success")
			return
		}

		if os.Args[1] == "remove" {
			s.Uninstall()
			klog.V(4).Infof("Remove Service Success")
			return
		}
	}

	err = s.Run()
	if err != nil {
		klog.Errorf("windows service run with error: %s", err.Error())
	}
}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	klog.V(4).Infof("==========  Start dashboard Service ==========")
	if service.Interactive() {
		logger.Infof("dashboard running %v. Running in terminal.", service.Platform())
		klog.V(4).Infof("Running in terminal.")
	} else {
		logger.Infof("dashboard running %v. Running under service manager.", service.Platform())
		klog.V(4).Infof("Running under service manager.")
	}
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	klog.Errorf("==========  Stop dashboard Service ==========")
	logger.Info("dashboard Stopping!")

	logs.FlushLogs()
	return nil
}

func (p *Program) Shutdown(s service.Service) error {
	klog.Errorf("==========  Windows shutdown ==========")
	// klog.Errorf("Exit dashboard.exe")
	logger.Info("OS shutdown, dashboard Stopping!")
	logs.FlushLogs()
	return nil
}

func (p *Program) run() {
	defer func() {
		if err := recover(); err != nil {
			klog.Errorf("windows service dashboard failed with %s", err)
		}
	}()

	klog.V(4).Infof("###########  Start the dashboard...! ###########")
	//monitorplug.RegisterServiceCallbackHandler()

	registerModules()
	// start all modules
	core.Run()
}
