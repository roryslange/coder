package cmd

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/roryslange/coder/interfaces"
	"github.com/spf13/cobra"
)

type config struct {
	helloMsg string
	configFile os.File
}

var configCmd = &cobra.Command{
	Use: "config",
	Short: "View and Edit configuration",
	Long: "Edit configurations declared in ",
	Run: configure,
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func configure(cmd *cobra.Command, args []string) {
	p := tea.NewProgram(interfaces.SimpleProgram("hello from config"))
	_, err := p.Run()
	if err != nil {
		panic(err)
	}
}