package gobpmn_builder_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// Defaults
	DefaultFiletestPath     = "temp"
	DefaultFiletestNameBPMN = "test.bpmn"
	DefaultFiletestNameJSON = "test.json"
)

type TestQuantities struct {
	Process int
}

// TestReflectQuantities ...
func TestReflectQuantities(t *testing.T) {
	t.Run("TestReflectQuantities()",
		func(t *testing.T) {
			want := 1
			q := TestQuantities{Process: 1}
			r := reflect.ValueOf(&q).Elem()
			r1 := r.FieldByName("Process").Int()
			assert.Equal(t, want, int(r1), "want %v, got %v", want, r1)
		})
}
