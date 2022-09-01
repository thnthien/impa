package mongodbrepository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseRepo struct {
	db         *mongo.Database
	client     *mongo.Client
	collection *mongo.Collection
}

func NewBaseRepo(db *mongo.Database, client *mongo.Client, collection *mongo.Collection) *BaseRepo {
	return &BaseRepo{
		db:         db,
		client:     client,
		collection: collection,
	}
}

func (r *BaseRepo) GetDB() *mongo.Database {
	return r.db
}

func (r *BaseRepo) GetClient() *mongo.Client {
	return r.client
}

func (r *BaseRepo) GetCollection() *mongo.Collection {
	return r.collection
}
