package posts

import (
	"context"

	"github.com/IndraNurfa/fastcampus/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, userID, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, userID, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id from database")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikedByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error count like post by post id from database")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get comment post by post id from database")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:           postDetail.ID,
			UserID:       postDetail.UserID,
			Username:     postDetail.Username,
			PostTitle:    postDetail.PostTitle,
			PostContent:  postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked:      postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comment:   comments,
	}, nil
}
