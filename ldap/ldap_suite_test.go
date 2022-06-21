package ldap_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

var suite spec.Suite

func init() {
	suite = spec.New("ldap", spec.Report(report.Terminal{}))
	suite("search", testFind)
}

func TestLDAP(t *testing.T) {
	suite.Run(t)
}
