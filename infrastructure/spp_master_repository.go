package infrastructure

import (
    "fmt"
    repository "ldap/domain/service"
)

type sppMaster struct {
    Host            string
    BindDN          string
    BindUser        string
    BindPassword    string
}

type LdapRepository interface {
    repository.LdapRepository
}

// コンストラクタ
// 呼び出しの際はこのメソッドをコールすること
func NewSppMaster() *sppMaster {
    return &sppMaster{
        Host: "LDAP SERVER HOST",
        BindDN: "BIND DN",
        BindUser: "BIND USER",
        BindPassword: "BIND USER PASSWORD",
    }
}


func (s sppMaster) Connect() bool {
    fmt.Println("#infrastructure: LDAP server connected.")
    return true
}

func (s sppMaster) Search() (string, string, string) {
    fmt.Println("#infrastructure: LDAP server searched.")
    return "UserId", "UserName", "Email"
}
