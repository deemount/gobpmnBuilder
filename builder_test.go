package gobpmn_builder

import (
	"encoding/xml"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/deemount/gobpmnModels/pkg/core"
)

func TestReflectCounter(t *testing.T) {

	type A struct{ Process int }

	counter := A{
		Process: 1,
	}

	r := reflect.ValueOf(&counter)
	r1 := r.FieldByName("Process").Int()

	t.Logf("result of r is %v", r1)

}

func TestToBPMN(t *testing.T) {
	var err error
	// create a new repository
	repo := core.NewDefinitions()
	repo.SetDefaultAttributes()
	repo.SetID("definitions", "1234")
	repo.SetMainElements(1)
	proc := repo.GetProcess(0)
	t.Logf("result of proc is %v", proc)
	// marshal xml to byte slice
	b, err := xml.MarshalIndent(&repo, " ", "  ")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	// path to temporary bpmn files
	path := "temp"
	// create .bpmn file
	f, err := os.Create(path + "/test.bpmn")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	defer f.Close()
	// add xml header
	w := []byte(fmt.Sprintf("%v", xml.Header+string(b)))
	// write bytes to file
	_, err = f.Write(w)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	err = f.Sync()
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

}
