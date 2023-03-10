package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"mh4af/internal/model"
)

var validate *validator.Validate

func Validate(d model.Definition) (*model.RearrangedDefinition, error) {
	err := validate.Struct(d)
	if err != nil {
		fmt.Println("Load Error")
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.ActualTag())
		}
		return nil, err
	}

	arranged, err := d.Rearrange()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &arranged, nil
}

func init() {
	validate = validator.New()
}
