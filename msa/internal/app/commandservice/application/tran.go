package application

import (
	"context"
	"database/sql"
	"log"

	"github.com/shion0625/grpc/msa/internal/app/commandservice/infra/sqlboiler/handler"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// トランザクション制御
type transaction struct{}

// トランザクションを開始する
func (t *transaction) begin(ctx context.Context) (*sql.Tx, error) {
	// トランザクションを開始する
	tran, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, handler.DBErrHandler(err)
	}

	return tran, nil
}

// トランザクションを終了する
func (t *transaction) complete(tran *sql.Tx, err error) error {
	if err != nil {
		if e := tran.Rollback(); e != nil {
			return handler.DBErrHandler(err)
		} else {
			log.Println("トランザクションをロールバックしました。")
		}
	} else {
		if e := tran.Commit(); e != nil {
			return handler.DBErrHandler(err)
		} else {
			log.Println("トランザクションをコミットしました。")
		}
	}

	return nil
}

// トランザクションを実行する共通関数
func ExecuteTransaction(ctx context.Context, t *transaction, action func(tran *sql.Tx) error) error {
	tran, err := t.begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err := t.complete(tran, err); err != nil {
			log.Println("トランザクションの完了中にエラーが発生しました:", err)
		}
	}()

	err = action(tran)

	return err
}
