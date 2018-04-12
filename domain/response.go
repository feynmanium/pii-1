package domain

// Response is the result of a data being run against a list of processors
type Response struct {
	Version    string
	ItemCount  int64 `json:"item-count"`
	IssueCount int64 `json:"issue-count"`
	Format     string
	Info []*ProcessorResult
}
