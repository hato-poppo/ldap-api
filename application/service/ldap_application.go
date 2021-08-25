package application

import (
    repository "ldap/domain/service"
    model "ldap/domain/model"
)

type LdapService struct {
    Repository repository.LdapRepository
}

// ここでコントローラー的な役割を担う
func (l LdapService) SearchUser(uid string, password string) model.LdapResult {
    ldapServer := model.Ldap{Server: l.Repository}
    return ldapServer.Search(uid, password)
}
