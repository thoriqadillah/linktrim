package link

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thoriqadillah/linktrim/ent"
	"github.com/thoriqadillah/linktrim/ent/link"
	"github.com/thoriqadillah/linktrim/lib/helper"
)

type Store interface {
	Create(ctx context.Context, payload linkCreate) error
	GetAll(ctx context.Context, userID uuid.UUID, pagination helper.Pagination) ([]Link, error)
	GetOne(ctx context.Context, linkID uuid.UUID) (*Link, error)
	Update(ctx context.Context, linkID uuid.UUID, payload linkUpdate) error
	Delete(ctx context.Context, linkID uuid.UUID) error
}

type entStore struct {
	db *ent.Client
}

func NewStore(db *ent.Client) Store {
	return &entStore{
		db: db,
	}
}

func (s *entStore) Create(ctx context.Context, payload linkCreate) error {
	return s.db.Link.Create().
		SetOwnerID(payload.Owner).
		SetOriginal(payload.Original).
		SetTrimmed(payload.Trimmed).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Exec(ctx)
}

func (s *entStore) GetAll(ctx context.Context, userID uuid.UUID, pagination helper.Pagination) ([]Link, error) {
	result, err := s.db.Link.Query().
		Where(link.OwnerIDEQ(userID)).
		Limit(pagination.Limit).
		Offset(pagination.Page).
		All(ctx)
	if err != nil {
		return nil, err
	}

	links := make([]Link, len(result))
	for i, link := range result {
		links[i] = Link{
			ID:        link.ID,
			OwnerID:   link.OwnerID,
			Original:  link.Original,
			Trimmed:   link.Trimmed,
			CreatedAt: link.CreatedAt,
			UpdatedAt: link.UpdatedAt,
		}
	}

	return links, nil
}

func (s *entStore) GetOne(ctx context.Context, linkID uuid.UUID) (*Link, error) {
	result, err := s.db.Link.Get(ctx, linkID)
	if err != nil {
		return nil, fmt.Errorf("link not found")
	}

	return &Link{
		ID:        result.ID,
		OwnerID:   result.OwnerID,
		Original:  result.Original,
		Trimmed:   result.Trimmed,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (s *entStore) Update(ctx context.Context, linkID uuid.UUID, payload linkUpdate) error {
	return s.db.Link.Update().
		SetOriginal(payload.Original).
		SetTrimmed(payload.Trimmed).
		SetUpdatedAt(time.Now()).Exec(ctx)
}

func (s *entStore) Delete(ctx context.Context, linkID uuid.UUID) error {
	return s.db.Link.DeleteOneID(linkID).Exec(ctx)
}
