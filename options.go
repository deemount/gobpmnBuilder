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

// WithCounter reads the directory with the path to the specified files
// and sets the number and count up for each created file. If the path is empty,
// the default path will be used.
func WithCounter(path ...string) Option {
	return func(bldr *Builder) {
		var filepath string
		// check if the path is empty
		if len(path) > 0 {
			filepath = path[0]
		} else {
			filepath = bldr.FilePathBPMN
		}
		// reading the directory with the path to the specified files
		files, err := os.ReadDir(filepath)
		if err != nil {
			log.Panic(err)
		}
		// set number and count up for each created file
		if len(files) == 0 {
			bldr.Counter = 1
		} else {
			bldr.Counter = bldr.Counter + 1
		}
	}
}

// WithFilenamePrefix returns the filename prefix, which is used to create the file name.
func WithFilenamePrefix(filenamePrefix string) Option {
	return func(bldr *Builder) {
		bldr.FilenamePrefix = filenamePrefix
	}
}

// WithDefinitions stores the definitions repository in the builder.
func WithDefinitions(Repo core.DefinitionsRepository) Option {
	return func(bldr *Builder) {
		bldr.Repo = Repo
	}
}
