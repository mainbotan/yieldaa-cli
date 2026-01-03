package cmd

import (
	"fmt"
	"os"
	"time"
	"yieldaa/cli/internal/scan"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan package",
	Run: func(cmd *cobra.Command, args []string) {
		// определение рабочей директории
		wd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error defining the working directory: %v\n", err)
			os.Exit(1)
		}

		// сканирование сущностей с таймером и спиннером
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Scanning package entities..."
		s.Start()

		startTime := time.Now()
		entities, scanErr := scan.ScanEntities(wd)
		elapsed := time.Since(startTime)

		s.Stop()

		// Вывод времени сканирования
		fmt.Printf("✓ Scanned %d entities in %v\n",
			len(entities),
			formatDuration(elapsed))

		// Вывод ошибок, если есть
		if scanErr != nil {
			fmt.Printf("\nErrors:\n%s\n", scanErr)
		}
	},
}
