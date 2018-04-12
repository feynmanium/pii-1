package domain_test

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tyndyll/pii/domain"
)

func TestResponse_Unmarshal(t *testing.T) {
	Convey(`Given a response in JSON format`, t, func() {
		testJSON := []byte(`{
			"version": "1",
			"item-count": 123,
			"issue-count": 456,
			"format": "CSV",
			"info": [],
			"warnings": [],
			"errors": [],
			"supplementary":[
				{
					"type": "",
					"source": "",
					"name": ""
			  	} 
			]
		}`)

		Convey(`When I unmarshal the JSON into a Response struct`, func() {
			response := &domain.Response{}

			err := json.Unmarshal(testJSON, response)

			Convey(`Then the error will be nil`, func() {
				So(err, ShouldBeNil)
			})

			Convey(`Then the response Version field will be set correctly`, func() {
				So(response.Version, ShouldEqual, "1")
			})

			Convey(`Then the response ItemCount field will be set correctly`, func() {
				So(response.ItemCount, ShouldEqual, 123)
			})

			Convey(`Then the response IssueCount field will be set correctly`, func() {
				So(response.IssueCount, ShouldEqual, 456)
			})

			Convey(`Then the response Format field will be set correctly`, func() {
				So(response.Format, ShouldEqual, "CSV")
			})

			Convey(`Then the response Info should be populated with 1 item`, func() {
				So(len(response.Info), ShouldEqual, 1)
			})
		})
	})
}
