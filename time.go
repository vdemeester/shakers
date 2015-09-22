package shakers

import (
	"fmt"
	"time"

	check "gopkg.in/check.v1"
)

// Default format when parsing (in addition to RFC and default time formats..)
const shortForm = "2006-01-02"

// IsBetween checker verifies the specified time is between the specified start
// and end. It's exclusive so if the specified time is at the tip of the interval.
//
//    c.Assert(myTime, IsBetween, startTime, endTime, check.Commentf("bouuuhhh"))
//
var IsBetween check.Checker = &isBetweenChecker{
	&check.CheckerInfo{
		Name:   "IsBetween",
		Params: []string{"time", "start", "end"},
	},
}

type isBetweenChecker struct {
	*check.CheckerInfo
}

func (checker *isBetweenChecker) Check(params []interface{}, names []string) (bool, string) {
	return isBetween(params[0], params[1], params[2])
}

func isBetween(value, start, end interface{}) (bool, string) {
	startTime, ok := parseTime(start)
	if !ok {
		return false, "Start must be a Time struct."
	}
	endTime, ok := parseTime(end)
	if !ok {
		return false, "End must be a Time struct."
	}
	valueTime, valueIsTime := parseTime(value)
	if valueIsTime {
		return valueTime.After(startTime) && valueTime.Before(endTime), ""
	}
	return false, "Obtained value is not a time.Time struct or parseable as a time."
}

func parseTime(datetime interface{}) (time.Time, bool) {
	switch datetime.(type) {
	case time.Time:
		return datetime.(time.Time), true
	case string:
		return parseTimeAsString(datetime.(string))
	default:
		if datetimeWithStr, ok := datetime.(fmt.Stringer); ok {
			return parseTimeAsString(datetimeWithStr.String())
		}
		return time.Time{}, false
	}
}

func parseTimeAsString(timeAsStr string) (time.Time, bool) {
	forms := []string{shortForm, time.RFC3339, time.RFC3339Nano, time.RFC822, time.RFC822Z}
	for _, form := range forms {
		datetime, err := time.Parse(form, timeAsStr)
		if err == nil {
			return datetime, true
		}
	}
	return time.Time{}, false
}
