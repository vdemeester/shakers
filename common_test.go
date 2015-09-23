package shakers

import (
	"reflect"
	"testing"
	"time"

	check "gopkg.in/check.v1"
)

func init() {
	check.Suite(&CommonCheckerS{})
}

type CommonCheckerS struct{}

func (s *CommonCheckerS) TestEqualsInfo(c *check.C) {
	testInfo(c, Equals, "Equals", []string{"obtained", "expected"})
}

func (s *CommonCheckerS) TestEqualsValidsEquals(c *check.C) {
	myTime, err := time.Parse("2006-01-02", "2018-01-01")
	if err != nil {
		c.Fatal(err)
	}

	testCheck(c, Equals, true, "", "string", "string")
	testCheck(c, Equals, true, "", 0, 0)
	testCheck(c, Equals, true, "", anEqualer{1}, anEqualer{1})
	testCheck(c, Equals, true, "", myTime, myTime)
	testCheck(c, Equals, true, "", myTime, "2018-01-01")
	testCheck(c, Equals, true, "", myTime, "2018-01-01T00:00:00Z")
}

func (s *CommonCheckerS) TestEqualsValidsDifferent(c *check.C) {
	myTime1, err := time.Parse("2006-01-02", "2018-01-01")
	if err != nil {
		c.Fatal(err)
	}
	myTime2, err := time.Parse("2006-01-02", "2018-01-02")
	if err != nil {
		c.Fatal(err)
	}

	testCheck(c, Equals, false, "", "string", "astring")
	testCheck(c, Equals, false, "", 0, 1)
	testCheck(c, Equals, false, "", anEqualer{1}, anEqualer{0})
	testCheck(c, Equals, false, "", myTime1, myTime2)
	testCheck(c, Equals, false, "", myTime1, "2018-01-02")
	testCheck(c, Equals, false, "", myTime1, "2018-01-02T00:00:00Z")
}

func (s *CommonCheckerS) TestEqualsInvalids(c *check.C) {
	// Incompatible type time.Time
	testCheck(c, Equals, false, "obtained value and expected value have not the same type.", "2015-01-01", time.Now())
	testCheck(c, Equals, false, "expected must be a Time struct, or parseable.", time.Now(), 0)

	// Incompatible type Equaler
	testCheck(c, Equals, false, "expected value must be an Equaler - implementing Equal(Equaler).", anEqualer{0}, 0)

	testCheck(c, Equals, false, "obtained value and expected value have not the same type.", 0, anEqualer{0})

	// Nils
	testCheck(c, Equals, false, "obtained value and expected value have not the same type.", nil, 0)
	testCheck(c, Equals, false, "obtained value and expected value have not the same type.", 0, nil)
}

type anEqualer struct {
	value int
}

func (a anEqualer) Equal(b Equaler) bool {
	if bEqualer, ok := b.(anEqualer); ok {
		return a.value == bEqualer.value
	}
	return false
}

func Test(t *testing.T) {
	check.TestingT(t)
}

func testInfo(c *check.C, checker check.Checker, name string, paramNames []string) {
	info := checker.Info()
	if info.Name != name {
		c.Fatalf("Got name %s, expected %s", info.Name, name)
	}
	if !reflect.DeepEqual(info.Params, paramNames) {
		c.Fatalf("Got param names %#v, expected %#v", info.Params, paramNames)
	}
}

func testCheck(c *check.C, checker check.Checker, expectedResult bool, expectedError string, params ...interface{}) ([]interface{}, []string) {
	info := checker.Info()
	if len(params) != len(info.Params) {
		c.Fatalf("unexpected param count in test; expected %d got %d", len(info.Params), len(params))
	}
	names := append([]string{}, info.Params...)
	result, error := checker.Check(params, names)
	if result != expectedResult || error != expectedError {
		c.Fatalf("%s.Check(%#v) returned (%#v, %#v) rather than (%#v, %#v)",
			info.Name, params, result, error, expectedResult, expectedError)
	}
	return params, names
}
