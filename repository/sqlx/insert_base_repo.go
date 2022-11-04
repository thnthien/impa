package sqlxrepository

import (
	"context"
	"fmt"
	"strings"

	"github.com/thnthien/impa/entity"
	querybuilder "github.com/thnthien/impa/internal/common/query-builder"
	commonreflect "github.com/thnthien/impa/internal/common/reflect"
)

type InsertBaseRepo[Struct entity.IEntity, K any] struct {
	*BaseRepo
}

func NewInsertBaseRepo[Struct entity.IEntity, K any](baseRepo *BaseRepo) *InsertBaseRepo[Struct, K] {
	return &InsertBaseRepo[Struct, K]{baseRepo}
}

func (r *InsertBaseRepo[Struct, K]) Insert(ctx context.Context, e *Struct) error {
	db := r.GetDB(ctx)
	driverName := db.DriverName()
	tableName := querybuilder.StandardizeIdentifier(driverName, r.TableName)
	cols := commonreflect.GetAllUnignoredNames("db", *e)
	columns := querybuilder.StandardizeIdentifiers(driverName, cols...)
	namedColumns := querybuilder.NamedColumns(cols)
	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, tableName, strings.Join(columns, ", "), strings.Join(namedColumns, ", "))
	_, err := r.GetDB(ctx).NamedExec(query, e)
	return err
}
