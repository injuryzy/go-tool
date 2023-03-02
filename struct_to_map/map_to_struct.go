package struct_to_map

import "reflect"

func StructToMap(v interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	of := reflect.ValueOf(v)
	typeOf := reflect.TypeOf(v)
	if of.Kind() != reflect.Struct {
		panic(" input interface not is struct ")
	}
	for i := 0; i < of.NumField(); i++ {
		switch of.Field(i).Kind() {
		case reflect.Int, reflect.Int16, reflect.Int64:
			m[typeOf.Field(i).Name] = of.Field(i).Int()
		case reflect.String:
			m[typeOf.Field(i).Name] = of.Field(i).String()
		}
	}
	return m
}
