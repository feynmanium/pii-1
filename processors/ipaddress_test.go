package processors_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tyndyll/pii/processors"
)

func TestIPAddress_Process(t *testing.T) {
	Convey(`Given I have an instance of an IPAddress processor`, t, func() {
		processor := &processors.IPAddress{}

		Convey(`And I have no IP address in the data`, func() {
			data := `Localhost IP`

			Convey(`When I call the Process method with the data`, func() {
				result, err := processor.Process(data)

				Convey(`Then the result will be nil`, func() {
					So(result, ShouldBeNil)
				})

				Convey(`Then the error will be nil`, func() {
					So(err, ShouldBeNil)
				})
			})
		})

		Convey(`And I have two valid IP4 addresses`, func() {
			data := ` 192.168.1.1 127.0.0.1`

			Convey(`When I call the Process method with the data`, func() {
				result, err := processor.Process(data)

				Convey(`Then the result FoundItems field will contain 2 items`, func() {
					So(len(result.FoundResults), ShouldEqual, 2)
				})

				Convey(`Then the error will be nil`, func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
