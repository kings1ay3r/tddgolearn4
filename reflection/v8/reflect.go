package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val, kind := getValue(x)
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	/* if field.Kind() == reflect.String {
		fn(field.String())
	} */

	switch kind {
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			walkValue(field)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			field := val.Index(i)
			walkValue(field)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}
func getValue(x interface{}) (val reflect.Value, kind reflect.Kind) {

	val = reflect.ValueOf(x)
	kind = val.Kind()
	if kind == reflect.Ptr {
		val = val.Elem()
		return
	}
	return
}
