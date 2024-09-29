package config

import (
	"strings"

	"github.com/spf13/viper"
)

// New - функция для инициализации Viper и загрузки конфигурации из файла
func New(cfgPath string) (Wrapper, error) {
	v := viper.New()

	// Настройка для считывания переменных окружения с префиксом "ENV"
	v.AutomaticEnv()
	v.SetEnvPrefix("ENV")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Указание файла конфигурации и его типа (yaml)
	v.SetConfigFile(cfgPath)
	v.SetConfigType("yaml")

	// Чтение конфигурации из файла
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return &wrapper{viper: v}, nil
}

// Wrapper - интерфейс для доступа к конфигурационным параметрам
type Wrapper interface {
	GetBool(key string) bool
	GetString(key string) string
	GetInt64(key string) int64
	IsSet(key string) bool
	UnmarshalKey(key string, rawVal interface{}) error
}

// wrapper - реализация интерфейса Wrapper на основе Viper
type wrapper struct {
	viper *viper.Viper
}

func (w *wrapper) GetBool(key string) bool {
	return w.viper.GetBool(key)
}

func (w *wrapper) GetString(key string) string {
	return w.viper.GetString(key)
}

func (w *wrapper) GetInt64(key string) int64 {
	return w.viper.GetInt64(key)
}

func (w *wrapper) IsSet(key string) bool {
	return w.viper.IsSet(key)
}

func (w *wrapper) UnmarshalKey(key string, rawVal interface{}) error {
	return w.viper.UnmarshalKey(key, rawVal)
}
