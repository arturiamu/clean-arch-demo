package config

var DefaultConf Config

type Config struct {
	App struct {
		Name      string `mapstructure:"name"`
		Port      int    `mapstructure:"port"`
		SecretKey string `mapstructure:"secret_key"`
	} `mapstructure:"app"`

	Auth struct {
		HashSalt   string `mapstructure:"hash_salt"`
		SigningKey string `mapstructure:"signing_key"`
		TokenTTL   int    `mapstructure:"token_ttl"`
	} `mapstructure:"auth"`

	Mysql struct {
		Name     string `mapstructure:"name"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"mysql"`

	Logging struct {
		Level    string `mapstructure:"level"`
		FilePath string `mapstructure:"file_path"`
	}
}
