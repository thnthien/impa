package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/thnthien/impa/repository"
	"github.com/thnthien/impa/repository/mongodb"
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
	repository.IUpdate[Object, int64]
	repository.IDelete[Object, int64]
	repository.IFindByID[Object, int64]

	FindAll(ctx context.Context, filter Filter) ([]*Object, error)
}

type Filter struct {
	Name string `json:"name"`
}

type repoImpl struct {
	*mongodbrepository.BaseRepo
	*mongodbrepository.InsertRepo[Object, int64]
	*mongodbrepository.UpdateRepo[Object, int64]
	*mongodbrepository.DeleteRepo[Object, int64]
	*mongodbrepository.FindByIDRepo[Object, int64]
}

func (r *repoImpl) FindAll(ctx context.Context, filter Filter) ([]*Object, error) {
	queryFilter := bson.M{}
	if filter.Name != "" {
		queryFilter["name"] = filter.Name
	}
	cur, err := r.GetCollection().Find(ctx, queryFilter)
	if err != nil {
		return nil, err
	}
	var objs []*Object
	err = cur.All(ctx, &objs)
	return objs, err
}

func NewRepository(db *mongo.Database, client *mongo.Client) *repoImpl {
	obj := Object{}
	collection := db.Collection(obj.TableName())
	base := mongodbrepository.NewBaseRepo(db, client, collection)
	insert := mongodbrepository.NewInsertRepo[Object, int64](base)
	update := mongodbrepository.NewUpdateRepo[Object, int64](base)
	del := mongodbrepository.NewDeleteRepo[Object, int64](base)
	findByID := mongodbrepository.NewFindByIDRepo[Object, int64](base)
	return &repoImpl{base, insert, update, del, findByID}
}
