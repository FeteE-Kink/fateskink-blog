package formAttributes

import (
	"fmt"
	"server/app/pkg/helpers"
	"time"
)

type FieldAttribute struct {
	Name         string
	Code         string
	Errors       []interface{}
	DefaultValue interface{}
}

type FieldAttributeInterface interface {
	AddError(message interface{})
	GetCode() string
	GetFieldI18nCode() string
	GetErrors() []interface{}
	Time() *time.Time
	IsClean() bool
	Present() bool
	Blank() bool

	// Validators
	ValidateRequired()
	ValidateMin(min interface{})
	ValidateMax(max interface{})
	ValidateIsPowerOf(value interface{})
	ValidateFormat(formatter string, formateterRemind string)
	ValidateGt(value interface{}, message *string)
	ValidateGtEq(value interface{}, message *string)
	ValidateLt(value interface{}, message *string)
	ValidateLtEq(value interface{}, message *string)
}

func ValidateionMesage(column string, message interface{}) interface{} {
	return fmt.Sprintf("%s%s", helpers.CamelToPascalCase(column), message)
}
