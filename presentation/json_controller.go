package presentation

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    app "ldap/application/service"
)

type ldapRequest struct {
    Uid      string  `json:"uid"`
    Password string  `json:"password"`
}

// ここで戻り値をちゃんと定義しておいた方が良いと思う
// 現状ではレスポンスの中身がここで分からなくて読み手が困る
//type ldapResponse struct {}

func JsonEndPoint(svc app.LdapService) endpoint.Endpoint {
    return func(_ context.Context, request interface{}) (interface{}, error) {
        req := request.(ldapRequest)
        result := svc.SearchUser(req.Uid, req.Password)
        return result, nil
    }
}
