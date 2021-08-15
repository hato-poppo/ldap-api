package domain

import (
    "fmt"
    repository "ldap/domain/service"
)

type Ldap struct {
    Server repository.LdapRepository
}

type LdapResult struct {
    Uid   string
    Name  string
    Email string
}

func Message() string {
    fmt.Println("#domain model: Start the LDAP authentication.")
    return "Successed"
}

func (l *Ldap) Search() LdapResult {
    fmt.Println("#domain model: Start the LDAP authentication.")
    l.Server.Connect()
    uid, name, email := l.Server.Search()
    return LdapResult{
        Uid: uid,
        Name: name,
        Email: email,
    }
}
