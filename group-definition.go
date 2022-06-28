package traverse

import "gopkg.in/yaml.v2"

// 1. Find people, using LDAP, MSGraph, Slack, Workday, Google, etc.
// 2. Find the group that the people should belong to, and identify current group membership.
// 3. Converge the expected group membership with the actual group membership, adding and removing people from the group as needed.

/*
---
  name: team@example.com
  base: "DC=example,DC=com"
  includes:
    - username: "manager@example.com"
      expand-tree: true
    - username: "individual@example.com"
      expand-tree: false
  excludes:
    - "excluded@example.com"
*/

type GroupDefinition struct {
	Name     string            `yaml:"name"`
	Type     string            `yaml:"type"`
	Includes []MemberDirective `yaml:"includes"`
	Excludes []string          `yaml:"excludes"`
	Extra    map[string]any
}

func UnmarshalGroupDefinition(data []byte) (*GroupDefinition, error) {
	var group GroupDefinition
	err := yaml.Unmarshal(data, &group)
	if err != nil {
		return nil, err
	}

	var extra map[string]any
	err = yaml.Unmarshal(data, &extra)
	if err != nil {
		return nil, err
	}
	group.Extra = extra
	return &group, nil
}

type MemberDirective struct {
	Username   string `yaml:"username"`
	ExpandTree bool   `yaml:"expand-tree"`
}
