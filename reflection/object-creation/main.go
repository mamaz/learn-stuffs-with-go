package main

import (
	"fmt"
	m "object-creation/map-helper"
	u "object-creation/user"
	"reflect"
)

func main() {
	var usr *u.UserCreated
	created := generate(usr)

	fmt.Printf("%+v\n", created)
}

func generate[T *u.UserCreated](usrCreated T) T {
	usrType := reflect.TypeOf(usrCreated).Elem()
	usrValue := reflect.New(usrType)
	indirectUsrValue := reflect.Indirect(usrValue)

	userMap := m.ToMap[T](&u.UserCreated{
		ID:   "123",
		Type: "UserCreated",
	})

	for key, value := range userMap {
		field := indirectUsrValue.FieldByName(key)
		field.Set(reflect.ValueOf(value))
	}

	return usrValue.Interface().(T)
}
