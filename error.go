package govalidator

import (
	"reflect"
	"sort"
	"strings"
)

// Errors is an array of multiple errors and conforms to the error interface.
type Errors []error

// Errors returns itself.
func (es Errors) Errors() []error {
	return es
}

func (es Errors) Error() string {
	var errs []string
	for _, e := range es {
		errs = append(errs, e.Error())
	}
	sort.Strings(errs)
	return strings.Join(errs, ";")
}

type ErrOpt func(err *Error)

func WithPath(path []string) ErrOpt {
	return func(err *Error) {
		err.Path = path
	}
}

func WithCustomErrMessage(customMsgExists bool) ErrOpt {
	return func(err *Error) {
		err.CustomErrorMessageExists = customMsgExists
	}
}

func WithTag(tag reflect.StructTag) ErrOpt {
	return func(err *Error) {
		jsonTag := strings.Split(tag.Get("json"), ",")[0]
		if len(jsonTag) > 0 {
			err.JsonKey = &jsonTag
		}
	}
}

func NewError(name string, err error, validator string, value *string, opts ...ErrOpt) Error {
	newErr := &Error{
		Name:                     name,
		Err:                      err,
		CustomErrorMessageExists: false,
		Validator:                validator,
		Path:                     []string{},
		JsonKey:                  nil,
		Value:                    value,
	}
	for _, opt := range opts {
		opt(newErr)
	}
	return *newErr
}

// Error encapsulates a name, an error and whether there's a custom error message or not.
type Error struct {
	Name                     string
	Err                      error
	CustomErrorMessageExists bool

	// Validator indicates the name of the validator that failed
	Validator string
	Path      []string
	JsonKey   *string
	Value     *string
}

func (e Error) Error() string {
	if e.CustomErrorMessageExists {
		return e.Err.Error()
	}

	errName := e.Name
	if len(e.Path) > 0 {
		errName = strings.Join(append(e.Path, e.Name), ".")
	}

	return errName + ": " + e.Err.Error()
}
