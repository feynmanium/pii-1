package domain

// HelpInformation contains the information that will be associated with a processor.
type HelpInformation struct {
	// The name of the processor that this HelpInformation is associated with
	Name string
	// A description of the processor that this HelpInformation is associated with
	//
	// The description should provide enough information that when presented to a user they should enough understanding
	// about what functionality the processor provides and if it will suit their use case
	Description string
	// Options provides a mapping of the keywords that are used in the configs passed to the processor and what
	// modifications they perform. If possible any defaults and option values should be detailed
	Options map[string]string
}
