package domain

// ProcessorResult contains the result of a processor which has identified identifiable information
type ProcessorResult struct {
	// Name of the processor
	Name string `json:"processor"`
	// Unique code of the result. These should be provided by the processor
	Code string
	// A human readable version of the result
	Message string
	// The item that has been identified
	Item *Item `json:"item"`
}

type Item struct {
	Type     string `json:"itemType"`
	Location *Location
}

type Location struct {
	Column int
	Row    int64
}
