package gobpmn_builder_test

import (
	"reflect"
	"testing"
)

// TestReflectQuantities ...
func TestReflectQuantities(t *testing.T) {

	type Quantities struct {
		Process int
	}

	a := Quantities{Process: 1}

	r := reflect.ValueOf(&a).Elem()
	r1 := r.FieldByName("Process").Int()

	t.Logf("result of r is %+v", r1)

}
