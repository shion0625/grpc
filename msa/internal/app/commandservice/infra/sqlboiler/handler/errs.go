package handler

import (
	"log"
	"net"

	"github.com/shion0625/grpc/msa/internal/pkg/errs"

	"github.com/go-sql-driver/mysql"
)

// データベースアクセスエラーのハンドリング
func DBErrHandler(err error) error {
	switch e := err.(type) {
	case *net.OpError: // 接続がタイムアウトかネットワーク関連の問題が原因で接続が確立できない?
		log.Println(e.Error())
		return errs.NewInternalError(e.Error())
	case *mysql.MySQLError: // MySQLエラー
		log.Printf("Code:%d Message:%s \n", e.Number, e.Message)

		if e.Number == 1062 { // 一意制約違反
			return errs.NewCRUDError("一意制約違反です。")
		}

		return errs.NewInternalError(e.Message)
	default: // その他のエラー
		log.Println(err.Error())
		return errs.NewInternalError(err.Error())
	}
}
