package api

import (
	"github.com/Kcih4518/simpleBank/util"
	"github.com/go-playground/validator/v10"
)

// FieldLevel contains all the information and helper functions to validate fields
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	// check	currency is string
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
