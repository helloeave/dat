package runner

import (
	"context"

	"github.com/homelight/dat/dat"
)

// Connection is a queryable connection and represents a DB or Tx.
type Connection interface {
	Begin() (*Tx, error)
	BeginContext(ctx context.Context) (*Tx, error)
	Call(sproc string, args ...interface{}) *dat.CallBuilder
	DeleteFrom(table string) *dat.DeleteBuilder
	Exec(cmd string, args ...interface{}) (*dat.Result, error)
	ExecBuilder(b dat.Builder) error
	ExecMulti(commands ...*dat.Expression) (int, error)
	InsertInto(table string) *dat.InsertBuilder
	Insect(table string) *dat.InsectBuilder
	Select(columns ...string) *dat.SelectBuilder
	SelectDoc(columns ...string) *dat.SelectDocBuilder
	SQL(sql string, args ...interface{}) *dat.RawBuilder
	Update(table string) *dat.UpdateBuilder
	Upsert(table string) *dat.UpsertBuilder
	ExecContext(ctx context.Context, cmd string, args ...interface{}) (*dat.Result, error)
	ExecBuilderContext(ctx context.Context, b dat.Builder) error
	ExecMultiContext(ctx context.Context, commands ...*dat.Expression) (int, error)
}
