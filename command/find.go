package command

import (
	"fmt"

	"github.com/joefitzgerald/traverse/ldap"
)

type LDAP struct {
	Email      string `name:"email" arg:"" help:"The person's email address."`
	URL        string `env:"LDAP_URL" required:"true" help:"The URL of the LDAP server. E.g. ldap.example.com:636."`
	User       string `env:"LDAP_USER" required:"true" help:"The username for connecting to the LDAP server."`
	Password   string `env:"LDAP_PASSWORD" required:"true" help:"The password for connecting to the LDAP server."`
	SearchBase string `env:"LDAP_SEARCH_BASE" optional:"true" help:"The base DN to search for the user. E.g. ou=people,dc=example,dc=com."`
}

type Find struct {
	LDAPCmd LDAP `cmd:"" name:"ldap"`
}

func (l *LDAP) Run() error {
	finder := ldap.New(l.URL, l.User, l.Password, l.SearchBase)

	result, err := finder.Find(l.Email)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
