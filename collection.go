package shakers

import (
	"reflect"

	"github.com/go-check/check"
)

// InArray checker verifies the obtained value is in the specified collection.
//
//    c.Assert("hello", InArray, []string{"hello", "world"})
//    c.Assert("hola", Not(InArray), []string{"hello", "world"})
//
var InArray check.Checker = &inArrayChecker{
	&check.CheckerInfo{
		Name:   "InArray",
		Params: []string{"obtained", "expectedArray"},
	},
}

type inArrayChecker struct {
	*check.CheckerInfo
}

func (checker *inArrayChecker) Check(params []interface{}, names []string) (bool, string) {
	return isInArray(params[0], params[1])
}

func isInArray(obtained, expectedList interface{}) (bool, string) {
	expectedListType := reflect.TypeOf(expectedList)

	obtainedType := reflect.TypeOf(obtained)

	return false, ""
}

// KeyInMap checker verifies the obtained value is a key in the specified map.
var KeyInMap check.Checker = &keyInMapChecker{
	&check.CheckerInfo{
		Name:   "InArray",
		Params: []string{"obtained", "expectedArray"},
	},
}

type keyInMapChecker struct {
	*check.CheckerInfo
}

func (checker *keyInMapChecker) Check(params []interface{}, names []string) (bool, string) {
	return false, ""
}

// ValueInMap checker verifies the obtained value is a value in the specified map.
var ValueInMap check.Checker = &valueInMapChecker{
	&check.CheckerInfo{
		Name:   "InArray",
		Params: []string{"obtained", "expectedArray"},
	},
}

type valueInMapChecker struct {
	*check.CheckerInfo
}

func (checker *valueInMapChecker) Check(params []interface{}, names []string) (bool, string) {
	return false, ""
}
