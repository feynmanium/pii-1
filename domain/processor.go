package domain

type Processor interface {
	Process(string) (*ProcessorResult, error)
}
