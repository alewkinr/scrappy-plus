package internal

import "github.com/kelseyhightower/envconfig"

// Settings — объект настроек
type Settings struct {
	// ParseURL — хост для парсинга
	ParseURL string `envconfig:"PARSE_URL"`
	// TelegramToken – токен для телеграма
	TelegramToken string `envconfig:"TELEGRAM_TOKEN"`
	// TelegramChannelID — ID канала для телеграма
	TelegramChannelID string `envconfig:"TELEGRAM_CHANNEL_ID"`
}

// initConfig возвращает конфиг
func initConfig() (*Settings, error) {
	var cfg Settings
	err := envconfig.Process("", &cfg)

	return &cfg, err
}

// MustInitConfig возвращает конфиг или паникует при ошибке
func MustInitConfig() *Settings {
	cfg, err := initConfig()
	if err != nil {
		panic(err)
	}

	return cfg
}
