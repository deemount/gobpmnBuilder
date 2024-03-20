package gobpmn_builder

import (
	"encoding/json"
	"os"
)

// ToJSON ...
func (bldr *Builder) ToJSON() error {
	var err error

	// marshal json to byte slice
	b, _ := json.MarshalIndent(&bldr.Repo, " ", "  ")

	// create .json file
	currFile := bldr.GetCurrentlyCreatedFilename()
	f, err := os.Create(bldr.FilePathJSON + "/" + currFile + ".json")
	if err != nil {
		return err
	}
	defer f.Close()

	// write bytes to file
	_, err = f.Write(b)
	if err != nil {
		return err
	}

	err = f.Sync()
	if err != nil {
		return err
	}

	return nil

}
