package util

import (
	"reflect"
)

func IterateStruct(stru interface{}, yield func(key string, value any)) {
	v := reflect.ValueOf(stru)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		yield(typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}
