package models

type User struct {
	Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsActive    bool   `json:"is_active"`
	IsStaff     bool   `json:"is_staff"`
	IsSuperuser bool   `json:"is_superuser"`
	Timestamp
	DeletedAt
}

func NewUser(isStaff bool, isSuperuser bool) *User {
	u := new(User)
	u.IsStaff = isStaff
	u.IsSuperuser = isSuperuser

	return u
}
