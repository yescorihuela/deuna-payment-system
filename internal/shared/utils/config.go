package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	HTTPServicePaymentPort       string `mapstructure:"HTTP_SERVICE_PAYMENT_PORT"`
	HTTPServiceAcquiringBankPort string `mapstructure:"HTTP_SERVICE_ACQUIRING_BANK"`
	DeunaDbDsn                   string `mapstructure:"DEUNA_DB_DSN"`
	HostAcquiringBank            string `mapstructure:"HOST_ACQUIRING_BANK"`
	TimeoutHTTPRequests          string `mapstructure:"TIMEOUT_HTTP_REQUESTS"`
	PathToMigrations             string `mapstructure:"PATH_TO_MIGRATIONS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
