package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// トランザクション内で扱うメソッド群
type Transacter interface {
	Tx() *sqlx.Tx
	Begin(ctx context.Context) error
	Commit() error
	Rollback() error
}

type Transaction struct {
	// DBインスタンス
	db Beginner
	// トランザクションで利用するDBインスタンス
	tx *sqlx.Tx
}

// トランザクション
func NewTransaction(db Beginner) *Transaction {
	return &Transaction{db: db}
}

// トラザクション開始
func (ac *Transaction) Begin(ctx context.Context) error {
	tx, err := ac.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("cannet connect transaction: %w", err)
	}
	ac.tx = tx
	return nil
}

// コミット
// トランザクションの最後に実行
func (ac *Transaction) Commit() error {
	if err := ac.tx.Commit(); err != nil {
		return fmt.Errorf("cannot commit: %w ", err)
	}
	return nil
}

// ロールバック
// トラザクションを開いてから、エラーが起きた時に実行する
func (ac *Transaction) Rollback() error {
	if err := ac.tx.Rollback(); err != nil {
		return fmt.Errorf("cannot rollback: %w", err)
	}
	return nil
}

// トランザクション用DBインスタンス
func (ac *Transaction) DB() *sqlx.Tx {
	return ac.tx
}
