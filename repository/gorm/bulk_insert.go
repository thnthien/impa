package gormrepository

import (
	"context"

	"github.com/thnthien/impa/entity"
)

type BulkInsertBaseRepo[E entity.IEntity, K any] struct {
	*BaseRepo
}

func NewBulkInsertBaseRepo[E entity.IEntity, K any](baseRepo *BaseRepo) *BulkInsertBaseRepo[E, K] {
	return &BulkInsertBaseRepo[E, K]{baseRepo}
}

func (b *InsertBaseRepo[E, K]) BulkInsert(ctx context.Context, es []E) ([]*E, error) {
	err := b.GetDB(ctx).Create(&es).Error
	if err != nil {
		return nil, err
	}
	pointers := make([]*E, 0, len(es))
	for i := range es {
		pointers = append(pointers, &es[i])
	}
	return pointers, err
}
