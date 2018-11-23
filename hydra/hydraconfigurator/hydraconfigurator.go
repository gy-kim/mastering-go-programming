package hydraconfigurator

import (
	"errors"
	"reflect"
)

const (
	CUSTOM uint8 = iota
	JSON
	XML
)

var wrongTypeError error = errors.New("Type must be a pointer to a struct")

func GetConfiguration(confType uint8, obj interface{}, filename string) (err error) {
	// check if this is type pointer
	mysRValue := reflect.ValueOf(obj)
	if mysRValue.Kind() != reflect.Ptr || mysRValue.IsNil() {
		return wrongTypeError
	}

	// get and confirm the struct value
	mysRValue = mysRValue.Elem()
	// *object => object
	// reflection value of *object .Elem() => boject() (Settable!!)

	if mysRValue.Kind() != reflect.Struct {
		return wrongTypeError
	}

	switch confType {
	case CUSTOM:
		err = MarchalCustomConfig(mysRValue, filename)
	case JSON:
		err = decodeJSONConfig(obj, filename)
	case XML:
		err = decodeXMLConfig(obj, filename)
	}

	return err
}
