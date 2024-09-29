package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	cfgPath string // Путь к файлу конфигурации

	// Канал для уведомления об остановке приложения
	stopNotification = make(chan struct{})

	// Определение корневой команды CLI
	rootCmd = &cobra.Command{
		Use:           "scraper [command]",
		Long:          "eda.ru scraper service",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Обработка системных сигналов для корректного завершения приложения
			go func() {
				c := make(chan os.Signal, 1)
				signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
				<-c
				stopNotification <- struct{}{}
			}()
			return nil
		},
	}
)

// Execute запускает корневую команду и все вложенные
func Execute() error {
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "config.yaml", "Путь к файлу конфигурации")
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
