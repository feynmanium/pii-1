package adapters_test

import (
	"testing"

	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tyndyll/pii/adapters"
	"github.com/tyndyll/pii/domain"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func TestValidate_Unmarshal(t *testing.T) {
	Convey(`Given I have a JSON representation of a ValidationRequest`, t, func() {
		testJSON := []byte(`{
			"data": [
				"The attackers IP address was 192.168.0.1",
				"The victims IP address was 192.168.0.2",
				5
			],
			"processors": [
				"ipaddress"
			]
		}`)

		Convey(`When I Unmarshal the JSON in to a ValidationRequest`, func() {
			request := &adapters.ValidationRequest{}

			err := json.Unmarshal(testJSON, request)
			if err != nil {
				panic(err.Error())
			}

			Convey(`Then the request Data field will contain the correct number of items`, func() {
				So(len(request.Data), ShouldEqual, 3)
			})

			Convey(`Then the request Processors field will contain the correct number of items`, func() {
				So(len(request.Processors), ShouldEqual, 1)
			})

		})
	})
}

func TestValidationHTTPAdapter_Validate(t *testing.T) {
	Convey(`Given I have a ValidationHTTPAdapter`, t, func() {
		adapter := &adapters.ValidationHTTPAdapter{}

		Convey(`And I have a configured processor`, func() {
			processor := &testProcessor{
				HelpInfo: &domain.HelpInformation{
					Name: `test-processor`,
				},
			}

			adapter.ProcessorFunctions = map[string]domain.Processor{}
			adapter.ProcessorFunctions[processor.Name()] = processor

			Convey(`When I have a request for the naked route`, func() {
				req, err := http.NewRequest(`GET`, `/demo`, nil)
				if err != nil {
					panic(err.Error())
				}

				Convey(`And I make the request`, func() {
					response := httptest.NewRecorder()
					adapter.ProcessorInformation(response, req)

					output, err := ioutil.ReadAll(response.Body)
					if err != nil {
						panic(err.Error())
					}

					Convey(`Then the output should equal the JSON representation of the list of processor functions`, func() {
						So(output, ShouldEqual, []byte(``))
					})
				})
			})
		})
	})
}

// testProcessor implements the Processor interface and provides a tool for testing where the
// output of a processor needs to be controller for testing purposes
type testProcessor struct {
	ProcessorName string
	ProcessFunc   func(data *string, config *map[string]interface{}) (*domain.ProcessorResult, error)
	HelpInfo *domain.HelpInformation
}

func (processor *testProcessor) Process(data *string, config *map[string]interface{}) (*domain.ProcessorResult, error) {
	return processor.ProcessFunc(data, config)
}

func (processor *testProcessor) Help() *domain.HelpInformation {
	return processor.HelpInfo
}

func (processor *testProcessor) Name() string {
	return processor.ProcessorName
}
