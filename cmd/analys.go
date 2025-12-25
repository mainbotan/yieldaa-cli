package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var analyseCmd = &cobra.Command{
	Use:	"analys",
	Short: 	"pkg analys",
	Long: 	"penis",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("penis")
	},
}

func init() {
	rootCmd.AddCommand(analyseCmd)
}