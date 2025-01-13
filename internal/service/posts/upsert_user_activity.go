package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/IndraNurfa/fastcampus/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {
	now := time.Now()
	model :=
		posts.UserActivityModel{
			PostId:    postID,
			UserId:    userID,
			IsLiked:   request.IsLiked,
			CreatedAt: now,
			UpdatedAt: now,
			CreatedBy: strconv.FormatInt(userID, 10),
			UpdatedBy: strconv.FormatInt(userID, 10),
		}
	userActivity, err := s.postRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error to get user activity")
	}
	if userActivity == nil {
		if !request.IsLiked {
			return errors.New("you not liked this post")
		}
		err = s.postRepo.CreateUserActivity(ctx, model)
	} else {
		err = s.postRepo.UpdateUserActivity(ctx, model)
	}
	if err != nil {
		log.Error().Err(err).Msg("error create or update user activity")
		return err
	}
	return nil
}
