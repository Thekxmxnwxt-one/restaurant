package models

type User struct {
	ID           int
	Name         string
	Username     string
	PasswordHash string
	Role         string // "employee" หรือ "customer"
}
