package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yieldaa",
	Short: "Yieldaa Command Line Interface",
	Long:  `CLI tool for Yieldaa server management`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Yieldaa CLI - Type --help for usage")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}