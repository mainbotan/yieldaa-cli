package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	blue     = color.New(color.FgBlue)
	cyan     = color.New(color.FgCyan)
	cyanBold = color.New(color.FgCyan, color.Bold)
	blueBg   = color.New(color.BgBlue, color.FgWhite)
	red      = color.New(color.FgRed)
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
