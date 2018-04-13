package domain

// ProcessorResult contains the result of a processor which has identified identifiable information
type ProcessorResult struct {
	// Name of the processor
	Name string `json:"processor"`
	// Unique code of the result. These should be provided by the processor
	FoundResults []string `json:"found-results"`
}
