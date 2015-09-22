package shakers

import (
	"time"

	check "gopkg.in/check.v1"
)

func init() {
	check.Suite(&TimeCheckerS{})
}

type TimeCheckerS struct{}

var _ = check.Suite(&TimeCheckerS{})

func (s *TimeCheckerS) TestIsBetweenInfo(c *check.C) {
	testInfo(c, IsBetween, "IsBetween", []string{"time", "start", "end"})
}

type randomStruct struct {
	foo string
	baz int
}

func (s *TimeCheckerS) TestIsBetweenValidBetween(c *check.C) {
	between := []struct {
		value interface{}
		start interface{}
		end   interface{}
	}{
		{
			value: "2018-01-02",
			start: "2018-01-01",
			end:   "2018-01-03",
		},
		{
			value: "2018-01-02T15:04:05Z",
			start: "2018-01-01T15:04:05Z",
			end:   "2018-01-03T15:04:05Z",
		},
		{
			value: "2018-01-02T15:04:05Z",
			start: "2018-01-02T15:03:05Z",
			end:   "2018-01-02T15:05:05Z",
		},
		{
			value: "2018-01-02T15:04:05+07:00",
			start: "2018-01-01T15:04:05+07:00",
			end:   "2018-01-03T15:04:05+07:00",
		},
		{
			value: "2018-01-02T15:04:05+07:00",
			start: "2018-01-02T15:03:05+07:00",
			end:   "2018-01-02T15:05:05+07:00",
		},
		{
			value: "2018-01-02T15:04:05+07:00",
			start: "2018-01-02T15:04:05+08:00",
			end:   "2018-01-02T15:04:05+06:00",
		},
		{
			value: "2018-01-02T15:04:05-07:00",
			start: "2018-01-02T15:04:05-06:00",
			end:   "2018-01-02T15:04:05-08:00",
		},
		{
			value: "2018-01-02T15:04:05.999999999Z",
			start: "2018-01-01T15:04:05.999999999Z",
			end:   "2018-01-03T15:04:05.999999999Z",
		},
		{
			value: "2018-01-02T15:04:05.999999999Z",
			start: "2018-01-02T15:03:05.999999999Z",
			end:   "2018-01-02T15:05:05.999999999Z",
		},
		{
			value: "2018-01-02T15:04:05.999999999+07:00",
			start: "2018-01-01T15:04:05.999999999+07:00",
			end:   "2018-01-03T15:04:05.999999999+07:00",
		},
		{
			value: "2018-01-02T15:04:05.999999999+07:00",
			start: "2018-01-02T15:03:05.999999999+07:00",
			end:   "2018-01-02T15:05:05.999999999+07:00",
		},
		{
			value: "2018-01-02T15:04:05.999999999+07:00",
			start: "2018-01-02T15:04:05.999999999+08:00",
			end:   "2018-01-02T15:04:05.999999999+06:00",
		},
		{
			value: "2018-01-02T15:04:05.999999999-07:00",
			start: "2018-01-02T15:04:05.999999999-06:00",
			end:   "2018-01-02T15:04:05.999999999-08:00",
		},
		{
			value: "02 Jan 18 15:04 MST",
			start: "01 Jan 18 15:04 MST",
			end:   "03 Jan 18 15:04 MST",
		},
		{
			value: "01 Jan 18 15:04 MST",
			start: "01 Jan 18 15:03 MST",
			end:   "01 Jan 18 15:05 MST",
		},
		{
			value: "02 Jan 18 15:04 +0700",
			start: "01 Jan 18 15:04 +0700",
			end:   "03 Jan 18 15:04 +0700",
		},
		{
			value: "01 Jan 18 15:04 +0700",
			start: "01 Jan 18 15:03 +0700",
			end:   "01 Jan 18 15:05 +0700",
		},
		{
			value: "01 Jan 18 15:04 +0700",
			start: "01 Jan 18 15:04 +0800",
			end:   "01 Jan 18 15:04 +0600",
		},
		{
			value: "01 Jan 18 15:04 -0700",
			start: "01 Jan 18 15:04 -0600",
			end:   "01 Jan 18 15:04 -0800",
		},
	}
	for _, d := range between {
		testCheck(c, IsBetween, true, "", d.value, d.start, d.end)
	}
}

func (s *TimeCheckerS) TestIsBetweenValidOutside(c *check.C) {
	outside := []struct {
		value interface{}
		start interface{}
		end   interface{}
	}{
		{
			value: "2018-01-02",
			start: "2018-01-12",
			end:   "2018-01-22",
		},
		{
			value: "2018-01-02T15:04:05Z",
			start: "2018-01-12T15:04:05Z",
			end:   "2018-01-22T15:04:05Z",
		},
		{
			value: "2018-01-02T15:04:05Z",
			start: "2018-01-02T15:05:05Z",
			end:   "2018-01-02T15:06:05Z",
		},
		{
			value: "2018-01-02T15:04:05+07:00",
			start: "2018-01-02T15:04:05+06:00",
			end:   "2018-01-02T15:04:05+05:00",
		},
		{
			value: "2018-01-02T15:04:05.999999990Z",
			start: "2018-01-02T15:04:05.999999995Z",
			end:   "2018-01-02T15:04:05.999999996Z",
		},
		{
			value: "2018-01-02T15:04:05.999999999-06:00",
			start: "2018-01-02T15:04:05.999999999-07:00",
			end:   "2018-01-02T15:04:05.999999999-08:00",
		},
		{
			value: "02 Jan 18 15:04 MST",
			start: "02 Jan 18 15:05 MST",
			end:   "02 Jan 18 15:06 MST",
		},
		{
			value: "02 Jan 18 15:04 +0700",
			start: "02 Jan 18 15:04 +0800",
			end:   "02 Jan 18 15:04 +0900",
		},
		{
			value: "2018-01-02",
			start: "2018-01-02",
			end:   "2018-01-10",
		},
		{
			value: "2018-01-02",
			start: "2018-01-01",
			end:   "2018-01-02",
		},
		{
			value: "2018-01-02",
			start: "2018-01-02",
			end:   "2018-01-02",
		},
		{
			value: "2018-01-02T15:04:05Z",
			start: "2018-01-02T15:03:05Z",
			end:   "2018-01-02T15:04:05Z",
		},
		{
			value: "2018-01-02T15:04:05Z",
			start: "2018-01-02T15:04:05Z",
			end:   "2018-01-02T15:05:05Z",
		},
		{
			value: "2018-01-02T15:04:05Z",
			start: "2018-01-02T15:04:05Z",
			end:   "2018-01-02T15:04:05Z",
		},
	}
	for _, d := range outside {
		testCheck(c, IsBetween, false, "", d.value, d.start, d.end)
	}
}

func (s *TimeCheckerS) TestIsBetweenInvalids(c *check.C) {
	// Nils
	testCheck(c, IsBetween, false, "Start must be a Time struct.", time.Now(), nil, time.Now())
	testCheck(c, IsBetween, false, "End must be a Time struct.", time.Now(), time.Now(), nil)
	testCheck(c, IsBetween, false, "Obtained value is not a time.Time struct or parseable as a time.", nil, time.Now(), time.Now())

	// wrong type
	testCheck(c, IsBetween, false, "Start must be a Time struct.", time.Now(), 0, time.Now())
	testCheck(c, IsBetween, false, "End must be a Time struct.", time.Now(), time.Now(), 0)
	testCheck(c, IsBetween, false, "Obtained value is not a time.Time struct or parseable as a time.", 0, time.Now(), time.Now())
	testCheck(c, IsBetween, false, "Start must be a Time struct.", time.Now(), randomStruct{}, time.Now())
	testCheck(c, IsBetween, false, "End must be a Time struct.", time.Now(), time.Now(), randomStruct{})
	testCheck(c, IsBetween, false, "Obtained value is not a time.Time struct or parseable as a time.", randomStruct{}, time.Now(), time.Now())

	// Invalid strings
	testCheck(c, IsBetween, false, "Start must be a Time struct.", time.Now(), "this is not a date", time.Now())
	testCheck(c, IsBetween, false, "End must be a Time struct.", time.Now(), time.Now(), "this is not a date")
	testCheck(c, IsBetween, false, "Obtained value is not a time.Time struct or parseable as a time.", "this is not a date", time.Now(), time.Now())

	// Invalids dates
	testCheck(c, IsBetween, false, "Start must be a Time struct.", time.Now(), "2018-31-02", time.Now())
	testCheck(c, IsBetween, false, "End must be a Time struct.", time.Now(), time.Now(), "2018-31-02")
	testCheck(c, IsBetween, false, "Obtained value is not a time.Time struct or parseable as a time.", "2018-31-02", time.Now(), time.Now())

}
