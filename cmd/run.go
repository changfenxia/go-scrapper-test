package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	// Параметры парсера
	parseType   int // Тип парсинга (1 - полный, 2 - только категории)
	maxRecipes  int // Количество рецептов для парсинга в каждой категории
	concurrency int // Количество одновременных потоков (горутин)

	// Команда для запуска парсера
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Запуск парсера",
		Long:  "Запуск парсера с параметрами: тип парсинга, количество рецептов и одновременных потоков.",
		Run:   runParser,
	}
)

// init добавляет флаги для команды run
func init() {
	rootCmd.AddCommand(runCmd)

	// Добавление флагов к команде run
	runCmd.Flags().IntVarP(&parseType, "type", "t", 1, "Тип парсинга: 1 - полный, 2 - только категории")
	runCmd.Flags().IntVarP(&maxRecipes, "recipes", "r", 10, "Количество рецептов для каждой категории")
	runCmd.Flags().IntVarP(&concurrency, "concurrency", "g", 5, "Количество одновременных потоков")
}

// runParser запускает парсер с заданными параметрами
func runParser(cmd *cobra.Command, args []string) {
	log.Println("Запуск парсера с параметрами:")
	log.Printf("Тип парсинга: %d\n", parseType)
	log.Printf("Количество рецептов: %d\n", maxRecipes)
	log.Printf("Количество одновременных потоков: %d\n", concurrency)

	// Здесь добавьте логику запуска парсера
	// Например, вызов функции, которая запускает парсинг с указанными параметрами.
}
