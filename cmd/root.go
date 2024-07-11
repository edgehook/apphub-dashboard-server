package cmd

import (
	"os"

	"github.com/edgehook/apphub-dashboard-server/common/dbm"
	_ "github.com/edgehook/apphub-dashboard-server/common/dbm"
	"github.com/edgehook/apphub-dashboard-server/common/dbm/model"
	"github.com/edgehook/apphub-dashboard-server/common/global"
	"github.com/edgehook/apphub-dashboard-server/webserver"
	"github.com/jwzl/beehive/pkg/core"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
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
