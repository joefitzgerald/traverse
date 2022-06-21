package ldap

import (
	"fmt"
	"strings"

	ldapv3 "github.com/go-ldap/ldap/v3"
)

var attributes []string = []string{"mail", "manager", "extensionAttribute10", "userPrincipalName"} // TODO: do we really need extensionAttribute10?

func (a *API) Find(identifier string) (string, error) {
	return a.GetPersonByEmail(identifier)
}

func (a *API) GetPersonByEmail(email string) (string, error) {
	filter := "(&(objectCategory=user)(mail=%s))"
	search := a.searchRequestForFilter(fmt.Sprintf(filter, email))
	sr, err := a.Searcher.Search(search)
	if err != nil {
		return "", fmt.Errorf("error searching for %s: %w", email, err)
	}
	if sr == nil || len(sr.Entries) == 0 {
		return "", fmt.Errorf("unable to find a user with email [%s]", email)
	}
	result := sr.Entries[0].GetAttributeValue("mail")
	if len(sr.Entries) != 1 || sr.Entries[0] == nil || !strings.EqualFold(email, result) {
		return "", fmt.Errorf("unable to find a unique user with email [%s]", email)
	}
	return result, nil
}

// func entryToPerson(entry *ldapv3.Entry) *traverse.Person {
// 	return &traverse.Person{
// 		DN:        entry.DN,
// 		Email:     entry.GetAttributeValue("mail"),
// 		Username:  entry.GetAttributeValue("userPrincipalName"),
// 		ManagerDN: entry.GetAttributeValue("manager"),
// 		IsManager: entry.GetAttributeValue("extensionAttribute10") == "isManager-Yes",
// 	}
// }

func (a *API) searchRequestForFilter(filter string) *ldapv3.SearchRequest {
	return ldapv3.NewSearchRequest(
		a.SearchBase,
		ldapv3.ScopeWholeSubtree,
		ldapv3.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		attributes,
		nil,
	)
}
