package processors

import (
	"regexp"

	"github.com/tyndyll/pii/domain"
)

var ip4re = regexp.MustCompile(`\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}\b`)

type IPAddress struct{}

func (processor *IPAddress) Process(data string) (*domain.ProcessorResult, error) {
	result := ip4re.FindAllStringIndex(data, -1)

	// IP Address not found. Return nil
	if result == nil {
		return nil, nil
	}

	// IP Address found. Extract IP addresses and return ProcessorResult
	response := &domain.ProcessorResult{
		Name: "IPAddress",
		Code: "ip4.found",
	}
	return response, nil
}
