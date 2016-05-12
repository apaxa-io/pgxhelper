package pgxhelper

/*
import "github.com/apaxa-io/databasehelper/sqlhelper"

func ScanConvertAll(conn PgxQueryer, sql string, dst sqlhelper.MultiScannableConverter, args ...interface{}) error {
	rows, err := conn.Query(sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(dst.SqlScanInterface()...); err != nil {
			return err
		}
		if err := dst.ConvertFromSql(); err != nil {
			return err
		}
	}

	return rows.Err()
}
*/