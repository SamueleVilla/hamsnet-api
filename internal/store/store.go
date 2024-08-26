package store

import "context"

type Store interface {
	FindHamstersFeed(ctx context.Context) ([]HamsterPost, error)
	FindHamsterById(ctx context.Context, id string) (*HamsterPost, error)
	CreateHamsterPost(ctx context.Context, post *CreateHamsterPost) (postId *string, err error)

	FindUserByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*User, error)
	CreateUser(ctx context.Context, user *CreateUser) (userId string, err error)
}

type CreateHamsterPost struct {
	AuthorId string
	Content  string
}

type CreateUser struct {
	Username       string
	Email          string
	HashedPassword string
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

type Role struct {
	RoleName string `db:"role_name"`
}

type User struct {
	Id             string `db:"id"`
	Username       string `db:"username"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
	CreatedAt      string `db:"created_at"`
	Roles          []Role
}
