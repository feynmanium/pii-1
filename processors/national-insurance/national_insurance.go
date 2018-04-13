package national_insurance

import (
	"github.com/tyndyll/pii/domain"
	"regexp"
)

var helpInfo = &domain.HelpInformation{
	Name: `national-insurance`,
	Description: `national-insurance checks data for the presence of a UK National Insurance number`,
}

var nationalInsuranceRe = regexp.MustCompile(`\b(?:[A-Z]{2}\s?\d{2}\s?\d{2}\s?\d{2}\s?[A-DFMP])\b`)

type NationalInsurance struct{}

func (processor *NationalInsurance) Process(data *string, config *map[string]interface{}) (*domain.ProcessorResult, error) {
	result := nationalInsuranceRe.FindAllString(*data, -1)

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

func (process *NationalInsurance) Help() *domain.HelpInformation {
	return helpInfo
}

var Processor *NationalInsurance