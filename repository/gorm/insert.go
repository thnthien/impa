package gormrepository

import (
	"context"

	"github.com/thnthien/impa/entity"
)

type InsertBaseRepo[E entity.IEntity, K any] struct {
	*BaseRepo
}

func NewInsertBaseRepo[E entity.IEntity, K any](baseRepo *BaseRepo) *InsertBaseRepo[E, K] {
	return &InsertBaseRepo[E, K]{baseRepo}
}

func (b *InsertBaseRepo[E, K]) Insert(ctx context.Context, e *E) error {
	return b.GetDB(ctx).Create(e).Error
}
