package cmd

import (
	"github.com/spf13/cobra"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/roryslange/coder/interfaces"
)

var openCmd = &cobra.Command{
	Use: "open [filePath]",
	Short: "open a file",
	Long: "open a file to make changes",
	Run: openFile,
} 

func init() {
	rootCmd.AddCommand(openCmd)
}

func openFile(cobra *cobra.Command, args []string) {
	var p *tea.Program

	switch len(args) {
	case 0:
		p = tea.NewProgram(interfaces.SimpleProgram("empty editor!"))
	case 1:
		p = tea.NewProgram(interfaces.OpenModel(args[0]))
	default:
		p = tea.NewProgram(interfaces.SimpleProgram("you put too many files to open"))
	}
	_, err := p.Run() // _ is supposed to be a model but idk what to do with it yet
	if err != nil {
		panic(err)
	}
}
