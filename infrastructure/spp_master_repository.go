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
    // ここで接続情報を参照する
    fmt.Println("#infrastructure: LDAP server connected.")
    return true
}

func (s sppMaster) Search() (string, string, string) {
    // Connectは外部で宣言するよりも内部から読んだ方が良いが、connectにするとうまく動作しないので一旦諦めている
    fmt.Println("#infrastructure: LDAP server searched.")
    // 接続に失敗したら例外投げたい
    return "UserId", "UserName", "Email"
}
