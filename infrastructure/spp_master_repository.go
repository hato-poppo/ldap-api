package infrastructure

import (
    "fmt"
    repository "ldap/domain/service"
    ldapConnect "gopkg.in/ldap.v2"
)

type sppMaster struct {
    Host            string
    Port            string
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
        Port: "LDAP SERVER  PORT",
        BindDN: "BIND DN",
        BindUser: "BIND USER",
        BindPassword: "BIND USER PASSWORD",
    }
}


func (s sppMaster) Connect() bool {
    l, err := ldapConnect.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
    if err != nil {
        fmt.Printf("Cannot connect server.\n%s", err.Error())
        return false
    }
    defer l.Close()

    
    err = l.Bind(ldap.BindUser, ldap.BindPassword)
    if err != nil {
        log.Printf("Cannot connect with bind user\n%s", err.Error())
        return false
    }

    return true
}

func (s sppMaster) Search() (string, string, string) {
    // Connectは外部で宣言するよりも内部から読んだ方が良いが、connectにするとうまく動作しないので一旦諦めている
    fmt.Println("#infrastructure: LDAP server searched.")
    // 接続に失敗したら例外投げたい
    return "UserId", "UserName", "Email"
}
