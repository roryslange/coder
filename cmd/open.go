package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
	fmt.Println("hello from open in its own file!")
}