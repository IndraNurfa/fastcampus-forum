package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/IndraNurfa/fastcampus/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil {
		return nil
	}

	if user != nil {
		return errors.New("username or email alredy exist")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := memberships.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(pass),
		CreatedAt: now,
		CreatedBy: req.Email,
		UpdatedAt: now,
		UpdatedBy: req.Email,
	}
	return s.membershipRepo.CreateUser(ctx, model)
}