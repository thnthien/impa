package mongodbrepository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/thnthien/impa/entity"
)

type FindByIDRepo[E entity.IEntity, K any] struct {
	*BaseRepo
}

func NewFindByIDRepo[E entity.IEntity, K any](base *BaseRepo) *FindByIDRepo[E, K] {
	return &FindByIDRepo[E, K]{base}
}

func (r *FindByIDRepo[E, K]) FindByID(ctx context.Context, id K) (*E, error) {
	result := r.collection.FindOne(ctx, getIdFilter(id))
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, result.Err()
	}

	obj := new(E)
	err := result.Decode(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
