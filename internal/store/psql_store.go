package store

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type PsqlStore struct {
	db      *sqlx.DB
	loggger *log.Logger
}

func NewPsqlStore(db *sqlx.DB, logger *log.Logger) *PsqlStore {
	return &PsqlStore{db: db, loggger: logger}
}

func (ps *PsqlStore) FindHamstersFeed(ctx context.Context) ([]HamsterPost, error) {
	// query the database
	var hamsterPosts []HamsterPost
	err := ps.db.SelectContext(ctx, &hamsterPosts, "select * from hamster_post_feed hpf")
	if err != nil {
		ps.loggger.Printf("Error querying database: %v", err)
		return nil, fmt.Errorf("error querying hamsters feed")
	}

	return hamsterPosts, err
}

func (ps *PsqlStore) FindHamsterById(ctx context.Context, id string) (*HamsterPost, error) {
	// query the database
	var hamsterPost HamsterPost
	err := ps.db.Get(&hamsterPost, "select * from hamster_post_feed hpf where hpf.id = $1", id)
	if err != nil {
		ps.loggger.Printf("Error querying hamster post by id %s %v", id, err)
		return nil, fmt.Errorf("error querying hamster by id %s", id)
	}
	return &hamsterPost, err
}

func (ps *PsqlStore) CreateHamsterPost(ctx context.Context, post *CreateHamsterPost) (postId *string, err error) {
	err = ps.db.QueryRowxContext(ctx, "insert into hamster_posts (user_id, content) values ($1, $2) returning id", post.AuthorId, post.Content).Scan(&postId)
	if err != nil {
		ps.loggger.Printf("Error creating hamster post %v", err)
		return nil, fmt.Errorf("error creating hamster post")
	}
	return postId, err
}

func (ps *PsqlStore) CreateUser(ctx context.Context, user *CreateUser) (userId string, err error) {
	return
}
