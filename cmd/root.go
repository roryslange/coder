package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "coder",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.coder.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	setupConfig()

}

func setupConfig() {
	viper.SetConfigName("coder")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("priority", "medium")
	viper.SetDefault("file", filepath.Join(os.Getenv("$HOME"), ".coder.json"))

	viper.ReadInConfig()
	fmt.Println("here is the config file:" + viper.GetViper().ConfigFileUsed())
}
