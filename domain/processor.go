package domain

// Processor is the interface that all Personally identifiable information processors must implement. It provides a
// static interface that anyone who wishes to implement a processor must satisfy, so that it can be imported into the
// service without recompiling or rewriting the main server code via Go's plugin interface
//
// Types that implement the Processor interface should access any configuration required via environmental variables.
// The implementing type will be created, so any information that needs to be gathered should be pulled in the init()
// function
//
type Processor interface {
	// Process invokes the functionality of the processor, taking in the data as string, and a map
	// of options with string keys and any type of
	Process(*string, *map[string]interface{}) (*ProcessorResult, error)
	// Help returns the HelpInformation for the processor
	Help() *HelpInformation
}
