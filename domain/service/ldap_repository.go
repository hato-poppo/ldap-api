package domain

type LdapRepository interface {
    Connect() bool
    Search(string, string) (string, string, string, error)
}
