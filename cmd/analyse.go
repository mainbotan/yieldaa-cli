package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var analyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "Analyse package",
	Long:  "Analyse package data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting analysis...")

		bar := progressbar.NewOptions(100,
			progressbar.OptionSetDescription("Processing data..."),
			progressbar.OptionSetWriter(os.Stderr),
			progressbar.OptionShowCount(),
			progressbar.OptionShowIts(),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "=",
				SaucerHead:    ">",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}),
			progressbar.OptionOnCompletion(func() {
				fmt.Fprint(os.Stderr, "\n")
			}),
		)

		for i := 0; i < 100; i++ {
			bar.Add(1)
			time.Sleep(50 * time.Millisecond)
		}

		fmt.Println("✅ Analysis complete!")
	},
}
var analyse2Cmd = &cobra.Command{
	Use:   "analyse",
	Short: "Analyse package",
	Long:  "Analyse package data",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Processing package data..."
		s.Start()

		// Ваша обработка данных
		time.Sleep(3 * time.Second) // Имитация работы

		s.Stop()
		fmt.Println("✅ Package analysis completed!")
	},
}

func init() {
	rootCmd.AddCommand(analyseCmd)
}
