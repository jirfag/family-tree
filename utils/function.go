package utils

import (
	"reflect"
)

// Contains is a func to check whether item in a array
func Contains(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {

			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}
