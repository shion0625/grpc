package dao_test

import (
	"context"

	"github.com/shion0625/grpc/msa/internal/app/commandservice/domain/models/categories"
	"github.com/shion0625/grpc/msa/internal/app/commandservice/domain/repository"
	"github.com/shion0625/grpc/msa/internal/app/commandservice/infra/sqlboiler/dao"
	"github.com/shion0625/grpc/msa/internal/pkg/errs"

	"database/sql"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ = Describe("categoryRepositorySQLBoiler構造体", Ordered, Label("CategoryRepositoryインターフェイスメソッドのテスト"), func() {
	var rep repository.Category
	var ctx context.Context
	var tran *sql.Tx
	// 前処理
	BeforeAll(func() {
		// リポジトリの生成
		rep = dao.NewCategoryRepositorySQLBoiler()
	})

	// テスト毎の前処理
	BeforeEach(func() {
		ctx = context.Background()       // Contextの生成
		tran, _ = boil.BeginTx(ctx, nil) // トランザクションの開始
	})

	// テスト毎の後処理
	AfterEach(func() {
		err := tran.Rollback() // トランザクションのロールバック
		Expect(err).To(BeNil())
	})

	// Exists()メソッドのテスト
	Context("同名の商品カテゴリが存在確認結果を返す", Label("Exists"), func() {
		It("存在しない商品の場合nilが返る", func() {
			name, _ := categories.NewCategoryName("食品")
			category, _ := categories.NewCategory(name)
			result := rep.Exists(ctx, tran, category)
			Expect(result).To(BeNil())
		})
		It("存在するカテゴリ名の場合、errs.CRUDErrorが返る", func() {
			name, _ := categories.NewCategoryName("文房具")
			category, _ := categories.NewCategory(name)
			result := rep.Exists(ctx, tran, category)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("%sは既に登録されています。", category.Name().Value()))))
		})
	})
	// Create()メソッドのテスト
	Context("新しい商品カテゴリを永続化する", Label("Create"), func() {
		It("カテゴリが登録成功し、nilが返る", func() {
			name, _ := categories.NewCategoryName("食品")
			category, _ := categories.NewCategory(name)
			result := rep.Create(ctx, tran, category)
			Expect(result).To(BeNil())
		})
		It("obj_idが同じカテゴリを追加すると、errs.CRUDErrorが返る", func() {
			name, _ := categories.NewCategoryName("文房具")
			category, _ := categories.NewCategory(name)
			result := rep.Create(ctx, tran, category)
			Expect(result).To(Equal(errs.NewCRUDError("一意制約違反です。")))
		})
	})
	// UpdateById()メソッドのテスト
	Context("商品カテゴリを変更する", Label("UpdateById"), func() {
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			name, _ := categories.NewCategoryName("文房具")
			category, _ := categories.NewCategory(name)
			result := rep.UpdateById(ctx, tran, category)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("カテゴリ番号:%sは存在しないため、更新できませんでした。", category.Id().Value()))))
		})
		It("存在するobj_idを指定すると、nilが返る", func() {
			name, _ := categories.NewCategoryName("文房具1")
			category, _ := categories.NewCategory(name)
			result := rep.UpdateById(ctx, tran, category)
			Expect(result).To(BeNil())
		})
	})
	// DeleteById()メソッドのテスト
	Context("商品カテゴリを削除する", Label("DeleteById"), func() {
		It("存在しないobj_idを指定すると、errs.CRUDErrorが返る", func() {
			name, _ := categories.NewCategoryName("文房具1")
			category, _ := categories.NewCategory(name)
			result := rep.DeleteById(ctx, tran, category)
			Expect(result).To(Equal(errs.NewCRUDError(
				fmt.Sprintf("カテゴリ番号:%sは存在しないため、削除できませんでした。",
					category.Id().Value()))))
		})
		It("存在するobj_idを指定すると、nilが返る", func() {
			// 削除対象のカテゴリを登録する
			name, _ := categories.NewCategoryName("食品")
			category, _ := categories.NewCategory(name)
			err := rep.Create(ctx, tran, category)
			Expect(err).To(BeNil())
			// 登録したカテゴリを削除する
			result := rep.DeleteById(ctx, tran, category)
			Expect(result).To(BeNil())
		})
	})
})
