package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val, kind := getValue(x)

	/* if field.Kind() == reflect.String {
		fn(field.String())
	} */

	switch kind {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			walk(field.Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			field := val.Index(i)
			walk(field.Interface(), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}
}
func getValue(x interface{}) (val reflect.Value, kind reflect.Kind) {

	val = reflect.ValueOf(x)
	kind = val.Kind()
	if kind == reflect.Ptr {
		val = val.Elem()
	}
	return
}
