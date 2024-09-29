package cmd

import (
	"github.com/changfenxia/scrapper-test/worker"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
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
func runParser(logger *zap.Logger) {
	// Создание воркеров
	categoryWorker := worker.NewCategoryWorker(logger)
	recipeWorker := worker.NewRecipeWorker(logger)

	// Парсинг категорий
	categories, err := categoryWorker.Start()
	if err != nil {
		logger.Fatal("Error parsing categories", zap.Error(err))
	}

	// Парсинг рецептов в каждой категории
	for _, category := range categories {
		recipes, err := recipeWorker.Start(category)
		if err != nil {
			logger.Error("Error parsing recipes", zap.String("category", category.Name), zap.Error(err))
		}

		// Логирование каждого рецепта
		for _, recipe := range recipes {
			logger.Info("Recipe processed", zap.String("Name", recipe.Name))
		}
	}
}
