package gormrepository

import (
	"context"

	"gorm.io/gorm"

	"github.com/thnthien/impa"
	"github.com/thnthien/impa/entity"
)

type BaseRepo struct {
	db        *gorm.DB
	TableName string
	IDField   string
}

func NewBaseRepo[E entity.IEntity](db *gorm.DB) *BaseRepo {
	var e E
	return &BaseRepo{db: db, TableName: e.TableName(), IDField: e.IDField()}
}

func (b *BaseRepo) GetDB(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(impa.CtxDBKey).(*gorm.DB)
	if !ok {
		return b.db
	}
	return db
}
