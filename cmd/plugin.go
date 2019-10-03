package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"opendev.org/airship/airshipui/internal/environment"
	"opendev.org/airship/airshipui/internal/plugin"
)

var pluginName = "airshipui"

// pluginCmd represents the plugin command
var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Serve up an octant plugin",
	Long: `The airshipui is a wrapper around octant, 
https://github.com/vmware-tanzu/octant, but also includes
a plugin providing airship-specific functionality.  When
invoked with this argument, it runs the plugin, which normally
has to be executed from octant.

When airshipui is invoked without arguments and its parent
process is either airshipui or octant, then the plugin argument
is assumed`,
	Run: LaunchPlugin,
}

func init() {
	rootCmd.AddCommand(pluginCmd)
}

// This is a sample plugin showing the features of Octant's plugin API.
func LaunchPlugin(cmd *cobra.Command, args []string) {
	// Remove the prefix from the go logger since Octant will print logs with timestamps.
	log.SetPrefix("")

	description := fmt.Sprintf("Airship UI version %s", environment.Version())
	// Use the plugin service helper to register this plugin.
	p, err := plugin.Register(pluginName, description)
	if err != nil {
		log.Fatal(err)
	}

	// The plugin can log and the log messages will show up in Octant.
	log.Printf("%s is starting", pluginName)
	p.Serve()
}
