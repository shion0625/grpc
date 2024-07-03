package sqlboiler

import (
	"github.com/shion0625/grpc/msa/internal/app/commandservice/infra/sqlboiler/dao"

	"go.uber.org/fx"
)

// SQLBoilerを利用したRepositoryの依存定義
var RepDepend = fx.Options(
	fx.Provide(
		// Repositoryインターフェイス実装のコンストラクタを指定
		dao.NewCategoryRepositorySQLBoiler, // カテゴリ用Reposititory
		dao.NewProductRepositorySQLBoiler,  // 商品用Repository
	),
)
