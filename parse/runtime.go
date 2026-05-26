package parse

import (
	"reflect"
)

func objectGoType(t reflect.Type, structT reflect.Type) string {
	_ = "STUB: not implemented"
	return ""
}

// drop package name from qualified identifier if type is defined in the same package

// Object extracts struct information from given object.
func Object(obj interface{}, schema, table string) (res *StructInfo, err error) {
	_ = "STUB: not implemented"
	// convert any panic to error
	return nil, nil
}

// nothing

// check for anonymous fields

// check for exported name

// parse tag and type
