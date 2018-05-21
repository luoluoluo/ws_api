package library

import (
	"database/sql"
)

// DB db struct
type DB struct {
	*sql.DB
	driver string
}

// NewDB new
func NewDB(driver, dsn string) (*DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return &DB{db, driver}, err
}

// Insert insert
func (db *DB) Insert(sqlstr string, args ...interface{}) (int64, error) {
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Exec update/delete
func (db *DB) Exec(sqlstr string, args ...interface{}) (int64, error) {
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// SelectOne select one
func (db *DB) SelectOne(sqlstr string, args ...interface{}) (map[string]string, error) {
	res := make(map[string]string)
	rows, err := db.Select(sqlstr, args...)
	if err != nil {
		return res, err
	}
	if len(rows) != 0 {
		res = rows[0]
	}
	return res, nil
}

// Select 查询
func (db *DB) Select(sqlstr string, args ...interface{}) ([]map[string]string, error) {
	res := make([]map[string]string, 0)

	rows, err := db.Query(sqlstr, args...)
	if err != nil {
		return res, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return res, err
	}

	values := make([]string, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return res, err
		}
		vmap := make(map[string]string, len(scanArgs))
		for i, col := range values {
			vmap[columns[i]] = col
		}
		res = append(res, vmap)
	}
	return res, nil
}

// Query query
func (db *DB) Query(sqlstr string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
