package sqlxrepository

import (
	"context"
	"fmt"

	"github.com/thnthien/impa/entity"
	querybuilder "github.com/thnthien/impa/internal/common/query-builder"
)

type DeleteBaseRepo[Struct entity.IEntity, K any] struct {
	*BaseRepo
}

func NewDeleteBaseRepo[Struct entity.IEntity, K any](baseRepo *BaseRepo) *DeleteBaseRepo[Struct, K] {
	return &DeleteBaseRepo[Struct, K]{baseRepo}
}

func (r *DeleteBaseRepo[Struct, K]) Delete(ctx context.Context, id K) error {
	db := r.GetDB(ctx)
	driverName := db.DriverName()
	tableName := querybuilder.StandardizeIdentifier(driverName, r.TableName)
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=?`, tableName)
	_, err := r.GetDB(ctx).Exec(query, id)
	return err
}
