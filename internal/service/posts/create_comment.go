package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/IndraNurfa/fastcampus/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreateComment(ctx context.Context, postID, userID int64, req posts.CreateCommentRequest) error {
	now := time.Now()
	commentModel := posts.CommentModel{
		PostId:         postID,
		UserId:         userID,
		CommentContent: req.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userID, 10),
		UpdatedBy:      strconv.FormatInt(userID, 10),
	}
	err := s.postRepo.CreateComment(ctx, commentModel)
	if err != nil {
		log.Error().Err(err).Msg("error to create comment to repository")
	}
	return nil
}
