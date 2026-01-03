package cmd

import (
	"fmt"
	"os"
	"yieldaa/cli/internal/info"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "info",
	Short: " Package info",
	Run: func(cmd *cobra.Command, args []string) {
		// определение рабочей директории
		wd, err := os.Getwd()
		if err != nil {
			red.Printf("Error defining the working directory: %v\n", err)
			os.Exit(1)
		}

		// запуск сканирования
		result, err := info.LoadPackage(wd)
		if err != nil {
			red.Printf("Error scanning: %v\n", err)
			os.Exit(1)
		}

		// форматирование результата
		cyan.Printf("Package:         %s\n", result.Config.Name)
		fmt.Printf("Version:         %s\n", result.Config.Version)
		fmt.Printf("Region:          %s\n", result.Config.Region)
		fmt.Printf("Size:            %d\n", result.Sum.TotalSize)
		fmt.Printf("Entities count:  %d\n", result.Sum.EntitiesCount)
		fmt.Printf("Hash:            0x%08X\n", result.Sum.ControlHash)
	},
}
