package traverse_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

var suite spec.Suite

func init() {
	suite = spec.New("traverse", spec.Report(report.Terminal{}))
	suite("group definition", testGroupDefinition)
}

func Test(t *testing.T) {
	suite.Run(t)
}
