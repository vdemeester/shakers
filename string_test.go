package shakers

import (
	check "gopkg.in/check.v1"
)

func init() {
	check.Suite(&StringCheckerS{})
}

type StringCheckerS struct{}

func (s *StringCheckerS) TestContains(c *check.C) {
	testInfo(c, Contains, "Contains", []string{"obtained", "substring"})

	testCheck(c, Contains, true, "", "abcd", "bc")
	testCheck(c, Contains, false, "", "abcd", "efg")
	testCheck(c, Contains, false, "", "", "bc")
	testCheck(c, Contains, true, "", "abcd", "")
	testCheck(c, Contains, true, "", "", "")

	testCheck(c, Contains, false, "obtained value is not a string and has no .String().", 12, "1")
	testCheck(c, Contains, false, "substring value must be a string.", "", 1)
}

func (s *StringCheckerS) TestContainsAny(c *check.C) {
	testInfo(c, ContainsAny, "ContainsAny", []string{"obtained", "chars"})

	testCheck(c, ContainsAny, true, "", "abcd", "b")
	testCheck(c, ContainsAny, true, "", "abcd", "b & c")
	testCheck(c, ContainsAny, false, "", "abcd", "e")
	testCheck(c, ContainsAny, false, "", "", "bc")
	testCheck(c, ContainsAny, false, "", "abcd", "")
	testCheck(c, ContainsAny, false, "", "", "")

	testCheck(c, ContainsAny, false, "obtained value is not a string and has no .String().", 12, "1")
	testCheck(c, ContainsAny, false, "substring value must be a string.", "", 1)
}

func (s *StringCheckerS) TestHasPrefix(c *check.C) {
	testInfo(c, HasPrefix, "HasPrefix", []string{"obtained", "prefix"})

	testCheck(c, HasPrefix, true, "", "abcd", "ab")
	testCheck(c, HasPrefix, false, "", "abcd", "efg")
	testCheck(c, HasPrefix, false, "", "", "bc")
	testCheck(c, HasPrefix, true, "", "abcd", "")
	testCheck(c, HasPrefix, true, "", "", "")

	testCheck(c, HasPrefix, false, "obtained value is not a string and has no .String().", 12, "1")
	testCheck(c, HasPrefix, false, "substring value must be a string.", "", 1)
}

func (s *StringCheckerS) TestHasSuffix(c *check.C) {
	testInfo(c, HasSuffix, "HasSuffix", []string{"obtained", "suffix"})

	testCheck(c, HasSuffix, true, "", "abcd", "cd")
	testCheck(c, HasSuffix, false, "", "abcd", "efg")
	testCheck(c, HasSuffix, false, "", "", "bc")
	testCheck(c, HasSuffix, true, "", "abcd", "")
	testCheck(c, HasSuffix, true, "", "", "")

	testCheck(c, HasSuffix, false, "obtained value is not a string and has no .String().", 12, "1")
	testCheck(c, HasSuffix, false, "substring value must be a string.", "", 1)
}
