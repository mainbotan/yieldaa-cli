package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Инициализация нового пакета в текущей директории.",
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			red.Printf("Error defining the working directory: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("wd: %s", wd)
	},
}
