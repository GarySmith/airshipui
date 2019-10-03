package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("plugin called")
	},
}

func init() {
	rootCmd.AddCommand(pluginCmd)
}
