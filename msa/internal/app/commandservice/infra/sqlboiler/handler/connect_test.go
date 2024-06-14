package handler_test

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shion0625/grpc/msa/internal/app/commandservice/infra/sqlboiler/handler"
)

func TestConnect(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "infra/sqlboiler/handler パッケージのテスト")
}

var _ = Describe("データベース接続テスト", func() {
	It("データベース接続が成功した場合、nilが返る", Label("DB接続"), func() {
		absPath, _ := filepath.Abs("../config/database.toml")
		os.Setenv("DB_CONFIG_PATH", absPath)
		result := handler.DBConnect()
		Expect(result).To(BeNil())
	})
})
