package gormrepository

import (
	"context"

	"github.com/thnthien/impa/entity"
)

type BulkInsertBaseRepo[Struct entity.IEntity, K any] struct {
	*BaseRepo
}

func NewBulkInsertBaseRepo[Struct entity.IEntity, K any](baseRepo *BaseRepo) *BulkInsertBaseRepo[Struct, K] {
	return &BulkInsertBaseRepo[Struct, K]{baseRepo}
}

func (b *InsertBaseRepo[Struct, K]) BulkInsert(ctx context.Context, es []Struct) ([]*Struct, error) {
	err := b.GetDB(ctx).Create(&es).Error
	if err != nil {
		return nil, err
	}
	pointers := make([]*Struct, 0, len(es))
	for i := range es {
		pointers = append(pointers, &es[i])
	}
	return pointers, err
}
