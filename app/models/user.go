package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"ezpz/pkg/common"
	"ezpz/pkg/db"
	"ezpz/pkg/notification"
	"ezpz/pkg/otp"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/mgo.v2/bson"
)

const (
	USER_COLLECTION string = "users"
)

type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FirstName   string        `json:"first_name" bson:"first_name"`
	LastName    string        `json:"last_name" bson:"last_name"`
	Username    string        `json:"username" bson:"username" binding:"required"`
	Email       string        `json:"email" bson:"email" binding:"required,email"`
	Password    string        `json:"password" bson:"password" binding:"required"`
	IsActive    bool          `json:"is_active" bson:"is_active"`
	IsStaff     bool          `json:"is_staff" bson:"is_staff"`
	IsSuperuser bool          `json:"is_superuser" bson:"is_superuser"`
	ActivatedAt sql.NullTime  `json:"activated_at" bson:"activated_at"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
	DeletedAt   sql.NullTime  `json:"deleted_at" bson:"deleted_at"`
}

func NewUser(isStaff bool, isSuperuser bool) *User {
	u := new(User)
	u.IsStaff = isStaff
	u.IsSuperuser = isSuperuser

	return u
}

type UserService interface {
	SendOtp() error
	HashPassword()
	Create()
	FindByUser()
}

func userCollection() db.Query {
	return db.NewQuery(USER_COLLECTION)
}

func (u User) String() string {
	return fmt.Sprintf(u.Username)
}

func (u User) SendOtp() error {
	message := fmt.Sprintf("Welcom to ezpz \n opt: %v", otp.Create(4))
	if err := notification.SendMail("register", []string{u.Email}, message); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u User) HashPassword(password string) {
	password, err := common.Hash(password)
	if err != nil {
		log.Println(err)
	}
	
	u.Password = password
	u.Create()
}

func (u User) Create() {
	userCollection().Create(u)
}

func (u User) FindByUser(username string) {
	data, err := userCollection().Where("username", u.Username).First()
	if err != nil {
		log.Println(err)
	}

	mapstructure.Decode(data, &u)
}