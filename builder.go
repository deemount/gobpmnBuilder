package gobpmn_builder

import (
	"errors"
	"fmt"
	"reflect"

	gobpmn_count "github.com/deemount/gobpmnCounter"
	"github.com/deemount/gobpmnModels/pkg/core"
)

var (
	// Defaults
	DefaultCounter        = 1            // DefaultCounter is the default counter
	DefaultFilenamePrefix = "diagram"    // DefaultFilenamePrefix is the default filename prefix
	DefaultModelType      = "human"      // DefaultModelType is the default model type
	DefaultPathBPMN       = "files/bpmn" // DefaultPathBPMN is the default path to the bpmn files
	DefaultPathJSON       = "files/json" // DefaultPathJSON is the default path to the json files

	// Errors
	ErrPathNotFound      = errors.New("path not found")           // ErrPathNotFound is the error for the path not found
	ErrEmptyFilePathBPMN = errors.New("empty file path for bpmn") // ErrEmptyFilePathBPMN is the error for the empty file path for bpmn
)

type (

	// BuilderRepository is the interface for the builder repository.
	// The interface defines the methods for the builder repository.
	// The methods are used to set the definitions, set the definitions by argument,
	// set the defaults, convert the model to bpmn, convert the model to json,
	// get the currently created filename, and build the model.
	BuilderRepository interface {
		SetDefinitions()
		SetDefinitionsByArg(r core.DefinitionsRepository)
		Defaults(p interface{}, c *gobpmn_count.Quantities)
		ToBPMN() error // Sets the bpmn file
		ToJSON() error // Sets the json file
		GetCurrentlyCreatedFilename() string
		Build() (Builder, error)
	}

	// Builder is the main struct of the builder package and holds
	// the options and the current file. The options are the parameters
	// for the builder and the current file is the name of the current
	// bpmn file in the particular directory.
	Builder struct {
		Counter        int                        // Counter is the number of created files
		FilenamePrefix string                     // FilenamePrefix is a part of the name of the current file
		ModelType      string                     // ModelType is the type of the model
		FilePathBPMN   string                     // FilePathBPMN is the path to the bpmn files
		FilePathJSON   string                     // FilePathJSON is the path to the json files
		Def            *core.Definitions          // Def is the definition of the model
		Repo           core.DefinitionsRepository // Repo is the repository of the model
	}
)

// New initializes the builder with the given options
// and returns the builder repository.
// The method sets the default values and applies the options
// to the builder.
func New(opts ...Option) (BuilderRepository, error) {
	// Set the default values
	bldr := &Builder{
		Counter:        DefaultCounter,
		FilenamePrefix: DefaultFilenamePrefix,
		ModelType:      DefaultModelType,
		FilePathBPMN:   DefaultPathBPMN,
		FilePathJSON:   DefaultPathJSON,
		Def:            nil,
		Repo:           nil,
	}
	// Apply the options to the builder
	for _, opt := range opts {
		opt(bldr)
	}
	if err := bldr.validate(); err != nil {
		return nil, err
	}
	return bldr, nil
}

// SetDefinitions ...
func (bldr *Builder) SetDefinitions() {
	bldr.Repo = core.NewDefinitions()
}

// SetDefinitionsByArg ...
func (bldr *Builder) SetDefinitionsByArg(r core.DefinitionsRepository) {
	bldr.Repo = r
}

// Defaults receives the definitions repository by the app in p argument
// and calls the main elements to set the maps, including process parameters
// n of process. The method contains the reflected process definition (p interface{})
// and calls it by the reflected method name.
//
// Note:
// This method hides specific setters (SetProcess, SetCollaboration, SetDiagram).
func (bldr *Builder) Defaults(p interface{}, c *gobpmn_count.Quantities) {

	// el is the interface {} of a given definition
	el := reflect.ValueOf(&p).Elem()
	counter := reflect.ValueOf(&c).Elem()

	// Get the number of processes
	numParticipants := counter.Elem().FieldByName("Participant").Int()
	numProcess := counter.Elem().FieldByName("Process").Int()

	// Allocate a temporary variable with type of the struct.
	// el.Elem() is the value contained in the interface
	definitions := reflect.New(el.Elem().Type()).Elem() // *core.Definitions
	definitions.Set(el.Elem())                          // reflected process definitions el will be assigned to the core definitions

	if numParticipants > 0 {
		collaboration := definitions.MethodByName("SetCollaboration")
		collaboration.Call([]reflect.Value{})
	}

	if numProcess > 0 {
		process := definitions.MethodByName("SetProcess")
		process.Call([]reflect.Value{reflect.ValueOf(int(numProcess))})
	}

	/*
		Actually, diagram is decoupled. So, no func needs to be called here ...

		diagram := definitions.MethodByName("SetDiagram")
		diagram.Call([]reflect.Value{reflect.ValueOf(1)}) // 1 represents number of diagrams
	*/
}

// GetCurrentlyCreatedFilename returns the current bpmn filename
// It is a helper method to get the current bpmn filename, which
// relies on the same instance
func (bldr Builder) GetCurrentlyCreatedFilename() string {
	return fmt.Sprintf("%s_%d", bldr.FilenamePrefix, bldr.Counter)
}

// Build builds the bpmn and json files and returns the builder
// and an error if an error occurs during the process.
// The method calls the ToBPMN and ToJSON methods.
func (bldr Builder) Build() (Builder, error) {
	if err := bldr.ToBPMN(); err != nil {
		return Builder{}, err
	}
	if err := bldr.ToJSON(); err != nil {
		return Builder{}, err
	}
	return bldr, nil
}
