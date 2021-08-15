package application

import (
    "fmt"
    domain "ldap/domain/model"
)

// ここでコントローラー的な役割を担う
func Exchange() string {
    fmt.Println("#application service: Called by presentation.")
    fmt.Println("#application service: Call the domain model.")
    return domain.Message()
}
