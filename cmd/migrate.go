package cmd

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var (
	migrationsPath string
	name           string

	// Определение команды migrate для миграций базы данных
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Управление миграциями базы данных",
	}

	// Валидация аргументов для создания миграции
	createArgsValidator = func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Нужно указать один аргумент - имя миграции (строка)")
		}
		if args[0] == "" {
			return fmt.Errorf("Имя миграции не может быть пустым")
		}
		return nil
	}

	// Команда для применения всех миграций
	migrateUpCmd = &cobra.Command{
		Use:   "up",
		Short: "Применить все миграции",
		Long:  "Обновить базу данных до последней версии",
		RunE:  migrateUpCmdHandler,
	}

	// Команда для отката всех миграций
	migrateDownCmd = &cobra.Command{
		Use:   "down",
		Short: "Откатить все миграции",
		Long:  "Откатить все изменения базы данных",
		RunE:  migrateDownCmdHandler,
	}

	// Команда для отката одной миграции
	migrateDownByOneCmd = &cobra.Command{
		Use:   "down-by-one",
		Short: "Откатить одну транзакцию",
		Long:  "Откатить базу данных на одну версию",
		RunE:  migrateDownByOneCmdHandler,
	}

	// Команда для создания новой миграции
	migrateCreateCmd = &cobra.Command{
		Use:   "create [migration_name]",
		Short: "Создать миграцию",
		Long:  "Создать новый файл миграции с текущей временной меткой",
		Args:  createArgsValidator,
		RunE:  migrateCreateCmdHandler,
	}
)

// Инициализация команд миграции
func init() {
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	migrateCmd.AddCommand(migrateDownByOneCmd)
	migrateCmd.AddCommand(migrateCreateCmd)
	rootCmd.AddCommand(migrateCmd)

	// Флаги команд
	migrateCmd.PersistentFlags().StringVarP(&migrationsPath, "migrationsPath", "m", "migrations", "Путь к файлам миграции")
	migrateCmd.PersistentFlags().StringVarP(&name, "name", "n", "migrate1", "Имя файла миграции")
}

// Обработчик команды для применения всех миграций
func migrateUpCmdHandler(*cobra.Command, []string) error {
	var db *sql.DB
	var embedMigrations embed.FS

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, migrationsPath); err != nil {
		panic(err)
	}
	return nil
}

// Обработчик команды для отката всех миграций
func migrateDownCmdHandler(*cobra.Command, []string) error {
	var db *sql.DB
	var embedMigrations embed.FS

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Down(db, migrationsPath); err != nil {
		panic(err)
	}
	return nil
}

// Обработчик команды для отката одной миграции
func migrateDownByOneCmdHandler(*cobra.Command, []string) error {
	var db *sql.DB
	var embedMigrations embed.FS

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Down(db, migrationsPath); err != nil {
		panic(err)
	}
	return nil
}

// Обработчик команды для создания новой миграции
func migrateCreateCmdHandler(_ *cobra.Command, args []string) error {
	if name == "" {
		return fmt.Errorf("Не указано имя миграции")
	}

	var err error
	if err = os.MkdirAll(migrationsPath, 0755); err != nil {
		return err
	}

	// Шаблон SQL для создания файла миграции
	var sqlMigrationTemplate = template.Must(
		template.New("goose.sql-migration").
			Parse(
				`-- +goose Up 
-- SQL в этом разделе выполняется при применении миграции
      
-- +goose Down 
-- SQL в этом разделе выполняется при откате миграции
`))

	if err = goose.CreateWithTemplate(nil, migrationsPath, sqlMigrationTemplate, name, "sql"); err != nil {
		return err
	}
	return nil
}
