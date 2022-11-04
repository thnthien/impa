package sqlxrepository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/thnthien/impa/constant"
	"github.com/thnthien/impa/entity"
)

type ISqlxDB interface {
	sqlx.Execer
	sqlx.Queryer
	sqlx.QueryerContext
	sqlx.Preparer
	sqlx.PreparerContext
	sqlx.ExecerContext

	DriverName() string
	BindNamed(query string, arg any) (string, []any, error)
	NamedExec(query string, arg any) (sql.Result, error)
	NamedQuery(query string, arg any) (*sqlx.Rows, error)
	Select(dest any, query string, args ...any) error
	Get(dest any, query string, args ...any) error
	Preparex(query string) (*sqlx.Stmt, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
}

type BaseRepo struct {
	db        *sqlx.DB
	TableName string
	IDField   string
}

func NewBaseRepo[E entity.IEntity](db *sqlx.DB) *BaseRepo {
	var e E
	return &BaseRepo{db: db, TableName: e.TableName(), IDField: e.IDField()}
}

func (r *BaseRepo) GetDB(ctx context.Context) ISqlxDB {
	db, ok := ctx.Value(constant.CtxDBKey).(*sqlx.Tx)
	if !ok {
		return r.db
	}
	return db
}
