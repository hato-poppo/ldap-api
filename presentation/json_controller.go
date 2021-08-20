package presentation

import (
    app "ldap/application/service"
    "context"
    // "encoding/json"
    // "net/http"
    "github.com/go-kit/kit/endpoint"
)

type ldapRequest struct {
    Uid      string  `json:"uid"`
    Password string  `json:"password"`
}

// ここで戻り値をちゃんと定義しておいた方が良い気がする
//type ldapResponse struct {}

func JsonEndPoint(svc app.LdapService) endpoint.Endpoint {
    return func(_ context.Context, request interface{}) (interface{}, error) {
        req := request.(ldapRequest)
        result := svc.SearchUser(req.Uid, req.Password)
        return result, nil
    }
}
