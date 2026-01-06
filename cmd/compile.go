package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Компиляция пакета.",
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			red.Printf("Error defining the working directory: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("wd: %s", wd)
	},
}
