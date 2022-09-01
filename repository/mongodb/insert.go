package baserepo

import (
	"context"

	"github.com/thnthien/impa/entity"
)

type InsertRepo[E entity.IEntity] struct {
	*BaseRepo
}

func NewInsertRepo[E entity.IEntity](base *BaseRepo) *InsertRepo[E] {
	return &InsertRepo[E]{base}
}

func (r *InsertRepo[E]) Insert(ctx context.Context, obj *E) error {
	beforeCreate(obj)
	_, err := r.collection.InsertOne(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}
