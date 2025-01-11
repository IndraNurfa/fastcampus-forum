package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/IndraNurfa/fastcampus/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error {
	postHashtags := strings.Join(req.PostHashtags, ",")

	now := time.Now()
	model := posts.PostModel{
		UserID:      userID,
		PostTitle:   req.PostTitle,
		PostContent: req.PostContent,
		PostHastags: postHashtags,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   strconv.FormatInt(userID, 10),
		UpdatedBy:   strconv.FormatInt(userID, 10),
	}
	err := s.postRepo.CreatePost(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error to create post to repository")
	}
	return nil
}
