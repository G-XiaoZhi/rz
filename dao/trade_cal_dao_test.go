package dao

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTradeCal(t *testing.T) {
	Convey("20210820 is a trade date", t, func() {
		calDate, isOpen := GetTradeCal("20210820")
		So(calDate, ShouldEqual, "20210820")
		So(isOpen, ShouldEqual, 1)
	})

	Convey("20210821 is not a trade date", t, func() {
		calDate, isOpen := GetTradeCal("20210821")
		So(calDate, ShouldEqual, "20210821")
		So(isOpen, ShouldEqual, 0)
	})
}