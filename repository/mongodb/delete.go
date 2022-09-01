package baserepo

import (
	"context"
)

type DeleteRepo[E any] struct {
	*BaseRepo
}

func NewDeleteRepo[E any](base *BaseRepo) *DeleteRepo[E] {
	return &DeleteRepo[E]{base}
}

func (r *DeleteRepo[E]) Delete(ctx context.Context, id E) error {
	_, err := r.collection.DeleteOne(ctx, getIdFilter(id))
	return err
}
