package domain_test

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tyndyll/pii/domain"
)

func TestProcessorResult_Unmarshal(t *testing.T) {
	Convey(`Given I a processor result in JSON format`, t, func() {
		testJSON := []byte(`{
			"processor": "Processor Searcher",
			"code": "missing-value",
			"message": "Message Value",
			"item": {
				"itemType": "Cell",
				"location": {
					"column": 5,
					"row": 4
				},
				"definition": [],
				"attributes": {}
			},
			"context": [
				{
					"itemType": "",
					"location": {},
					"definition": [],
					"attributes": {}
				} 
    		],
			"error-data": {}
		}`)

		Convey(`When I unmarshal the JSON into a ProcessorResult struct`, func() {
			result := &domain.ProcessorResult{}

			err := json.Unmarshal(testJSON, result)

			Convey(`Then the error will be nil`, func() {
				So(err, ShouldBeNil)
			})

			Convey(`Then the processor result Name field will be set correctly`, func() {
				So(result.Name, ShouldEqual, "Processor Searcher")
			})

			Convey(`Then the processor result Code field will be set correctly`, func() {
				So(result.Code, ShouldEqual, "missing-value")
			})

			Convey(`Then the processor result Message field will be set correctly`, func() {
				So(result.Message, ShouldEqual, "Message Value")
			})

			Convey(`Then the processor result Item Type field will be set correctly`, func() {
				So(result.Item.Type, ShouldEqual, "Cell")
			})

			Convey(`Then the processor result Item Location Row field will be set correctly`, func() {
				So(result.Item.Location.Row, ShouldEqual, 4)
			})

			Convey(`Then the processor result Item Location Column field will be set correctly`, func() {
				So(result.Item.Location.Column, ShouldEqual, 5)
			})
		})
	})
}
