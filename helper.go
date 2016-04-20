package pgxhelper

import "github.com/jackc/pgx"

// PgxPreparer interface can hold any object that possible prepare SQL statements.
// Currently (and primary used for) it can hold pgx.Conn & pgx.ConnPool.
type PgxPreparer interface {
	Prepare(name, sql string) (*pgx.PreparedStatement, error)
}

// MustPrepare is like pgxConn[Pool].Prepare but panics if the SQL cannot be parsed.
// It simplifies safe initialization of global variables holding prepared statements.
// It also assign name to prepared statement (currently name is SQL itself, but it is subject to change in near future).
func MustPrepare(p PgxPreparer, sql string) (stmtName string) {
	stmtName = sql // TODO may be use some other naming rule?
	if _, err := p.Prepare(stmtName, sql); err != nil {
		panic(`pgxhelper: Prepare(` + sql + `): ` + err.Error())
	}
	return
}
