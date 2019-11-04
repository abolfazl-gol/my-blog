package models

import (
	"blog/pkg/invalid"
	"blog/proto"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time `xorm:"<-"`
}

func (u *User) ToProto() *proto.User {
	return &proto.User{
		Id:    u.ID,
		Name:  u.Name,
		Token: u.Token,
	}
}

func (u *User) TableName() string { return "users" }

func (u *User) GenerateHash(password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.Password = string(hash)
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func CreateUser(user *User) error {
	return createUser(engine, user)
}

func GetUserByEmail(email string) (*User, error) {
	return getUserBy(engine, "email", email)
}

func GetUserByToken(token string) (*User, error) {
	return getUserBy(engine, "token", token)
}

func GetUserByID(id int64) (*User, error) {
	return getUserBy(engine, "id", id)
}

func createUser(e Engine, user *User) error {
	if _, err := e.Insert(user); err != nil {
		if strings.Contains(err.Error(), "1062") {
			return invalid.ErrDuplicate{"User", "email", user.Email}
		}
		return err
	}

	return nil
}

func getUserBy(e Engine, column string, value interface{}) (*User, error) {
	user := new(User)
	if _, err := e.Where(fmt.Sprintf("%s = ?", column), value).Get(user); err != nil {
		if err == sql.ErrNoRows {
			return nil, invalid.ErrNotFound{"User", column, value}
		}
		return nil, err
	}

	return user, nil
}
