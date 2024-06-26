package cmd

import (
	_ "github.com/edgehook/apphub-dashboard-server/common/dbm"
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
}
