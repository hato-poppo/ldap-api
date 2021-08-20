package application

import (
    "fmt"
    repository "ldap/domain/service"
    model "ldap/domain/model"
)

// ここでコントローラー的な役割を担う
func LdapUseCase(server repository.LdapRepository) model.LdapResult {
    ldapServer := model.Ldap{Server: server}
    return ldapServer.Search()
}
