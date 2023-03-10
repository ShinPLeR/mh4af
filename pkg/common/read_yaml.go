package common

import (
	"gopkg.in/yaml.v2"
	"os"
)

func ReadYaml[T interface{}](model *T, filepath string) error {
	bin, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bin, &model)
	if err != nil {
		return err
	}

	return nil
}
