package infrastructure

import (
    "fmt"
    repository "ldap/domain/service"
)

type SppMaster struct {
    Host string
    User string
    Pass string
}

type LdapRepository interface {
    repository.LdapRepository
}

// コンストラクタ
// 呼び出しの際はこのメソッドをコールすること
func NewSppMaster() *SppMaster {
    return &SppMaster{
        Host: "SppMasterServer",
        User: "LdapUser",
        Pass: "LdapPassword",
    }
}


func (s SppMaster) Connect() bool {
    fmt.Println("#infrastructure: LDAP server connected.")
    return true
}

func (s SppMaster) Search() (string, string, string) {
    fmt.Println("#infrastructure: LDAP server searched.")
    return "UserId", "UserName", "Email"
}
