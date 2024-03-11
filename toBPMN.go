package gobpmn_builder

import (
	"encoding/xml"
	"fmt"
	"os"
)

// toBPMN ...
func (bldr *Builder) ToBPMN() error {

	var err error

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(&bldr.Options.Repo, " ", "  ")

	// path to bpmn files
	path := "files/bpmn"
	// create .bpmn file
	f, err := os.Create(path + "/" + bldr.Options.CurrentFile + ".bpmn")
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

	// create .json
	err = bldr.toJSON()
	if err != nil {
		return err
	}

	return nil

}
