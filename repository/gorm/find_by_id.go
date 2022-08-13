package gormrepository

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/thnthien/impa/entity"
)

type FindByIDBaseRepo[Struct entity.IEntity, K any] struct {
	*BaseRepo
}

func NewFindByIDBaseRepo[Struct entity.IEntity, K any](baseRepo *BaseRepo) *FindByIDBaseRepo[Struct, K] {
	return &FindByIDBaseRepo[Struct, K]{baseRepo}
}

func (b *FindByIDBaseRepo[Struct, K]) FindByID(ctx context.Context, id K) (*Struct, error) {
	obj := new(Struct)
	err := b.GetDB(ctx).First(obj, fmt.Sprintf("%s = ?", b.IDField), id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return obj, nil
}
