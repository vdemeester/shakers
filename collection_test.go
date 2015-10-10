package shakers

import (
	"github.com/go-check/check"
)

func init() {
	check.Suite(&CollectionCheckerS{})
}

type CollectionCheckerS struct{}

func (s *CollectionCheckerS) TestInArrayInfo(c *check.C) {
	testInfo(c, InArray, "InArray", []string{"obtained", "expectedArray"})
}

func (s *CollectionCheckerS) TestInArray(c *check.C) {
	testCheck(c, InArray, true, "", "hello", []string{"hello", "world"})
	testCheck(c, InArray, true, "", 1, []int{1, 2})
}
