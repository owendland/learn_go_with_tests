package collection

import (
	"reflect"
)

// Collect transforms src into Collection. The src could be json string, []string,
// []map[string]interface{}, map[string]interface{}, []int, []int16, []int32, []int64,
// []float32, []float64, []interface{}.
func Collect(src interface{}) *Collection {
	vMap := make(map[reflect.Value]reflect.Value)
	v, err := src.(reflect.Value)
	if !err {
		v = reflect.ValueOf(src)
	}
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		length := v.Len()
		for i := 0; i < length; i++ {
			vMap[reflect.ValueOf(i)] = v.Index(i)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			vMap[key] = v.MapIndex(key)
		}
	default:
	}

	return &Collection{vMap, v.Kind()}
}

type Collection struct {
	value map[reflect.Value]reflect.Value
	kind  reflect.Kind
}

// Every may be used to verify that all elements of a collection pass a given truth test.
func (c *Collection) Every(cb Callback) bool {
	for key, value := range c.value {
		if !cb(key, value) {
			return false
		}
	}
	return true
}

type Callback func(item reflect.Value, value reflect.Value) bool
