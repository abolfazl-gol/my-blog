package validate

import (
	"blog/proto"
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

var (
	v = validator.New()
)

func RegisterReq(req *proto.RegisterRequest) error {
	if err := v.Var(req.Email, "required,email"); err != nil {
		return fmt.Errorf("Email address is invalid")
	}

	return nil
}

func ValidateBlog(req *proto.CreateBlogRequest) error {
	if err := v.Var(req.Name, "alphanum"); err != nil {
		return fmt.Errorf("Name must by alpha")
	}

	if err := v.Var(req.Title, "min=6,max=50"); err != nil {
		return fmt.Errorf("Title must be between 6 and 50 chars")
	}
	return nil
}
