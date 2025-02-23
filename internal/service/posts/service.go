package posts

import (
	"context"

	"github.com/IndraNurfa/fastcampus/internal/configs"
	"github.com/IndraNurfa/fastcampus/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error

	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error

	GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error)


	GetPostByID(ctx context.Context, userID, postID int64) (*posts.Post, error)
	CountLikedByPostID(ctx context.Context, postID int64) (int, error)
	GetCommentByPostID(ctx context.Context, postID int64) ([]posts.Comment, error)

}

type service struct {
	cfg      *configs.Config
	postRepo postRepository
}

func NewService(cfg *configs.Config, postRepo postRepository) *service {
	return &service{
		cfg:      cfg,
		postRepo: postRepo,
	}
}
