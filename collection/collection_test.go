package collection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCollection(t *testing.T) {
	a := []int{2, 3, 4, 5, 6, 7}

	result := Collect(a).Every(func(item reflect.Value, value reflect.Value) bool {
		return true
	})

	fmt.Println(result)
}
