package command

import (
	"fmt"

	"github.com/joefitzgerald/traverse/ldap"
)

type LDAP struct {
	Email      string `name:"email" arg:"" help:"The person's email address."`
	Connection struct {
		URL        string `env:"URL" required:"true" help:"The URL of the LDAP server. E.g. ldap.example.com:636."`
		User       string `env:"USER" required:"true" help:"The username for connecting to the LDAP server."`
		Password   string `env:"PASSWORD" required:"true" help:"The password for connecting to the LDAP server."`
		SearchBase string `env:"SEARCH_BASE" optional:"true" help:"The base DN to search for the user. E.g. ou=people,dc=example,dc=com."`
	} `embed:"" prefix:"ldap-" envprefix:"LDAP_"`
}

type Find struct {
	LDAPCmd LDAP `cmd:"" name:"ldap"`
}

func (l *LDAP) Run() error {
	finder := ldap.New(l.Connection.URL, l.Connection.User, l.Connection.Password, l.Connection.SearchBase)

	result, err := finder.Find(l.Email)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
