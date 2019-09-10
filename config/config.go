package config

import (
	"errors"
	log "github.com/golang/glog"
	"github.com/spf13/viper"
)

func Initialize() error {
	viper.SetEnvPrefix("k8comp")
	viper.AddConfigPath("/opt/k8comp")
	viper.AddConfigPath(".")
	viper.SetConfigName("k8comp")

	err := viper.ReadInConfig()

	if err != nil {
		_, ok := err.(viper.ConfigParseError)
		if ok {
			return errors.New("Error parsing config")
		}
		log.V(1).Info("No config file used")
	} else {
		log.V(1).Infof("Using config file: %v", viper.ConfigFileUsed())
	}

	return nil
}
