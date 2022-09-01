package gormrepository

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/thnthien/impa/entity"
)

type FindByIDBaseRepo[E entity.IEntity, K any] struct {
	*BaseRepo
}

func NewFindByIDBaseRepo[E entity.IEntity, K any](baseRepo *BaseRepo) *FindByIDBaseRepo[E, K] {
	return &FindByIDBaseRepo[E, K]{baseRepo}
}

func (b *FindByIDBaseRepo[E, K]) FindByID(ctx context.Context, id K) (*E, error) {
	obj := new(E)
	err := b.GetDB(ctx).First(obj, fmt.Sprintf("%s = ?", b.IDField), id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return obj, nil
}
