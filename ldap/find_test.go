package ldap_test

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"

	ldapv3 "github.com/go-ldap/ldap/v3"
	"github.com/joefitzgerald/traverse/ldap"
	"github.com/joefitzgerald/traverse/ldap/ldapfakes"
)

func testFind(t *testing.T, when spec.G, it spec.S) {
	var api *ldap.API
	var fakeSearcher *ldapfakes.FakeSearcher

	it.Before(func() {
		RegisterTestingT(t)
		fakeSearcher = &ldapfakes.FakeSearcher{}
		api = &ldap.API{Searcher: fakeSearcher}
	})

	when("the search errors", func() {
		it.Before(func() {
			fakeSearcher.SearchReturns(nil, errors.New("fake error"))
		})

		it("returns an error", func() {
			user, err := api.Find("user@example.com")
			Expect(err).To(HaveOccurred())
			Expect(user).To(Equal(""))
		})
	})

	when("there are no users", func() {
		it.Before(func() {
			fakeSearcher.SearchReturns(&ldapv3.SearchResult{}, nil)
		})

		it("returns an error because no users can be found", func() {
			user, err := api.Find("user@example.com")
			Expect(err).To(HaveOccurred())
			Expect(user).To(Equal(""))
		})
	})

	when("there is a single user", func() {
		it.Before(func() {
			fakeSearcher.SearchReturns(&ldapv3.SearchResult{
				Entries: []*ldapv3.Entry{
					{
						DN: "cn=foo,dc=example,dc=com",
						Attributes: []*ldapv3.EntryAttribute{
							{
								Name:   "mail",
								Values: []string{"user@example.com"},
							},
						},
					},
				},
			}, nil)
		})

		it("finds a person by email", func() {
			user, err := api.Find("user@example.com")
			Expect(err).NotTo(HaveOccurred())
			Expect(user).To(Equal("user@example.com"))
		})

		it("returns an error if the user is not found", func() {
			user, err := api.Find("wronguser@example.com")
			Expect(err).To(HaveOccurred())
			Expect(user).To(Equal(""))
		})
	})

	when("there is more than one user", func() {
		it.Before(func() {
			fakeSearcher.SearchReturns(&ldapv3.SearchResult{
				Entries: []*ldapv3.Entry{
					{
						DN: "cn=foo,dc=example,dc=com",
						Attributes: []*ldapv3.EntryAttribute{
							{
								Name:   "mail",
								Values: []string{"user@example.com"},
							},
						},
					},
					{
						DN: "cn=foo,dc=example,dc=com",
						Attributes: []*ldapv3.EntryAttribute{
							{
								Name:   "mail",
								Values: []string{"user@example.com"},
							},
						},
					},
				},
			}, nil)
		})

		it("returns an error because the user is not unique", func() {
			user, err := api.Find("user@example.com")
			Expect(err).To(HaveOccurred())
			Expect(user).To(Equal(""))
		})
	})
}
