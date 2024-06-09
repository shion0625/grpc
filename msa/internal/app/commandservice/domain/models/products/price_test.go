package products

import (
	"github.com/shion0625/grpc/msa/internal/pkg/errs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Productエンティティを構成する値オブジェクト", Ordered, Label("ProductPrice構造体の生成"), func() {
	var min_err *errs.DomainError   // 50未満の単価
	var max_err *errs.DomainError   // 10000より大きい単価
	var product_price *productPrice // 商品単価
	// 前処理
	BeforeAll(func() {
		_, min_err = NewProductPrice(49)
		_, max_err = NewProductPrice(10001)
		product_price, _ = NewProductPrice(1500)
	})
	// 範囲外の単価の検証
	Context("範囲外の単価の検証", Label("無効な範囲"), func() {
		It("50未満の場合、errs.DomainErrorが返る", func() {
			Expect(min_err).To(Equal(errs.NewDomainError("単価は50以上、10000以下です。")))
		})
		It("10000より大きいの単価の場合、errs.DomainErrorが返る", func() {
			Expect(max_err).To(Equal(errs.NewDomainError("単価は50以上、10000以下です。")))
		})
	})
	// 範囲内の単価の検証
	Context("範囲外の単価の検証", Label("有効な範囲"), func() {
		It("50以上10000以下の単価の場合、ProductPriceが返る", func() {
			Expect(product_price.Value()).To(Equal(uint32(1500)))
		})
	})
})
