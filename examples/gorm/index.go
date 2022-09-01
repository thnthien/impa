package gorm

import (
	"context"

	"gorm.io/gorm"

	"github.com/thnthien/impa/repository"
	gormrepository "github.com/thnthien/impa/repository/gorm"
)

type Object struct {
	ID   int64  `gorm:"primaryKey" json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (Object) IDField() string {
	return "id"
}

func (Object) TableName() string {
	return "objects"
}

type IRepo interface {
	repository.IInsert[Object, int64]
	repository.IBulkInsert[Object, int64]
	repository.IUpdate[Object, int64]
	repository.IDelete[Object, int64]
	repository.IFindByID[Object, int64]

	FindAll(ctx context.Context, filter Filter) ([]*Object, error)
}

type Filter struct {
	Name string `json:"name"`
}

type repoImpl struct {
	*gormrepository.BaseRepo
	*gormrepository.InsertRepo[Object, int64]
	*gormrepository.BulkInsertRepo[Object, int64]
	*gormrepository.UpdateRepo[Object, int64]
	*gormrepository.DeleteRepo[Object, int64]
	*gormrepository.FindByIDRepo[Object, int64]
}

func (r *repoImpl) FindAll(ctx context.Context, filter Filter) ([]*Object, error) {
	var objs []*Object
	db := r.GetDB(ctx)

	if filter.Name != "" {
		db = db.Where("name = ?", filter.Name)
	}

	err := db.Find(&objs).Error
	return objs, err
}

func NewRepository(db *gorm.DB) *repoImpl {
	base := gormrepository.NewBaseRepo[Object](db)
	insert := gormrepository.NewInsertRepo[Object, int64](base)
	bulkInsert := gormrepository.NewBulkInsertRepo[Object, int64](base)
	update := gormrepository.NewUpdateRepo[Object, int64](base)
	del := gormrepository.NewDeleteRepo[Object, int64](base)
	findByID := gormrepository.NewFindByIDRepo[Object, int64](base)
	return &repoImpl{base, insert, bulkInsert, update, del, findByID}
}
