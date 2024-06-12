package repository

import (
	"database/sql"
	"errors"
)

type TransactionManager struct {
	db *sql.DB
	tx *sql.Tx
}

func NewTransactionManager(db *sql.DB) *TransactionManager {
	return &TransactionManager{db: db}
}

func (t *TransactionManager) StartTransaction() error {
	tx, err := t.db.Begin()
	if err != nil {
		t.tx = nil
		return err
	}
	t.tx = tx
	return nil
}

func (t *TransactionManager) Commit() error {
	if t.tx == nil {
		return errors.New("Транзакция не была открыта")
	}
	err := t.tx.Commit()
	t.tx = nil
	return err
}

func (t *TransactionManager) RollBack() error {
	if t.tx == nil {
		return errors.New("Транзакция не была открыта")
	}
	err := t.tx.Rollback()
	t.tx = nil
	return err
}

func (t *TransactionManager) QueryRow(query string, args ...any) *sql.Row {
	if t.tx == nil {
		return t.db.QueryRow(query, args...)
	}
	return t.tx.QueryRow(query, args...)
}

func (t *TransactionManager) Query(query string, args ...any) (*sql.Rows, error) {
	if t.tx == nil {
		return t.db.Query(query, args...)
	}
	return t.tx.Query(query, args...)
}

func (t *TransactionManager) Exec(query string, args ...any) (sql.Result, error) {
	if t.tx == nil {
		return t.db.Exec(query, args...)
	}
	return t.tx.Exec(query, args...)
}

func (t *TransactionManager) Prepare(query string) (*sql.Stmt, error) {
	if t.tx == nil {
		return t.db.Prepare(query)
	}
	return t.tx.Prepare(query)
}
