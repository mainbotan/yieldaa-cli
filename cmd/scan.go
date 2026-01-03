package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"yieldaa/cli/internal/scan"

	"github.com/briandowns/spinner"
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

		// загрузка конфигурации
		config, err := scan.ReadConfig(wd)
		if err != nil {
			red.Printf("Error read package configuration: %v\n", err)
			os.Exit(1)
		}
		var packageData scan.Package // пакетик
		cyan.Printf("Package: %s\n", config.Name)
		fmt.Printf("Version: %s\n", config.Version)
		fmt.Printf("Region:  %s\n", config.Region)

		// определение директории сущностей пакета
		entitiesDir := filepath.Join(wd, PACKAGE_ENTITIES_DIR_NAME)
		if _, err := os.Stat(entitiesDir); os.IsNotExist(err) {
			red.Printf("Package does not have '%s' directory\n", PACKAGE_ENTITIES_DIR_NAME)
			os.Exit(1)
		}

		// получение файлов сущностей
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = "Scanning package..."
		s.Start()

		entitiesFiles, err := scan.ScanEntities(entitiesDir)
		if err != nil {
			red.Printf("Error scan package: %v\n", err)
			os.Exit(1)
		}
		packageData.Files = entitiesFiles

		s.Stop()
		fmt.Println("✅ Package scan completed!")
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
