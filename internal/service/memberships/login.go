package memberships

import (
	"context"
	"errors"

	"github.com/IndraNurfa/fastcampus/internal/model/memberships"
	"github.com/IndraNurfa/fastcampus/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	// Get user by email
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil {
		log.Error().Err(err).Str("email", req.Email).Msg("failed to get user")
		return "", errors.New("invalid email or password")
	}

	// If user not found
	if user == nil {
		return "", errors.New("invalid email or password") // Avoid revealing which part is wrong
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Warn().Err(err).Str("email", req.Email).Msg("password mismatch")
		return "", errors.New("invalid email or password") // Same reason here
	}

	// Create JWT token
	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Str("email", req.Email).Msg("failed to create token")
		return "", err
	}

	return token, nil
}
