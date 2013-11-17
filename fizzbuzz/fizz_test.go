package fizzbuzz

// For me the best way to learn a new programming language is by
// watching another guy/gal program and then imitate him/her.
// So thanks to the GoConvey team for the testing framework
// and this awesome screencast: http://www.youtube.com/watch?v=vL_UD1oAF0E

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	Convey("Given an integer not divisible by 3 or 5", t, func() {
		Convey("should return the number as a String", func() {
			So(fizzbuzz(1), ShouldEqual, "1")
		})
	})
	Convey("Given an int divisible by 3", t, func() {
		Convey("should return fizz", func() {
			So(fizzbuzz(3), ShouldEqual, "fizz")
		})
	})
	Convey("Given an int divisible by 5", t, func() {
		Convey("should return buzz", func() {
			So(fizzbuzz(5), ShouldEqual, "buzz")
		})
	})
	Convey("Given an int divisible by 3 AND 5", t, func() {
		Convey("should return fizzbuzz", func() {
			So(fizzbuzz(15), ShouldEqual, "fizzbuzz")
		})
	})
}
