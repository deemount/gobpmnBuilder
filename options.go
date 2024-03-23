package gobpmn_builder

import (
	"os"

	"github.com/deemount/gobpmnModels/pkg/core"
)

/*
  Functional Options Pattern Type and Functions for the Builder
*/

// Option is a functional option type for the Builder.
type Option = func(bldr *Builder) error

// WithPath sets the path to the bpmn and json files.
// Ruleset:
//   - If the length of path is empty, the default path will be used.
//   - If the length of path is 1, the path to the bpmn files will be set.
//   - If the length of path is 2, the path to the bpmn and json files will be set.
//   - If the length of path is greater than 2, the path to the bpmn and json files will be set
//     and the rest of the paths will be ignored.
//
// The first path is the path to the bpmn files and the second path is the path to the json files.
// Note:
// The function uses variadic parameters. Parameters can be empty or more than one.
func WithPath(path ...string) Option {
	return func(bldr *Builder) error {

		length := len(path)

		switch true {
		case length == 0:
			bldr.FilePathBPMN = DefaultPathBPMN
			bldr.FilePathJSON = DefaultPathJSON

		case length == 1:
			if _, err := os.Stat(path[0]); os.IsNotExist(err) {
				return ErrPathNotFound
			}
			bldr.FilePathBPMN = path[0]
			bldr.FilePathJSON = DefaultPathJSON

		case length >= 2:
			if _, err := os.Stat(path[0]); os.IsNotExist(err) {
				return ErrPathNotFound
			}
			bldr.FilePathBPMN = path[0]
			bldr.FilePathJSON = path[1]

		}

		return nil

	}
}

// WithCounter reads the directory with the path to the specified files
// and sets the number and count up for each created file.
// Ruleset:
//   - If the path is empty, the default path will be used.
//   - If the path is not empty AND isn't the same like the default path, the path argument will be used.
//   - Only the first path will be used. All other paths will be ignored.
//
// The number of files in this directory will be used for all other file types.
// Note:
// The function uses variadic parameters. Parameters can be empty or more than one.
func WithCounter(path ...string) Option {
	return func(bldr *Builder) error {

		var filepath string

		// check if variadic parameter has a length of 1 and is not the default path
		if len(path) == 1 && path[0] != bldr.FilePathBPMN {
			filepath = path[0]
		} else {
			filepath = bldr.FilePathBPMN
		}

		// reading the directory with the path to the specified files
		files, err := os.ReadDir(filepath)
		if err != nil {
			return err
		}

		// get the length of the files
		length := len(files)

		// set number and count up for each created file
		if length == 0 {
			bldr.Counter = 1
		} else {
			bldr.Counter = length + 1
		}

		return nil
	}
}

// WithFilenamePrefix returns the filename prefix, which is used to create the file name.
func WithFilenamePrefix(filenamePrefix string) Option {
	return func(bldr *Builder) error {
		bldr.FilenamePrefix = filenamePrefix
		return nil
	}
}

// WithDefinitions stores the definitions repository in the builder.
func WithDefinitions(Repo core.DefinitionsRepository) Option {
	return func(bldr *Builder) error {
		bldr.Repo = Repo
		return nil
	}
}
