package gormrepository

import (
	"context"
	"fmt"

	"github.com/thnthien/impa/entity"
)

type DeleteBaseRepo[Struct entity.IEntity, K any] struct {
	*BaseRepo
}

func NewDeleteBaseRepo[Struct entity.IEntity, K any](baseRepo *BaseRepo) *DeleteBaseRepo[Struct, K] {
	return &DeleteBaseRepo[Struct, K]{baseRepo}
}

func (b *DeleteBaseRepo[Struct, K]) Delete(ctx context.Context, id K) error {
	var obj Struct
	return b.GetDB(ctx).Delete(&obj, fmt.Sprintf("%s = ?", b.IDField), id).Error
}
