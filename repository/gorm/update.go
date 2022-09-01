package gormrepository

import (
	"context"

	"github.com/thnthien/impa/entity"
)

type UpdateRepo[E entity.IEntity, K any] struct {
	*BaseRepo
}

func NewUpdateRepo[E entity.IEntity, K any](baseRepo *BaseRepo) *UpdateRepo[E, K] {
	return &UpdateRepo[E, K]{baseRepo}
}

func (b *UpdateRepo[E, K]) Update(ctx context.Context, e *E) error {
	return b.GetDB(ctx).Save(e).Error
}
