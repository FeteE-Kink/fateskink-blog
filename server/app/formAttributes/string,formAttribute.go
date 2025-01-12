package formAttributes

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

type StringAttribute struct {
	FieldAttribute

	Value string
}

func (attr *StringAttribute) GetCode() string {
	return attr.Code
}

func (attr *StringAttribute) GetFieldI18nCode() string {
	if attr.Name != "" {
		return attr.Name
	}

	return attr.Code
}

func (attr *StringAttribute) GetErrors() []interface{} {
	return attr.Errors
}

func (attr *StringAttribute) AddError(message interface{}) {
	attr.Errors = append(attr.Errors, ValidateionMesage(attr.GetFieldI18nCode(), message))
}

func (attr *StringAttribute) ValidateRequired() {
	if attr.Blank() {
		attr.AddError("is required")
	}
}

func (attr *StringAttribute) ValidateFormat(formatter string, formatterRemind string) {
	if attr.Value != "" {
		re := regexp.MustCompile(formatter)

		if re.MatchString(attr.Value) {
			return
		}

		attr.AddError("format error")
	}
}

func (form *StringAttribute) Time() *time.Time {
	return nil
}

func (attr *StringAttribute) IsClean() bool {
	return len(attr.Errors) == 0
}

func (attr *StringAttribute) ValidateMin(min interface{}) {
	switch v := min.(type) {
	case int64:
		if int64(utf8.RuneCountInString(attr.Value)) < v {
			attr.AddError(fmt.Sprintf("length min: %v", min))
		}
	default:
		panic("Need to provice int64 interface{} as params")
	}
}

func (attr *StringAttribute) ValidateMax(max interface{}) {
	switch v := max.(type) {
	case int64:
		if v < int64(utf8.RuneCountInString(attr.Value)) {
			attr.AddError(fmt.Sprintf("length max: %v", max))
		}
	default:
		panic("Need to provice int64 interface{} as params")
	}
}

func (attr *StringAttribute) ValidateGt(min interface{}) {
	switch v := min.(type) {
	case int64:
		if int64(utf8.RuneCountInString(attr.Value)) <= v {
			attr.AddError(fmt.Sprintf("length min: %v", min))
		}
	default:
		panic("Need to provice int64 interface{} as params")
	}
}

func (attr *StringAttribute) ValidateGtEq(min interface{}) {
	switch v := min.(type) {
	case int64:
		if int64(utf8.RuneCountInString(attr.Value)) < v {
			attr.AddError(fmt.Sprintf("length min: %v", min))
		}
	default:
		panic("Need to provice int64 interface{} as params")
	}
}

func (attr *StringAttribute) ValidateLt(min interface{}) {
	switch v := min.(type) {
	case int64:
		if int64(utf8.RuneCountInString(attr.Value)) >= v {
			attr.AddError(fmt.Sprintf("length max: %v", min))
		}
	default:
		panic("Need to provice int64 interface{} as params")
	}
}

func (attr *StringAttribute) ValidateLtEq(min interface{}) {
	switch v := min.(type) {
	case int64:
		if int64(utf8.RuneCountInString(attr.Value)) > v {
			attr.AddError(fmt.Sprintf("length max: %v", min))
		}
	default:
		panic("Need to provice int64 interface{} as params")
	}
}

func (attr *StringAttribute) Blank() bool {
	return attr.Present()
}

func (attr *StringAttribute) Present() bool {
	return strings.TrimSpace(attr.Value) != ""
}
