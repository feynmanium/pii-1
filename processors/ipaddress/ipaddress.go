package main

import (
	"regexp"

	"github.com/tyndyll/pii/domain"
)


var helpInfo = &domain.HelpInformation{
	Name: `ip-address`,
	Description: `ip-address will search for IPv4 addresses in supplied data. IPv6 addresses will be added at a future date`,
}

var ip4re = regexp.MustCompile(`\b(((?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|\b)){4})\b`)

type IPAddress struct{}

func (processor *IPAddress) Process(data *string, config *map[string]interface{}) (*domain.ProcessorResult, error) {
	result := ip4re.FindAllString(*data, -1)

	// IP Address not found. Return nil
	if result == nil {
		return nil, nil
	}

	// IP Address found. Extract IP addresses and return ProcessorResult
	response := &domain.ProcessorResult{
		Name:         helpInfo.Name,
		FoundResults: result,
	}

	return response, nil
}

func (process *IPAddress) Help() *domain.HelpInformation {
	return helpInfo
}

var Processor *IPAddress
