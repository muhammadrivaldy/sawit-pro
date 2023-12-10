package configs

// Configuration is a object configuration
type Configuration struct {
	Port     int    `json:"port" env:"PORT"`
	JWTKey   string `json:"jwt-key" env:"JWT_KEY"`
	Database struct {
		User     string `json:"user" env:"DB_USER"`
		Password string `json:"password" env:"DB_PASSWORD"`
		Address  string `json:"address" env:"DB_ADDRESS"`
		Port     int    `json:"port" env:"DB_PORT"`
		Schema   string `json:"schema" env:"DB_SCHEMA"`
	} `json:"database"`
	ThirdParty struct {
		Telegram struct {
			Token  string `json:"token" env:"TELEGRAM_TOKEN"`
			ChatId int64  `json:"chat_id" env:"TELEGRAM_CHAT_ID"`
		} `json:"telegram"`
	} `json:"third_party"`
}
