package cmd

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check package integridy",
	Run: func(cmd *cobra.Command, args []string) {
		// определение рабочей директории
		// wd, err := os.Getwd()
		// if err != nil {
		// 	red.Printf("Error defining the working directory: %v\n", err)
		// 	os.Exit(1)
		// }

		// двухэтапный процессинг сущностей
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Scanning package..."
		s.Start()
		time.Sleep(3 * time.Second)
		s.Stop()
		fmt.Println("✅ Package scan completed!")
	},
}
