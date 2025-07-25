package domain

type RoleType string

const (
	RoleTypeAdmin    RoleType = "admin"
	RoleTypeCustomer RoleType = "customer"
)

type User struct {
	Id           int64    `json:"id"`
	Email        string   `json:"email"`
	PasswordHash string   `json:"-"`
	Role         RoleType `json:"role"`
}

func (r RoleType) IsValid() bool {
	return r == RoleTypeAdmin || r == RoleTypeCustomer
}
