package repository

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/deemount/gobpmnLab/examples/collaborative_process"
	"github.com/deemount/gobpmnLab/models/bpmn/core"
)

// Builder ...
type Builder interface {
	Build() (builder, error)
	GetCurrentlyCreatedFile() string
}

// builder ...
type builder struct {
	// options
	Options Options
	// using pointer to interface and struct for more flexibel modelling
	Repo core.DefinitionsRepository
	Def  *core.Definitions
}

type BuilderOption func(o Options) Options

// NewBuilder ...
func NewBuilder(opt ...BuilderOption) Builder {
	//
	options := Options{}
	for _, o := range opt {
		options = o(options)
	}
	// path to bpmn files
	path := "files/bpmn"
	// read the dir for created bpmn files
	log.Printf("builder.go: %s", path)
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

	return &builder{Options: options}
}

// set is a private method and is called inside Create().
// It calls the Create() method in the written business model,
// when it fits to the correct expectations
func (builder *builder) set() {

	log.Println("builder.go: set model")

	// e.g. use interface pointer (no argument)

	//def := new(models.Definitions)
	//repositoryModel := examples.NewModel(def)

	// e.g. use struct pointer (no argument)

	// 1
	//builder.Repo = small_process.New().Build().Call()
	builder.Repo = collaborative_process.New().Build().Call()

	// 2
	// repositoryModel := examples.NewBlackBoxModel()
	//d := repositoryModel.Create()
	//builder.Repo = d.Def()

	// 3
	//repositoryModel := examples.NewSimpleModel001()
	//d := repositoryModel.Create()
	//builder.Repo = d.Def()

}

// Build...
func (bldr builder) Build() (builder, error) {

	var err error

	bldr.set()

	// create .bpmn
	log.Println("goBuilder.builder.go:Build create bpmn")
	err = bldr.toBPMN()
	if err != nil {
		return builder{}, err
	}

	return bldr, nil

}

// GetCurrentlyCreatedFilename ...
func (bldr builder) GetCurrentlyCreatedFile() string {
	return bldr.Options.CurrentFile
}

// toBPMN ...
func (bldr *builder) toBPMN() error {

	var err error

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(&bldr.Repo, " ", "  ")
	//b, _ := xml.MarshalIndent(&builder.Def, " ", "  ")
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
