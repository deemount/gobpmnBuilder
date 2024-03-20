package gobpmn_builder

import (
	"encoding/xml"
	"fmt"
	"os"
)

// ToBPMN marshals the repository to a bpmn file
// and writes the file to the file system.
// The method returns an error if the file could not be created.
// The method returns nil if the file was created successfully.
func (bldr *Builder) ToBPMN() error {

	var err error

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(&bldr.Repo, " ", "  ")

	// create .bpmn file
	currFile := bldr.GetCurrentlyCreatedFilename()
	f, err := os.Create(bldr.FilePathBPMN + "/" + currFile + ".bpmn")
	if err != nil {
		return err
	}
	defer f.Close()

	// add xml header
	w := []byte(fmt.Sprintf("%v", xml.Header+string(b)))

	// write bytes to file
	_, err = f.Write(w)
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}

	return nil

}
