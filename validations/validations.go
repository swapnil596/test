package validations

// happily copied and updated from
// https://blog.depa.do/post/gin-validation-errors-handling

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type JSONFormatter struct{}

// NewJSONFormatter will create a new JSON formatter and register a custom tag
// name func to gin's validator
func NewJSONFormatter() *JSONFormatter {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	return &JSONFormatter{}
}

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func Descriptive(verr validator.ValidationErrors) []ValidationError {
	var errs []ValidationError

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationError{Field: f.Field(), Reason: err})
	}

	return errs
}

func (JSONFormatter) Simple(ver validator.ValidationErrors) map[string]string {

	errs := make(map[string]string)

	// ranging over the field of errors and for each of them we're
	// setting the field name as the map key and the tag that matched as the value.
	for _, f := range ver {

		// f.ActualTag returns the tag that triggered the failure
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}

		// f.Field returns the field that failed
		errs[f.Field()] = err
	}

	// returning a map[string]string because that’s a type Go’s JSON library can deal with pretty easily.
	return errs
}

func ValidateErrors(err error) interface{} {
	// verr => validation errors
	var verr validator.ValidationErrors

	if errors.As(err, &verr) {
		return Descriptive(verr)
	}

	return err.Error()
}
