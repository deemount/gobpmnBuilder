package gobpmn_builder

// validate validates the builder fields
// and returns an error if the validation fails.
func (bldr *Builder) validate() error {
	if bldr.FilePathBPMN == "" {
		return ErrEmptyFilePathBPMN
	}
	return nil
}
