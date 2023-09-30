package api

import (
	"github.com/gistGitUser/course/project/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(level validator.FieldLevel) bool {
	if currency, ok := level.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false

}
