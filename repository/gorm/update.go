package gormrepository

import (
	"context"

	"github.com/thnthien/impa/entity"
)

type UpdateBaseRepo[E entity.IEntity, K any] struct {
	*BaseRepo
}

func NewUpdateBaseRepo[E entity.IEntity, K any](baseRepo *BaseRepo) *UpdateBaseRepo[E, K] {
	return &UpdateBaseRepo[E, K]{baseRepo}
}

func (b *UpdateBaseRepo[E, K]) Update(ctx context.Context, e *E) error {
	return b.GetDB(ctx).Save(e).Error
}
