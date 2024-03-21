package gobpmn_builder_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/deemount/gobpmnModels/pkg/core"
)

func TestToJSON(t *testing.T) {
	var err error

	// create a new repository
	repo := core.NewDefinitions()
	repo.SetDefaultAttributes()
	repo.SetID("definitions", "1234")
	repo.SetMainElements(1)

	// marshal xml to byte slice
	b, err := json.MarshalIndent(&repo, " ", "  ")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	// create .bpmn file
	f, err := os.Create(DefaultFiletestPath + "/test.json")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	defer f.Close()

	// write bytes to file
	_, err = f.Write(b)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	err = f.Sync()
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

}
