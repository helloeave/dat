package dat

import (
	"context"
	"time"
)

// Result serves the same purpose as sql.Result. Defining
// it for the package avoids tight coupling with database/sql.
type Result struct {
	LastInsertID int64
	RowsAffected int64
}

// Execer is any object that executes and queries SQL.
type Execer interface {
	Cache(id string, ttl time.Duration, invalidate bool) Execer
	Timeout(time.Duration) Execer
	Interpolate() (string, []interface{}, error)
	Exec() (*Result, error)
	ExecContext(ctx context.Context) (*Result, error)

	QueryScalar(destinations ...interface{}) error
	QueryScalarContext(ctx context.Context, destinations ...interface{}) error
	QuerySlice(dest interface{}) error
	QuerySliceContext(ctx context.Context, dest interface{}) error
	QueryStruct(dest interface{}) error
	QueryStructContext(ctx context.Context, dest interface{}) error
	QueryStructs(dest interface{}) error
	QueryStructsContext(ctx context.Context, dest interface{}) error
	QueryObject(dest interface{}) error
	QueryObjectContext(ctx context.Context, dest interface{}) error
	QueryJSON() ([]byte, error)
	QueryJSONContext(ctx context.Context) ([]byte, error)
}

var nullExecer = &disconnectedExecer{}

// disonnectedExecer is the execer assigned when a builder is first created.
// Runners override the execer to work with a live database.
type disconnectedExecer struct{}

func (nop *disconnectedExecer) Cache(id string, ttl time.Duration, invalidate bool) Execer {
	return nil
}

func (nop *disconnectedExecer) Timeout(time.Duration) Execer {
	return nil
}

// Exec panics when Exec is called.
func (nop *disconnectedExecer) Exec() (*Result, error) {
	return nil, ErrDisconnectedExecer
}

func (nop *disconnectedExecer) ExecContext(_ context.Context) (*Result, error) {
	return nil, ErrDisconnectedExecer
}

func (nop *disconnectedExecer) Interpolate() (string, []interface{}, error) {
	return NewDatSQLErr(ErrDisconnectedExecer)
}

// QueryScalar panics when QueryScalar is called.
func (nop *disconnectedExecer) QueryScalar(destinations ...interface{}) error {
	return ErrDisconnectedExecer
}

func (nop *disconnectedExecer) QueryScalarContext(_ context.Context, destinations ...interface{}) error {
	return ErrDisconnectedExecer
}

// QuerySlice panics when QuerySlice is called.
func (nop *disconnectedExecer) QuerySlice(dest interface{}) error {
	return ErrDisconnectedExecer
}

func (nop *disconnectedExecer) QuerySliceContext(_ context.Context, dest interface{}) error {
	return ErrDisconnectedExecer
}

// QueryStruct panics when QueryStruct is called.
func (nop *disconnectedExecer) QueryStruct(dest interface{}) error {
	return ErrDisconnectedExecer
}

func (nop *disconnectedExecer) QueryStructContext(_ context.Context, dest interface{}) error {
	return ErrDisconnectedExecer
}

// QueryStructs panics when QueryStructs is called.
func (nop *disconnectedExecer) QueryStructs(dest interface{}) error {
	return ErrDisconnectedExecer
}

func (nop *disconnectedExecer) QueryStructsContext(_ context.Context, dest interface{}) error {
	return ErrDisconnectedExecer
}

// QueryObject panics when QueryObject is called.
func (nop *disconnectedExecer) QueryObject(dest interface{}) error {
	return ErrDisconnectedExecer
}

func (nop *disconnectedExecer) QueryObjectContext(_ context.Context, dest interface{}) error {
	return ErrDisconnectedExecer
}

// QueryJSON panics when QueryJSON is called.
func (nop *disconnectedExecer) QueryJSON() ([]byte, error) {
	return nil, ErrDisconnectedExecer
}

func (nop *disconnectedExecer) QueryJSONContext(_ context.Context) ([]byte, error) {
	return nil, ErrDisconnectedExecer
}
