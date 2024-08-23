package store

import "context"

type Store interface {
	FindHamstersFeed(ctx context.Context) ([]HamsterPost, error)
	FindHamsterById(ctx context.Context, id string) (*HamsterPost, error)
	CreateHamsterPost(ctx context.Context, post *CreateHamsterPost) error
}

type CreateHamsterPost struct {
	AuthorId string
	Content  string
}

type HamsterPost struct {
	Id            string  `db:"id"`
	AuthorId      string  `db:"author_id"`
	Author        string  `db:"author"`
	Content       string  `db:"content"`
	ImageKey      *string `db:"image_key"`
	LikesCount    int     `db:"likes_count"`
	CommentsCount int     `db:"comments_count"`
	CreatedAt     string  `db:"created_at"`
}
