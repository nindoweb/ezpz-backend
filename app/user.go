package app

import (
	"database/sql"

	"ezpz/internals/common"
)

type User struct {
	common.Model
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Username    string       `json:"username" validate:"required"`
	Email       string       `json:"email" validate:"required,email"`
	Password    string       `json:"password" validate:"required"`
	IsActive    bool         `json:"is_active"`
	IsStaff     bool         `json:"is_staff"`
	IsSuperuser bool         `json:"is_superuser"`
	ActivatedAt sql.NullTime `json:"activated_at"`
	common.Timestamp
	common.DeletedAt
}

func NewUser(isStaff bool, isSuperuser bool) *User {
	u := new(User)
	u.IsStaff = isStaff
	u.IsSuperuser = isSuperuser

	return u
}

type UserService interface {
	Create()
	Find()
	Get()
	Update()
	Delete()
}

func (u User) Create() {

}

func (u User) Find() {

}

func (u User) Get() {

}

func (u User) Update() {

}

func (u User) Delete() {

}
