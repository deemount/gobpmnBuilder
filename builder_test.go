package gobpmn_builder_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/deemount/gobpmnModels/pkg/core"
)

// TestReflectQuantities
func TestReflectQuantities(t *testing.T) {

	type Quantities struct {
		Process int
	}

	a := Quantities{Process: 1}

	r := reflect.ValueOf(&a).Elem()
	r1 := r.FieldByName("Process").Int()

	t.Logf("result of r is %+v", r1)

}

type mockOptions struct {
	mock.Mock
	Counter      int                        // Counter is the number of created files
	CurrentFile  string                     // CurrentFile is the name of the current file
	ModelType    string                     // ModelType is the type of the model (can be human or technical)
	FilePathBPMN string                     // FilePathBPMN is the path to the bpmn files
	FilePathJSON string                     // FilePathJSON is the path to the json files
	Def          *core.Definitions          // Def is the definition of the model
	Repo         core.DefinitionsRepository // Repo is the repository of the model
}

type mockOption func(opt mockOptions) mockOptions

func TestSetOptions(t *testing.T) {

	opt := mockOptions{}
	t.Logf("result of o is %+v", opt)

}
