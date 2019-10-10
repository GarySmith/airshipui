package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	ps "github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"opendev.org/airship/airshipui/internal/environment"
)

var (
	cfgFile     string
	disableAuto bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "airshipui",
	Short:   "airshipui is a graphical user interface for airship",
	Run:     runOctantOrPlugin,
	Version: environment.Version(),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Define glags and configuration settings
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.airshipui.yaml)")

	rootCmd.Flags().BoolVarP(&disableAuto, "no-auto", "n", false, "Disable auto-detection for deciding whether to run the plugin")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".airshipui" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".airshipui")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func runOctantOrPlugin(cmd *cobra.Command, args []string) {

	launchOctant := true
	if !disableAuto {
		ppid := os.Getppid()
		proc, err := ps.FindProcess(ppid)
		if err == nil {
			parentName := proc.Executable()

			if parentName == "airshipui" || parentName == "octant" {
				launchOctant = false
			}
		}
	}

	if launchOctant {
		myCmd, err := os.Executable()
		if err == nil {
			exePath := filepath.Dir(myCmd)
			pluginPath := os.Getenv("OCTANT_PLUGIN_PATH")
			pathList := append(filepath.SplitList(pluginPath), exePath)
			pluginPath = strings.Join(pathList, string(os.PathListSeparator))
			//os.Setenv("OCTANT_PLUGIN_PATH", pluginPath)

			command := exec.Command("/projects/octant/build/octant","-v")
			log.Printf("Launching octant with plugin path: %s\n", pluginPath)
			err := command.Run()
			log.Printf("Command finished with error: %v", err)
		}
	} else {
		// fmt.Printf("Launching airshipui as a plugin\n")
		LaunchPlugin(cmd, args)
	}
}
