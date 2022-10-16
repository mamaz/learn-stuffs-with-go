package maphelper

import (
	"object-creation/user"
	"reflect"
)

func ToMap[T *user.UserCreated | user.User](anyObj T) map[string]any {
	value := reflect.ValueOf(anyObj)
	indirectValue := reflect.Indirect(value)
	mapResult := map[string]any{}

	for index := 0; index < indirectValue.NumField(); index++ {
		fieldName := indirectValue.Type().Field(index).Name
		val := indirectValue.Field(index)
		mapResult[fieldName] = val.Interface()
	}

	return mapResult
}
