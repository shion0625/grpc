package dao_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/shion0625/grpc/msa/internal/app/commandservice/infra/sqlboiler/handler"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRepimplPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "infra/sqlboiler/repositoryパッケージのテスト")
}

// 全テストが実行される前に1度だけ実行される関数
var _ = BeforeSuite(func() {
	absPath, _ := filepath.Abs("../config/database.toml")
	err := os.Setenv("DB_CONFIG_PATH", absPath)
	Expect(err).To(BeNil())
	err = handler.DBConnect() // データベースに接続する
	Expect(err).NotTo(HaveOccurred(), "データベース接続が失敗したのでテストを中止します。")
})
