package models

import (
	"blog/pkg/invalid"
	"blog/proto"
	"database/sql"
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

func createUser(e Engine, user *User) error {
	if _, err := e.Insert(user); err != nil {
		if strings.Contains(err.Error(), "1062") {
			return invalid.ErrDuplicate{"User", "email", user.Email}
		}
		return err
	}

	return nil
}

func GetUserByEmail(email string) (*User, error) {
	return getUserByEmail(engine, email)
}

func getUserByEmail(e Engine, email string) (*User, error) {
	user := new(User)
	if _, err := e.Where("email = ?", email).Get(user); err != nil {
		if err == sql.ErrNoRows {
			return nil, invalid.ErrNotFound{"User", "email", user.Email}
		}
		return nil, err
	}
	return user, nil
}

func GetUserByToken(token string) (*User, error) {
	return getUserByToken(engine, token)
}

func getUserByToken(e Engine, token string) (*User, error) {
	user := new(User)
	if _, err := e.Where("token = ?", token).Get(user); err != nil {
		if err == sql.ErrNoRows {
			return nil, invalid.ErrNotFound{"User", "token", user.Token}
		}
		return nil, err
	}
	return user, nil
}

func GetUserById(id interface{}) (*User, error) {
	return getUserById(engine, id.(int64))
}

func getUserById(e Engine, id int64) (*User, error) {
	user := new(User)
	if _, err := e.Where("id = ?", id).Get(user); err != nil {
		if err == sql.ErrNoRows {
			return nil, invalid.ErrNotFound{"User", "id", user.ID}
		}
		return nil, err
	}

	return user, nil
}
