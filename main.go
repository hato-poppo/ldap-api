package main

import (
	"fmt"
	"os"
	"log"
    ldapConnect "gopkg.in/ldap.v2"
	"errors"

	// "net/http"
	// "context"
	// "encoding/json"

	// "github.com/go-kit/kit/endpoint"
	// httptransport "github.com/go-kit/kit/transport/http"
)

// For onion
import (
    "fmt"
    api "ldap/presentation"
    infra "ldap/infrastructure"
    model "ldap/domain/model"
)
 
type Ldap struct {
    Host            string
    BindDN          string
    BindUser        string
    BindPassword    string
}
 
type LdapResult struct {
    LoginName   string
    Name        string
    DN          string
}
 
// LDAP 設定作成
func CreateLdapConn(ldapHost string, ldapBaseDN string, ldapBindUser string, ldapBindPassword string) *Ldap {
    return &amp;Ldap{
        Host: ldapHost,
        BindDN: ldapBaseDN,
        BindUser: ldapBindUser,
        BindPassword: ldapBindPassword,
    }
}
 
// LDAP 認証
func (ldap *Ldap) Login(username string, password string) (*LdapResult, error) {
    // LDAP サーバー接続
    l, err := ldapConnect.Dial("tcp", fmt.Sprintf("%s:%d", ldap.Host, 389))
    if err != nil {
        log.Printf("Cannot connect server.\n%s", err.Error())
        return nil, err
    }
    defer l.Close()
 
    // LDAP バインドユーザーでの接続
    err = l.Bind(ldap.BindUser, ldap.BindPassword)
 
    if err != nil {
        log.Printf("Cannot connect with bind user\n%s", err.Error())
        return nil, err
    }
 
    // LDAP 検索設定
    SearchRequest := ldapConnect.NewSearchRequest(
        ldap.BindDN,
        ldapConnect.ScopeWholeSubtree, ldapConnect.NeverDerefAliases, 0, 0, false,
        fmt.Sprintf("(sAMAccountName=%s)", username),
        []string{"dn", "sAMAccountName", "displayName"},
        nil,
    )
 
    log.Printf("Start search user: %s", username)
 
    // LDAP 検索実行
    sr, err := l.Search(SearchRequest)
    if err != nil {
        log.Printf("Search failed.\n%s", err.Error())
        return nil, err
    }
 
    // 検索結果数確認 (1 件に特定できない場合は失敗)
    if len(sr.Entries) == 0 {
        log.Printf("Not found user: %s", username)
        return nil, fmt.Errorf("not found user: %s", username)
    } else if len(sr.Entries) &gt; 1 {
        log.Printf("Too many found user: %s", username)
        return nil, fmt.Errorf("too many found user: %s", username)
    }
 
    entity := sr.Entries[0]
 
    // 認証ユーザーパスワード正誤確認
    err = l.Bind(entity.DN, password)
    if err != nil {
        log.Printf("Authorization failed user: %s", username)
        return nil, err
    }
 
    log.Printf("Authorization success user: %s", username)
 
    // sAMAccountName, displayName, DN を返却
    return &amp;LdapResult{
        LoginName: entity.GetAttributeValue("sAMAccountName"),
        Name: entity.GetAttributeValue("displayName"),
        DN: entity.GetAttributeValue("dn"),
    }, nil
}

// For onion
func main() {
    fmt.Println("#main: Created end points.")

    sppmaster := infra.NewSppMaster()
    service := application.LdapService{Repository: sppmaster}
    ep := api.JsonEndPoint(service)
    fmt.Println(ep)
}
