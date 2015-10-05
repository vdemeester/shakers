package shakers

import (
	"reflect"
	"time"

	check "gopkg.in/check.v1"
)

// As a commodity, we bring all check.Checker variables into the current namespace to avoid having
// to think about check.X versus checker.X.
var (
	DeepEquals   = check.DeepEquals
	ErrorMatches = check.ErrorMatches
	FitsTypeOf   = check.FitsTypeOf
	HasLen       = check.HasLen
	Implements   = check.Implements
	IsNil        = check.IsNil
	Matches      = check.Matches
	Not          = check.Not
	NotNil       = check.NotNil
	PanicMatches = check.PanicMatches
	Panics       = check.Panics
)

// Equaler is an interface implemented if the type has a Equal method.
// This is used to compare struct using shakers.Equals.
type Equaler interface {
	Equal(Equaler) bool
}

// Equals checker verifies the obtained value is equal to the specified one.
// It is smart in a way that it supports several *types* (built-in, Equaler,
// time.Time)
//
//    c.Assert(myStruct, Equals, aStruct, check.Commentf("bouuuhh"))
//    c.Assert(myTime, Equals, aTime, check.Commentf("bouuuhh"))
//
var Equals check.Checker = &equalChecker{
	&check.CheckerInfo{
		Name:   "Equals",
		Params: []string{"obtained", "expected"},
	},
}

// Lower checker verifies the obtained value is lower to the specified one.
// It is smart in a way that it supports several *types* (built-in, Less,
// time.Time)
//
//    c.Assert(myStruct, Lower, aStruct, check.Commentf("hh"))
//    c.Assert(myInt, Lower, 0, check.Commentf("hh"))
//
var Lower check.Checker = &lowerChecker{
	&check.CheckerInfo{
		Name:   "Lower",
		Params: []string{"obtained", "expected"},
	},
}

// Greater checker verifies the obtained value is greater to the specified one.
// It is smart in a way that it supports several *types* (built-in, Less,
// time.Time)
//
//    c.Assert(myStruct, Lower, aStruct, check.Commentf("hh"))
//    c.Assert(myInt, Lower, 0, check.Commentf("hh"))
//
var Greater check.Checker = &greaterChecker{
	&check.CheckerInfo{
		Name:   "Greater",
		Params: []string{"obtained", "expected"},
	},
}

type greaterChecker struct {
	*check.CheckerInfo
}

func (checker *greaterChecker) Check(params []interface{}, names []string) (bool, string) {
	return greater(params[0], params[1])
}

type lowerChecker struct {
	*check.CheckerInfo
	equal bool
}

type equalChecker struct {
	*check.CheckerInfo
	equal bool
}

func (checker *equalChecker) Check(params []interface{}, names []string) (bool, string) {
	return isEqual(params[0], params[1])
}

func isEqual(obtained, expected interface{}) (bool, string) {
	switch obtained.(type) {
	case time.Time:
		return timeEquals(obtained, expected)
	case Equaler:
		return equalerEquals(obtained, expected)
	default:
		if reflect.TypeOf(obtained) != reflect.TypeOf(expected) {
			return false, "obtained value and expected value have not the same type."
		}
		return obtained == expected, ""
	}
}

func equalerEquals(obtained, expected interface{}) (bool, string) {
	expectedEqualer, ok := expected.(Equaler)
	if !ok {
		return false, "expected value must be an Equaler - implementing Equal(Equaler)."
	}
	obtainedEqualer, ok := obtained.(Equaler)
	if !ok {
		return false, "obtained value must be an Equaler - implementing Equal(Equaler)."
	}
	return obtainedEqualer.Equal(expectedEqualer), ""
}
