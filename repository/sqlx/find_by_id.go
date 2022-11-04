package sqlxrepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/thnthien/impa/entity"
	querybuilder "github.com/thnthien/impa/internal/common/query-builder"
)

type FindByIDBaseRepo[Struct entity.IEntity, K any] struct {
	*BaseRepo
}

func NewFindByIDBaseRepo[Struct entity.IEntity, K any](baseRepo *BaseRepo) *FindByIDBaseRepo[Struct, K] {
	return &FindByIDBaseRepo[Struct, K]{baseRepo}
}

func (r *FindByIDBaseRepo[Struct, K]) FindByID(ctx context.Context, id K) (*Struct, error) {
	db := r.GetDB(ctx)
	driverName := db.DriverName()
	tableName := querybuilder.StandardizeIdentifier(driverName, r.TableName)
	idField := querybuilder.StandardizeIdentifier(driverName, r.IDField)
	query := fmt.Sprintf(`SELECT * FROM %s WHERE %s=?`, tableName, idField)
	obj := new(Struct)
	err := db.Get(obj, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return obj, nil
}
