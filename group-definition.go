package traverse

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
