package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var analyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "Проверка конфликтов ключей сущностей.",
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			red.Printf("Error defining the working directory: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("wd: %s", wd)
	},
}
