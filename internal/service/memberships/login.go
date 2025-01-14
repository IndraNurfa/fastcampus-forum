package memberships

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/IndraNurfa/fastcampus/internal/model/memberships"
	"github.com/IndraNurfa/fastcampus/pkg/jwt"
	tokenUtil "github.com/IndraNurfa/fastcampus/pkg/token"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) {
	// Get user by email
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Str("email", req.Email).Msg("failed to get user")
		return "", "", errors.New("invalid email or password")
	}

	// If user not found
	if user == nil {
		return "", "", errors.New("invalid email or password") // Avoid revealing which part is wrong
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Warn().Err(err).Str("email", req.Email).Msg("password mismatch")
		return "", "", errors.New("invalid email or password") // Same reason here
	}

	// Create JWT token
	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Str("email", req.Email).Msg("failed to create token")
		return "", "", err
	}

	existRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, user.ID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get latest refresh token")
		return "", "", err
	}

	if existRefreshToken != nil {
		return token, existRefreshToken.RefreshToken, nil
	}

	refreshToken := tokenUtil.GenerateRefreshToken()
	if refreshToken == "" {
		return token, "", errors.New("failed to generate refresh token")
	}

	err = s.membershipRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(10 * 24 * time.Hour),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    strconv.FormatInt(user.ID, 10),
		UpdatedBy:    strconv.FormatInt(user.ID, 10),
	})

	if err != nil {
		log.Error().Err(err).Msg("error inserting refresh token")
		return token, refreshToken, nil
	}

	return token, refreshToken, nil
}
