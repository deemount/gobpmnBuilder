package gobpmn_builder

import "github.com/deemount/gobpmnModels/pkg/core"

// Options
type Options struct {
	Counter     int
	CurrentFile string
	ModelType   string
	Def         *core.Definitions
	Repo        core.DefinitionsRepository
}
