package mongodbrepository

import (
	"context"

	"github.com/thnthien/impa/entity"
)

type InsertRepo[E entity.IEntity, K any] struct {
	*BaseRepo
}

func NewInsertRepo[E entity.IEntity, K any](base *BaseRepo) *InsertRepo[E, K] {
	return &InsertRepo[E, K]{base}
}

func (r *InsertRepo[E, K]) Insert(ctx context.Context, obj *E) error {
	beforeCreate(obj, r.idGenerator)
	_, err := r.collection.InsertOne(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}
