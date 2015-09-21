package shakers

import (
	check "gopkg.in/check.v1"
)

func init() {
	check.Suite(&StringCheckerS{})
}

type StringCheckerS struct{}

var _ = check.Suite(&StringCheckerS{})

func (s *StringCheckerS) TestContains(c *check.C) {
	testInfo(c, Contains, "Contains", []string{"value", "substring"})

	testCheck(c, Contains, true, "", "abcd", "bc")
	testCheck(c, Contains, false, "", "abcd", "efg")
	testCheck(c, Contains, false, "", "", "bc")
	testCheck(c, Contains, true, "", "abcd", "")
	testCheck(c, Contains, true, "", "", "")

	testCheck(c, Contains, false, "Obtained value is not a string and has no .String()", 12, "1")
	testCheck(c, Contains, false, "Substring must be a string", "", 1)
}
