package gormrepository

import (
	"context"

	"github.com/thnthien/impa/entity"
)

type UpdateBaseRepo[Struct entity.IEntity, K any] struct {
	*BaseRepo
}

func NewUpdateBaseRepo[Struct entity.IEntity, K any](baseRepo *BaseRepo) *UpdateBaseRepo[Struct, K] {
	return &UpdateBaseRepo[Struct, K]{baseRepo}
}

func (b *UpdateBaseRepo[Struct, K]) Update(ctx context.Context, e *Struct) error {
	return b.GetDB(ctx).Save(e).Error
}
