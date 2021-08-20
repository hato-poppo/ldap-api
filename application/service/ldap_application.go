package application

import (
    "fmt"
    repository "ldap/domain/service"
    model "ldap/domain/model"
)

type LdapService struct {
    Repository repository.LdapRepository
}

// ここでコントローラー的な役割を担う
func (l LdapService) SearchUser(uid string, password string) model.LdapResult {
    fmt.Println(uid)
    fmt.Println(password)
    ldapServer := model.Ldap{Server: l.Repository}
    return ldapServer.Search()
}
