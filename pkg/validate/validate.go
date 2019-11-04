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
