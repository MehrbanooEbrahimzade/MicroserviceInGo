package configs

import (
	"fmt"
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/pkg/models"

	"github.com/spf13/viper"
)

func NewConfig() *models.Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath("./pkg/configs")
	err := v.ReadInConfig()
	if err != nil {
		panic("CANNOT PARSE CONFIGS " + err.Error())
	}
	return &models.Config{
		Cors: models.Cors{
			AllowedOrigins: v.GetStringSlice("allowed-origins"),
			AllowedMethods: v.GetStringSlice("allowed-methods"),
			AllowedHeaders: v.GetStringSlice("allowed-headers"),
			ExposedHeaders: v.GetStringSlice("exposed-headers"),
			MaxAge:         v.GetInt("max-age"),
		},
		UsersPort: fmt.Sprintf(":%d", v.GetInt("users-port")),
		Port:      v.GetInt("port"),
	}
}
