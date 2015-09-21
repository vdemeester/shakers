// Package shakers provide some checker implementation the go-check.Checker interface.
package shakers

import (
	"fmt"
	"strings"

	check "gopkg.in/check.v1"
)

// Contains checker verifies that string value contains a substring.
var Contains check.Checker = &containsChecker{
	&check.CheckerInfo{
		Name:   "Contains",
		Params: []string{"value", "substring"},
	},
}

type containsChecker struct {
	*check.CheckerInfo
}

func (checker *containsChecker) Check(params []interface{}, names []string) (bool, string) {
	return contains(params[0], params[1])
}

func contains(value, substring interface{}) (bool, string) {
	substringStr, ok := substring.(string)
	if !ok {
		return false, "Substring must be a string"
	}
	valueStr, valueIsStr := value.(string)
	if !valueIsStr {
		if valueWithStr, valueHasStr := value.(fmt.Stringer); valueHasStr {
			valueStr, valueIsStr = valueWithStr.String(), true
		}
	}
	if valueIsStr {
		return strings.Contains(valueStr, substringStr), ""
	}
	return false, "Obtained value is not a string and has no .String()"
}
