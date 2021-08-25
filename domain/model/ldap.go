package domain

import (
    "fmt"
    repository "ldap/domain/service"
)

type Ldap struct {
    Server repository.LdapRepository
}

type LdapResult struct {
    Uid   string
    Name  string
    Email string
}

func (l *Ldap) Search(uid string, password string) LdapResult {
    fmt.Println("#domain model: Start the LDAP authentication.")
    isConnected := l.Server.Connect()
    if !isConnected {
        fmt.Println("接続失敗")
        // 本来はここでエラーか何かを返す
    }

    uid, name, email, err := l.Server.Search(uid, password)
    if err != nil {
        return LdapResult{}
    }

    return LdapResult{
        Uid: uid,
        Name: name,
        Email: email,
    }
}
