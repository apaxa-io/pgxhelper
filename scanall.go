package pgxhelper

import (
	"github.com/jackc/pgx"
	"github.com/apaxa-io/databasehelper/sqlhelper"
)

// PgxQueryer interface can hold any object that can query SQL statements.
// Currently (and is primary used for) it can hold pgx.Conn & pgx.ConnPool.
type PgxQueryer interface {
	Query(sql string, args ...interface{}) (*pgx.Rows, error)
}

// ScanAll is adaptation of "github.com/apaxa-io/databasehelper/sqlhelper" StmtScanAll for "github.com/jackc/pgx".
// ScanAll performs query sql on connection conn with arguments 'args' and stores all result rows in dst.
// sql passed as-is to conn.Query so it is possible to pass prepared statement name as sql.
// ScanAll stop working on first error.
// Example:
//  type Label struct {
//  	Id       int32
//  	Name     string
//  }
//
//  func (l *Label) SqlScanInterface() []interface{} {
//  	return []interface{}{
//  		&l.Id,
//  		&l.Name,
//  	}
//  }
//
//  type Labels []*Label
//
//  func (l *Labels) SqlNewElement() sqlhelper.SingleScannable {
//	e := &Label{}
//	*l = append(*l, e)
//	return e
//  }
//  ...
//  var labels Labels
//  if err := pgxhelper.ScanAll(conn, "SELECT id, name FROM LABELS where amount>$1", &labels, someAmount); err != nil {
//  	return err
//  }
func ScanAll(conn PgxQueryer, sql string, dst sqlhelper.MultiScannable, args ...interface{}) error {
	rows, err := conn.Query(sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		rowContainer := dst.SqlNewElement()
		if err := rows.Scan(rowContainer.SqlScanInterface()...); err != nil {
			return err
		}
	}

	return rows.Err()
}
