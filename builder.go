package gobpmnBuilder

import (
	"fmt"
	"log"
	"os"

	"github.com/deemount/gobpmnModels/pkg/core"
)

var (
	options Options
	path    string
)

type (

	// Option ...
	Option func(o Options) Options

	// BuilderRepository ...
	BuilderRepository interface {
		Build() (Builder, error)
		SetDefinitions()
		SetDefinitionsByArg(r core.DefinitionsRepository)
		ToBPMN() error
		GetCurrentlyCreatedFilename() string
	}

	// Builder ...
	Builder struct {
		Options Options
	}
)

// NewBuilder ...
func NewBuilder(option ...Option) BuilderRepository {
	//
	for _, o := range option {
		options = o(options)
	}

	// path to bpmn files
	path = "files/bpmn"
	// read the dir for created bpmn files
	log.Printf("goBuilder.builder.go:NewBuilder path to bpmn files %s", path)
	files, err := os.ReadDir("files/bpmn")
	if err != nil {
		log.Panic(err)
	}

	// set number and count up for each created file
	if options.Counter == 0 {
		options.Counter = len(files)
	} else {
		options.Counter += 1
	}

	// set default name for bpmn-file
	options.CurrentFile = "diagram_" + fmt.Sprintf("%d", options.Counter)
	log.Printf("goBuilder.builder.go:NewBuilder created filename %s.bpmn", options.CurrentFile)

	return &Builder{Options: options}
}

// SetDefinitions ...
func (bldr *Builder) SetDefinitions() {
	bldr.Options.Repo = core.NewDefinitions()
}

// SetDefinitionsByArg ...
func (bldr *Builder) SetDefinitionsByArg(r core.DefinitionsRepository) {
	bldr.Options.Repo = r
}

// GetCurrentlyCreatedFilename ...
func (bldr Builder) GetCurrentlyCreatedFilename() string {
	return bldr.Options.CurrentFile
}

// Build...
func (bldr Builder) Build() (Builder, error) {

	var err error

	bldr.SetDefinitions()

	// create .bpmn
	log.Println("goBuilder.builder.go:Build create bpmn")
	err = bldr.ToBPMN()
	if err != nil {
		return Builder{}, err
	}

	return bldr, nil

}
