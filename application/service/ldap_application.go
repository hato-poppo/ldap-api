package application

import (
    "fmt"
    repository "ldap/domain/service"
    model "ldap/domain/model" // ここでmodelを参照して良いかどうか未確認
)

// ここでコントローラー的な役割を担う
// 適切な名前が浮かばなかったから適当な名前
func Controller(server repository.LdapRepository) model.LdapResult {
    ldapServer := model.Ldap{Server: server}
    return ldapServer.Search()
}
