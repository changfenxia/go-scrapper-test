package main

import (
	"log"

	"github.com/changfenxia/scrapper-test/cmd" // Укажите правильный путь к пакету cmd
)

func init() {
	// Настройка логирования с отображением даты, времени и короткого пути файла
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile | log.LUTC)
}

func main() {
	// Запуск корневой команды CLI
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
