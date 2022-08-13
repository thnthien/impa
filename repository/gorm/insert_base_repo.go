package gormrepository

import (
	"context"

	"github.com/thnthien/impa/entity"
)

type InsertBaseRepo[Struct entity.IEntity, K any] struct {
	*BaseRepo
}

func NewInsertBaseRepo[Struct entity.IEntity, K any](baseRepo *BaseRepo) *InsertBaseRepo[Struct, K] {
	return &InsertBaseRepo[Struct, K]{baseRepo}
}

func (b *InsertBaseRepo[Struct, K]) Insert(ctx context.Context, e *Struct) error {
	return b.GetDB(ctx).Create(e).Error
}
