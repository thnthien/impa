package baserepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/thnthien/impa/entity"
)

type UpdateRepo[E entity.IEntity] struct {
	*BaseRepo
}

func NewUpdateRepo[E entity.IEntity](base *BaseRepo) *UpdateRepo[E] {
	return &UpdateRepo[E]{base}
}

func (r *UpdateRepo[E]) Update(ctx context.Context, obj *E) error {
	beforeUpdate(obj)
	id := getID(obj)
	_, err := r.collection.UpdateOne(ctx, getIdFilter(id), bson.M{"$set": obj})
	if err != nil {
		return err
	}
	return nil
}
