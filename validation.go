package validation

import "github.com/go-playground/validator/v10"

var (
	Validate *validator.Validate
)

func init() {
	Validate = validator.New()
}

func ValidateByStr(field interface{}, tag string) error {
	err := Validate.Var(field, tag)
	if err != nil {
		return err
	}
	return nil
}
