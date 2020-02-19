package env

import (
	"fmt"

	"github.com/spf13/viper"
)

type envConfigStruct struct {
	Port string `mapstructure:"port" json:"port,omitempty"`
	Env  string `mapstructure:"env" json:"env,omitempty"`
}

// EnvConfig ...
var EnvConfig envConfigStruct

func init() {
	viper.AutomaticEnv()

	viper.SetDefault("ENV", "local")
	if env := viper.GetString("ENV"); env == "local" {
		viper.SetConfigName("local")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(fmt.Sprintf("local.env read error %s", err.Error()))
		}
	}

	if err := viper.Unmarshal(&EnvConfig); err != nil {
		fmt.Println(fmt.Sprintf("Unmarshal into EnvConfig error: %s", err.Error()))
	}
}
