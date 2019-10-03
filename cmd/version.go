package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"opendev.org/airship/airshipui/internal/environment"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version number of airshipui",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(environment.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
