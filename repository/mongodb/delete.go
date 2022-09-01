package mongodbrepository

import (
	"context"

	"github.com/thnthien/impa/entity"
)

type DeleteRepo[E entity.IEntity, K any] struct {
	*BaseRepo
}

func NewDeleteRepo[E entity.IEntity, K any](base *BaseRepo) *DeleteRepo[E, K] {
	return &DeleteRepo[E, K]{base}
}

func (r *DeleteRepo[E, K]) Delete(ctx context.Context, id K) error {
	_, err := r.collection.DeleteOne(ctx, getIdFilter(id))
	return err
}
