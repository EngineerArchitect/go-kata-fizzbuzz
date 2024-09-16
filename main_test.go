package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_numbers(t *testing.T) {

	testCases := []struct {
		x        int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{4, "4"},
	}

	for _, tc := range testCases {
		Convey(fmt.Sprintf("Given test number %d", tc.x), t, func() {

			Convey("Since number not a multiple of 3 or 5", func() {

				Convey("original number should be returned", func() {
					result := fizzbuzz(tc.x)
					So(result, ShouldEqual, tc.expected)
				})

			})
		})
	}
}

func Test_multiple_three(t *testing.T) {
	testCases := []struct {
		x        int
		expected bool
	}{
		{1, false},
		{3, true},
		{5, false},
		{6, true},
	}

	for _, tc := range testCases {
		Convey(fmt.Sprintf("Given test number %d", tc.x), t, func() {
			Convey("When number multiple of three", func() {
				Convey("Then true else false if not", func() {
					result := isMultipleOfThree(tc.x)
					So(result, ShouldEqual, tc.expected)
				})
			})
		})
	}
}

func Test_multiple_five(t *testing.T) {
	testCases := []struct {
		x        int
		expected bool
	}{
		{1, false},
		{5, true},
		{8, false},
		{10, true},
	}

	for _, tc := range testCases {
		Convey(fmt.Sprintf("Given test number %d", tc.x), t, func() {
			Convey("When number multiple of five", func() {
				Convey("Then true else false if not", func() {
					result := isMultipleOfFive(tc.x)
					So(result, ShouldEqual, tc.expected)
				})
			})
		})
	}
}

func Test_fizz(t *testing.T) {
	testCases := []struct {
		x        int
		expected string
	}{
		{1, "1"},
		{3, "Fizz"},
		{4, "4"},
		{6, "Fizz"},
	}

	for _, tc := range testCases {
		Convey(fmt.Sprintf("Given test number %d", tc.x), t, func() {
			Convey("multiple of three", func() {
				Convey("return the string 'Fizz'", func() {
					result := fizzbuzz(tc.x)
					So(result, ShouldEqual, tc.expected)
				})
			})
		})
	}
}

func Test_buzz(t *testing.T) {
	testCases := []struct {
		x        int
		expected string
	}{
		{1, "1"},
		{5, "Buzz"},
		{8, "8"},
		{10, "Buzz"},
	}

	for _, tc := range testCases {
		Convey(fmt.Sprintf("Given test number %d", tc.x), t, func() {
			Convey("multiple of five", func() {
				Convey("return the string 'Buzz'", func() {
					result := fizzbuzz(tc.x)
					So(result, ShouldEqual, tc.expected)
				})
			})
		})
	}
}

func Test_fizz_buzz(t *testing.T) {
	testCases := []struct {
		x        int
		expected string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
	}

	for _, tc := range testCases {
		Convey(fmt.Sprintf("Given test number %d", tc.x), t, func() {
			Convey("when multiple of three and five", func() {
				Convey("return the string 'FizzBuzz'", func() {
					result := fizzbuzz(tc.x)
					So(result, ShouldEqual, tc.expected)
				})
			})
		})
	}
}
