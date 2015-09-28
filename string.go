// Package shakers provide some checker implementation the go-check.Checker interface.
package shakers

import (
	"fmt"
	"strings"

	check "gopkg.in/check.v1"
)

// Contains checker verifies that string value contains a substring.
var Contains check.Checker = &substringChecker{
	&check.CheckerInfo{
		Name:   "Contains",
		Params: []string{"obtained", "substring"},
	},
	strings.Contains,
}

// ContainsAny checker verifies that any Unicode code points in chars
// are in the obtained string.
var ContainsAny check.Checker = &substringChecker{
	&check.CheckerInfo{
		Name:   "ContainsAny",
		Params: []string{"obtained", "chars"},
	},
	strings.ContainsAny,
}

// HasPrefix checker verifies that string value has the specified substring as prefix
var HasPrefix check.Checker = &substringChecker{
	&check.CheckerInfo{
		Name:   "HasPrefix",
		Params: []string{"obtained", "prefix"},
	},
	strings.HasPrefix,
}

// HasSuffix checker verifies that string value has the specified substring as prefix
var HasSuffix check.Checker = &substringChecker{
	&check.CheckerInfo{
		Name:   "HasSuffix",
		Params: []string{"obtained", "suffix"},
	},
	strings.HasSuffix,
}

type substringChecker struct {
	*check.CheckerInfo
	substringFunction func(string, string) bool
}

func (checker *substringChecker) Check(params []interface{}, names []string) (bool, string) {
	return applyStringFunction(checker.substringFunction, params[0], params[1])
}

func applyStringFunction(stringFunction func(string, string) bool, obtained, substring interface{}) (bool, string) {
	substringStr, ok := substring.(string)
	if !ok {
		return false, "substring value must be a string."
	}
	obtainedString, obtainedIsStr := obtained.(string)
	if !obtainedIsStr {
		if obtainedWithStringer, obtainedHasStringer := obtained.(fmt.Stringer); obtainedHasStringer {
			obtainedString, obtainedIsStr = obtainedWithStringer.String(), true
		}
	}
	if obtainedIsStr {
		return stringFunction(obtainedString, substringStr), ""
	}
	return false, "obtained value is not a string and has no .String()."
}
