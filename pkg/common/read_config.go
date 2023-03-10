package common

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfig[T interface{}](model *T) error {
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Cound't read configs: %s \n", err)
	}

	if err := viper.Unmarshal(&model); err != nil {
		return fmt.Errorf("Cound't recognize model: %s \n", err)
	}

	return nil
}
