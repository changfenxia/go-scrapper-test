package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/seniorcat/scraper/worker" // Укажите правильный путь к пакету cmd
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

	// Инициализация логгера zap
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Не удалось инициализировать логгер: %v", err)
	}
	defer logger.Sync() // Синхронизация логов перед завершением работы

	// Создание воркеров для категорий и рецептов
	categoryWorker := worker.NewCategoryWorker(logger)
	recipeWorker := worker.NewRecipeWorker(logger)

	// В зависимости от parseType запускаем парсинг
	switch parseType {
	case 1:
		// Полный парсинг: категории + рецепты
		log.Println("Запуск полного парсинга...")

		// Парсинг категорий
		categories, err := categoryWorker.Start()
		if err != nil {
			logger.Error("Ошибка при парсинге категорий", zap.Error(err))
			return
		}

		// Логирование найденных категорий
		for _, category := range categories {
			logger.Info("Категория", zap.String("Name", category.Name), zap.String("Href", category.Href))

			// Парсинг рецептов в каждой категории
			recipes, err := recipeWorker.Start(category)
			if err != nil {
				logger.Error("Ошибка при парсинге рецептов", zap.String("Category", category.Name), zap.Error(err))
				continue
			}

			// Логирование найденных рецептов
			for _, recipe := range recipes {
				logger.Info("Рецепт", zap.String("Name", recipe.Name), zap.String("Href", recipe.Href))
			}
		}

	case 2:
		// Парсинг только категорий
		log.Println("Запуск парсинга категорий...")
		categories, err := categoryWorker.Start()
		if err != nil {
			logger.Error("Ошибка при парсинге категорий", zap.Error(err))
			return
		}

		// Логирование найденных категорий
		for _, category := range categories {
			logger.Info("Категория", zap.String("Name", category.Name), zap.String("Href", category.Href))
		}
	default:
		log.Printf("Неизвестный тип парсинга: %d\n", parseType)
	}
}
