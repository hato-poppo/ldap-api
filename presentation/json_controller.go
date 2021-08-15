package presentation

import (
    "fmt"
    application "ldap/application/service"
)

func Get() string {
    fmt.Println("#presentation: Get http request.")
    fmt.Println("#presentation: Call the application service.")
    return application.Exchange()
}
