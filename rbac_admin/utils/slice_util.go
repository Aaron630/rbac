package utils

import (
	"fmt"
	"reflect"
)

// Contain func
func Contain(e interface{}, target interface{}, key string) bool {
	targetValue := reflect.ValueOf(target)
	targetType := reflect.TypeOf(target).Kind()
	fmt.Println(targetValue)
	fmt.Println(targetValue.Len())
	switch targetType {
	case reflect.Slice, reflect.Array:
		if key != "" {
			for i := 0; i < targetValue.Len(); i++ {
				// fmt.Println(targetValue.Index(i).FieldByName(key))
				// fmt.Println(reflect.TypeOf(targetValue.Index(i).FieldByName(key)))
				// fmt.Println(reflect.TypeOf(e))
				if targetValue.Index(i).FieldByName(key).Interface() == e {
					return true
				}
			}
		} else {
			for i := 0; i < targetValue.Len(); i++ {
				if targetValue.Index(i).Interface() == e {
					return true
				}
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(e)).IsValid() {
			return true
		}
	}

	return false
}
