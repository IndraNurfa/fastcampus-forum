package memberships

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/IndraNurfa/fastcampus/internal/model/memberships"
)

func (s *service) UpdateUser(ctx context.Context, userID int64, request memberships.UpdateUserRequest) error {
	now := time.Now()
	model := memberships.UserModel{
		ID:        userID,
		Email:     request.Email,
		Username:  request.Username,
		UpdatedAt: now,
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	err := s.membershipRepo.UpdateUser(ctx, model)
	if err != nil {
		return errors.New("error updating user")
	}
	return nil
}
