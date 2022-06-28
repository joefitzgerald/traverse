package traverse_test

import (
	"testing"

	"github.com/joefitzgerald/traverse"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testGroupDefinition(t *testing.T, when spec.G, it spec.S) {
	it.Before(func() {
		RegisterTestingT(t)
	})

	it("does stuff", func() {
		input := `name: team@example.com
base: "DC=example,DC=com"
includes:
  - username: "manager@example.com"
    expand-tree: true
  - username: "individual@example.com"
    expand-tree: false
    excludes:
      - "excluded@example.com"
ldap:
  base: "blergh"`

		group, err := traverse.UnmarshalGroupDefinition([]byte(input))
		Expect(err).ToNot(HaveOccurred())
		Expect(group).NotTo(BeNil())
		Expect(group.Name).To(Equal("team@example.com"))
		Expect(group.Extra).To(HaveKey("ldap"))
	})
}
