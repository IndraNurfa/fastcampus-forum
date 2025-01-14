package posts

import (
	"context"

	"github.com/IndraNurfa/fastcampus/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * limit

	response, err := s.postRepo.GetAllPost(ctx, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error to get all post repositories")
		return response, err
	}
	return response, nil
}
