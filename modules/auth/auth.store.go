package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/thoriqadillah/linktrim/ent"
	"github.com/thoriqadillah/linktrim/ent/user"
	"github.com/thoriqadillah/linktrim/lib/security"
	"golang.org/x/crypto/bcrypt"
)

type Store interface {
	Create(ctx context.Context, payload userCreate) error
	Login(ctx context.Context, payload userLogin) (*User, error)
}

type entStore struct {
	db *ent.Client
}

func NewStore(db *ent.Client) Store {
	return &entStore{
		db: db,
	}
}

func (s *entStore) Create(ctx context.Context, payload userCreate) error {
	if _, err := s.db.User.Create().
		SetName(payload.Name).
		SetEmail(payload.Email).
		SetPassword(security.Hash(payload.Password)).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx); err != nil {
		return err
	}

	return nil
}

func (s *entStore) Login(ctx context.Context, payload userLogin) (*User, error) {
	user, err := s.db.User.Query().Where(user.EmailEQ(payload.Email)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &User{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		ProfilePic: user.ProfilePic,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}, nil
}
