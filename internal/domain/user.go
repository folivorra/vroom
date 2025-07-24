package domain

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	Id           int64  `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Role         Role   `json:"role"`
	Revoked      bool   `json:"revoked"`
}

func (r Role) IsValid() bool {
	return r == RoleAdmin || r == RoleUser
}
