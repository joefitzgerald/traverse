package ldap

import (
	"crypto/tls"
	"log"

	ldapv3 "github.com/go-ldap/ldap/v3"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Searcher

type Searcher interface {
	Search(searchRequest *ldapv3.SearchRequest) (*ldapv3.SearchResult, error)
	Close()
}

type API struct {
	Searcher   Searcher
	SearchBase string
}

func (a *API) Close() {
	a.Searcher.Close()
}

func New(ldapURL string, ldapUser string, ldapPassword string, searchBase string) *API {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS10,
		MaxVersion:         tls.VersionTLS13,
	}

	connection, err := ldapv3.DialTLS("tcp", ldapURL, tlsConfig)
	if err != nil {
		log.Fatal(err)
	}
	if connection != nil {
		if err = connection.Bind(ldapUser, ldapPassword); err != nil {
			connection.Close()
			log.Fatal(err)
		}
	}

	return &API{
		Searcher:   connection,
		SearchBase: searchBase,
	}
}
