package gobpmn_builder

import (
	"fmt"
	"log"
	"os"
	"reflect"

	gobpmn_count "github.com/deemount/gobpmnCounter"
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
		Defaults(p interface{}, c *gobpmn_count.Quantities)
		ToBPMN() error
		GetCurrentlyCreatedFilename() string
	}

	// Builder ...
	Builder struct {
		Options Options
	}
)

// New ...
func New(option ...Option) BuilderRepository {
	//
	for _, o := range option {
		options = o(options)
	}

	// path to bpmn files
	path = "files/bpmn"
	files, err := os.ReadDir(path)
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

// Defaults receives the definitions repository by the app in p argument
// and calls the main elements to set the maps, including process parameters
// n of process. The method contains the reflected process definition (p interface{})
// and calls it by the reflected method name.
// This method hides specific setters (SetProcess, SetCollaboration, SetDiagram).
func (bldr *Builder) Defaults(p interface{}, c *gobpmn_count.Quantities) {

	// el is the interface {} of a given definition
	el := reflect.ValueOf(&p).Elem()
	counter := reflect.ValueOf(&c).Elem()

	// Get the number of processes
	//numProcess := counter.FieldByName("Process")

	log.Printf("Number of processes: %+v", counter.Elem().FieldByName("Process").Int())

	// Allocate a temporary variable with type of the struct.
	// el.Elem() is the value contained in the interface
	definitions := reflect.New(el.Elem().Type()).Elem() // *core.Definitions
	definitions.Set(el.Elem())                          // reflected process definitions el will be assigned to the core definitions

	/*
		collaboration := definitions.MethodByName("SetCollaboration")
		collaboration.Call([]reflect.Value{})
	}*/

	//if numProcess.Int() > 0 {
	process := definitions.MethodByName("SetProcess")
	process.Call([]reflect.Value{reflect.ValueOf(1)})
	//}

	/*

		Actually, diagram is decoupled. So, no func needs to be called here ...

		diagram := definitions.MethodByName("SetDiagram")
		diagram.Call([]reflect.Value{reflect.ValueOf(1)}) // 1 represents number of diagrams

	*/
}

// GetCurrentlyCreatedFilename ...
func (bldr Builder) GetCurrentlyCreatedFilename() string {
	return bldr.Options.CurrentFile
}

// Build...
func (bldr Builder) Build() (Builder, error) {
	if err := bldr.ToBPMN(); err != nil {
		return Builder{}, err
	}
	return bldr, nil
}
