package cluster_test

import (
	"testing"

	. "github.com/pingcap/check"
)

type demoSuite struct{}

var _ = Suite(&demoSuite{})

func (s *demoSuite) TestNodeStartup(c *C) {
	c.Assert(nil, IsNil)
}

// 运行测试
func TestDemo(t *testing.T) {
	TestingT(t)
}
