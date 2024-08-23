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

func (ps *PsqlStore) CreateHamsterPost(ctx context.Context, post *CreateHamsterPost) error {
	_, err := ps.db.ExecContext(ctx, "insert into hamster_posts (user_id, content) values ($1, $2)", post.AuthorId, post.Content)
	if err != nil {
		ps.loggger.Printf("Error creating hamster post %v", err)
		return fmt.Errorf("error creating hamster post")
	}

	return nil
}
