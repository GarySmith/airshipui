package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	homedir "github.com/mitchellh/go-homedir"
	ps "github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	disableAuto bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "airshipui",
	Short: "airshipui is a graphical user interface for airship",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	Run: runOctantOrPlugin,
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
	fmt.Printf("Hello, world: %t\n", disableAuto)
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
		cmd := exec.Command("octant")
		fmt.Println("Launching octant\n")
		err := cmd.Run()
		log.Printf("Command finished with error: %v", err)
	} else {
		LaunchPlugin(cmd, args)
	}
}
