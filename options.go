package gobpmn_builder

import (
	"log"
	"os"

	"github.com/deemount/gobpmnModels/pkg/core"
)

/*
  Functional Options Pattern Type and Functions for the Builder
*/

// Option ...
type Option = func(bldr *Builder)

// WithPath ...
func WithPath(FilePathBPMN, FilePathJSON string) Option {
	return func(bldr *Builder) {
		bldr.FilePathBPMN = FilePathBPMN
		bldr.FilePathJSON = FilePathJSON
	}
}

// WithCounter ...
func WithCounter(filePath string) Option {
	return func(bldr *Builder) {
		// reading the directory with the path to the specified files
		files, err := os.ReadDir(filePath)
		if err != nil {
			log.Panic(err)
		}
		// set number and count up for each created file
		if len(files) == 0 {
			bldr.Counter = 1
		} else {
			bldr.Counter += 1
		}
	}
}

func WithCurrentFile(filenamePrefix string) Option {
	return func(bldr *Builder) {
		bldr.FilenamePrefix = filenamePrefix
	}
}

func WithDefinitions(Repo core.DefinitionsRepository) Option {
	return func(bldr *Builder) {
		bldr.Repo = Repo
	}
}
