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

    
    err = l.Bind(s.BindUser, s.BindPassword)
    if err != nil {
        fmt.Printf("Cannot connect with bind user\n%s", err.Error())
        return false
    }

    return true
}

func (s sppMaster) Search(uid string, password string) (string, string, string, error) {
    // Connectは外部で宣言するよりも内部から読んだ方が良いが、connectにするとうまく動作しないので一旦諦めている
    fmt.Println("#infrastructure: LDAP server searched.")


    // LDAP サーバー接続
    l, err := ldapConnect.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, 389))
    if err != nil {
        fmt.Printf("Cannot connect server.\n%s", err.Error())
        return "", "", "", err
    }
    defer l.Close()

    // LDAP バインドユーザーでの接続
    err = l.Bind(s.BindUser, s.BindPassword)
 
    if err != nil {
        fmt.Printf("Cannot connect with bind user\n%s", err.Error())
        return "", "", "", err
    }
 
    // LDAP 検索設定
    SearchRequest := ldapConnect.NewSearchRequest(
        s.BindDN,
        ldapConnect.ScopeWholeSubtree, ldapConnect.NeverDerefAliases, 0, 0, false,
        fmt.Sprintf("(sAMAccountName=%s)", uid),
        []string{"dn", "sAMAccountName", "displayName"},
        nil,
    )
 
    fmt.Printf("Start search user: %s", uid)
 
    // LDAP 検索実行
    sr, err := l.Search(SearchRequest)
    if err != nil {
        fmt.Printf("Search failed.\n%s", err.Error())
        return "", "", "", err
    }
 
    // 検索結果数確認 (1 件に特定できない場合は失敗)
    if len(sr.Entries) == 0 {
        fmt.Printf("Not found user: %s", uid)
        return "", "", "", fmt.Errorf("not found user: %s", uid)
    } else if len(sr.Entries) > 1 {
        fmt.Printf("Too many found user: %s", uid)
        return "", "", "", fmt.Errorf("too many found user: %s", uid)
    }
 
    entity := sr.Entries[0]
 
    // 認証ユーザーパスワード正誤確認
    err = l.Bind(entity.DN, password)
    if err != nil {
        fmt.Printf("Authorization failed user: %s", uid)
        return "", "", "", err
    }
 
    fmt.Printf("Authorization success user: %s", uid)


    // 接続に失敗したら例外投げたい
    return "UserId", "UserName", "Email", nil
}
