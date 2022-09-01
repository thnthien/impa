package mongodbrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/thnthien/impa/entity"
)

type UpdateRepo[E entity.IEntity, K any] struct {
	*BaseRepo
}

func NewUpdateRepo[E entity.IEntity, K any](base *BaseRepo) *UpdateRepo[E, K] {
	return &UpdateRepo[E, K]{base}
}

func (r *UpdateRepo[E, K]) Update(ctx context.Context, obj *E) error {
	beforeUpdate(obj)
	id := getID(obj)
	_, err := r.collection.UpdateOne(ctx, getIdFilter(id), bson.M{"$set": obj})
	if err != nil {
		return err
	}
	return nil
}
