package gormrepository

import (
	"context"
	"fmt"

	"github.com/thnthien/impa/entity"
)

type DeleteBaseRepo[E entity.IEntity, K any] struct {
	*BaseRepo
}

func NewDeleteBaseRepo[E entity.IEntity, K any](baseRepo *BaseRepo) *DeleteBaseRepo[E, K] {
	return &DeleteBaseRepo[E, K]{baseRepo}
}

func (b *DeleteBaseRepo[E, K]) Delete(ctx context.Context, id K) error {
	var obj E
	return b.GetDB(ctx).Delete(&obj, fmt.Sprintf("%s = ?", b.IDField), id).Error
}
