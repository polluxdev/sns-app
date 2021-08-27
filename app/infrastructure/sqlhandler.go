package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/polluxdev/sns-app/app/interfaces"
)

type SQLHandler struct {
	Conn *sql.DB
}

type Tx struct {
	Tx *sql.Tx
}

type Result struct {
	Result sql.Result
}

type Row struct {
	Rows *sql.Rows
}

func NewSQLHandler() (interfaces.SQLHandler, error) {
	sqlHandler := &SQLHandler{}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("HOST_URL"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"), os.Getenv("SSL_MODE"), os.Getenv("DB_TIME_ZONE"),
	)

	conn, err := sql.Open(os.Getenv("DB_DRIVER"), dsn)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	sqlHandler.Conn = conn

	return sqlHandler, nil
}

func (s *SQLHandler) Begin() (interfaces.Tx, error) {
	t, err := s.Conn.Begin()

	if err != nil {
		return nil, err
	}

	tx := &Tx{}
	tx.Tx = t

	return tx, nil
}

func (s *SQLHandler) Query(query string, args ...interface{}) (interfaces.Row, error) {
	rows, err := s.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	row := &Row{}
	row.Rows = rows

	return row, nil
}

func (s *SQLHandler) Exec(query string, args ...interface{}) (interfaces.Result, error) {
	result, err := s.Conn.Exec(query, args...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t Tx) Commit() error {
	err := t.Tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (t Tx) Rollback() error {
	err := t.Tx.Rollback()

	if err != nil {
		return err
	}

	return nil
}

func (t Tx) Exec(query string, args ...interface{}) (interfaces.Result, error) {
	result, err := t.Tx.Exec(query, args...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r Result) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r Result) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

func (r Row) Scan(value ...interface{}) error {
	return r.Rows.Scan(value...)
}

func (r Row) Next() bool {
	return r.Rows.Next()
}

func (r Row) Close() error {
	return r.Rows.Close()
}

func (r Row) Err() error {
	return r.Rows.Err()
}
