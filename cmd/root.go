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
	green    = color.New(color.FgGreen)
)

var rootCmd = &cobra.Command{
	Use:     "ypm",
	Short:   cyan.Sprint("YPM - Yieldaa! Package Manager"),
	Long:    cyan.Sprint("YPM (Yieldaa! Package Manager)") + ` - инструмент для работы с пакетами конфигураций бизнес-сущностей.`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func init() {
	// Добавляем функции для шаблона
	cobra.AddTemplateFunc("cyan", func(s string) string {
		return cyan.Sprint(s)
	})

	cobra.AddTemplateFunc("blue", func(s string) string {
		return blue.Sprint(s)
	})

	cobra.AddTemplateFunc("green", func(s string) string {
		return green.Sprint(s)
	})

	// Функция для проверки принадлежности к группе
	cobra.AddTemplateFunc("isProcessing", func(name string) bool {
		processingCmds := map[string]bool{
			"info": true, "scan": true, "analyse": true,
			"compile": true, "build": true,
		}
		return processingCmds[name]
	})

	cobra.AddTemplateFunc("isCreation", func(name string) bool {
		creationCmds := map[string]bool{
			"new": true, "entity": true,
		}
		return creationCmds[name]
	})

	// Шаблон с группировкой команд
	rootCmd.SetHelpTemplate(`{{blue "YPM v.1.0.0 | Yieldaa! Package Manager"}}

{{blue "ИСПОЛЬЗОВАНИЕ:"}}
{{.UseLine}}

{{blue "КОМАНДЫ ОБРАБОТКИ:"}}
{{range .Commands}}{{if and (or .IsAvailableCommand (eq .Name "help")) (isProcessing .Name)}}

  {{green .Name | printf "%-20s"}} {{.Short}}{{end}}{{end}}

{{blue "КОМАНДЫ СОЗДАНИЯ:"}}
{{range .Commands}}{{if and (or .IsAvailableCommand (eq .Name "help")) (isCreation .Name)}}

  {{green .Name | printf "%-20s"}} {{.Short}}{{end}}{{end}}

{{blue "СЛУЖЕБНЫЕ КОМАНДЫ:"}}
{{range .Commands}}{{if and (or .IsAvailableCommand (eq .Name "help")) (not (isProcessing .Name)) (not (isCreation .Name)) (ne .Name "help")}}

  {{green .Name | printf "%-20s"}} {{.Short}}{{end}}{{end}}

{{blue "ГЛОБАЛЬНЫЕ ФЛАГИ:"}}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}

{{blue "ИСПОЛЬЗУЙТЕ:"}}
  {{cyan "ypm [команда] --help"}} для подробностей о команде
`)

	// Добавляем команды в нужном порядке
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(analyseCmd)
	rootCmd.AddCommand(compileCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(entityCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		fmt.Fprintf(os.Stderr, "Используйте 'ypm --help' для справки\n")
		os.Exit(1)
	}
}
