package name

import (
	"regexp"

	"github.com/tyndyll/pii/domain"
)

func init() {
	// TODO: We could gather some names here
}

var helpInfo = &domain.HelpInformation{
	Name: `name`,
	Description: `name checks data for the presence of the names listed in a predefined list of forenames and surnames`,
}

type Name struct {
	Forenames []*regexp.Regexp
	Surnames  []*regexp.Regexp
}

func (processor *Name) Process(data *string, config *map[string]interface{}) (*domain.ProcessorResult, error) {
	// Based on the config we could pull from a resource at this point that has been provided by the user
	result := []string{}
	result = append(result, processor.searchForNames(data, processor.Forenames)...)
	result = append(result, processor.searchForNames(data, processor.Surnames)...)

	if len(result) == 0 {
		return nil, nil
	}
	return &domain.ProcessorResult{
		Name:         helpInfo.Name,
		FoundResults: result,
	}, nil
}

func (processor *Name) searchForNames(data *string, list []*regexp.Regexp) []string {
	result := []string{}

	for _, name := range processor.Forenames {
		found := name.FindAllString(*data, -1)
		result = append(result, found...)
	}
	return result
}

func (processor *Name) Help() *domain.HelpInformation {
	return helpInfo
}

var Processor *Name