package cmd

import (
	"fmt"
	"os"
	"yieldaa/cli/internal/scan"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scann package",
	Long:  "Scan package data",
	Run: func(cmd *cobra.Command, args []string) {
		// определение рабочей директории
		wd, err := os.Getwd()
		if err != nil {
			red.Printf("Error defining the working directory: %v\n", err)
			os.Exit(1)
		}

		// запуск сканирования
		result, err := scan.LoadPackage(wd)
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

func init() {
	rootCmd.AddCommand(scanCmd)
}

// // двухэтапный процессинг сущностей
// s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
// s.Suffix = "Scanning package..."
// s.Start()
// s.Stop()
// fmt.Println("✅ Package scan completed!")
