package domain_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tyndyll/pii/domain"
	. "github.com/smartystreets/goconvey/convey"
)

func ExampleHelpInformation() {
	helper := &domain.HelpInformation{
		Name: "Lintol Processor",
		Description: "This processor examines the provided data to identify whether the string Lintol is in the data",
		Options: map[string]string{
			"case-insensitive": "Boolean. Determines whether the search should be case insensitive (default: false)",
			"language": "String. Determins which language Lintol should be searched for. Options: en, fr, de. Default: en",
		},
	}

	data, _ := json.Marshal(helper)
	fmt.Println(string(data))
}

func TestHelpInformation_Unmarshal(t *testing.T) {
	Convey(`Given I have a JSON representation of a HelpInformation instance`,t, func() {
		name := "Test Processor"
		description := "Test Description"
		optionKey := "key"
		optionValue := "value"
		testJSON := []byte(fmt.Sprintf(`{
			"name": "%s",
			"description": "%s",
			"options": {
				"%s": "%s"
			}	
		}`, name, description, optionKey, optionValue))

		Convey(`When I unmarshal the JSON into a HelpInformation struct`, func() {
			help := &domain.HelpInformation{}
			if err := json.Unmarshal(testJSON, help); err != nil {
				panic(err.Error())
			}

			Convey(`Then the Name field will be set correctly`, func() {
				So(help.Name, ShouldEqual, name)
			})

			Convey(`Then the Description field will be set correctly`, func() {
				So(help.Description, ShouldEqual, description)
			})

			Convey(`Then the Options field will be set correctly`, func() {
				So(help.Options[optionKey], ShouldEqual, optionValue)
			})
		})
	})
}