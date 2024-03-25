package gobpmn_builder_test

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	"github.com/deemount/gobpmnModels/pkg/core"
)

func TestToBPMN(t *testing.T) {
	t.Run("TestToBPMN()",
		func(t *testing.T) {
			var err error

			// create a new repository
			repo := core.NewDefinitions()
			repo.SetDefaultAttributes()
			repo.SetID("definitions", "1234")
			repo.SetMainElements(1)

			// marshal xml to byte slice
			b, err := xml.MarshalIndent(&repo, " ", "  ")
			if err != nil {
				t.Errorf("expected nil, got %v", err)
			}

			// create .bpmn file
			f, err := os.Create(DefaultFiletestPath + "/test.bpmn")
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
		})
}
