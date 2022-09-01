package gormrepository

import (
	"context"
	"fmt"

	"github.com/thnthien/impa/entity"
)

type DeleteRepo[E entity.IEntity, K any] struct {
	*BaseRepo
}

func NewDeleteRepo[E entity.IEntity, K any](baseRepo *BaseRepo) *DeleteRepo[E, K] {
	return &DeleteRepo[E, K]{baseRepo}
}

func (b *DeleteRepo[E, K]) Delete(ctx context.Context, id K) error {
	var obj E
	return b.GetDB(ctx).Delete(&obj, fmt.Sprintf("%s = ?", b.IDField), id).Error
}
