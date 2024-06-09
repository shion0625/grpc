package products

import (
	"github.com/shion0625/grpc/msa/internal/pkg/errs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Productエンティティを構成する値オブジェクト", Ordered, Label("ProductName構造体の生成"), func() {
	var empty_str *errs.DomainError   // 空文字列　長さ5文字以上、30文字以内に違反する
	var length_over *errs.DomainError // 30文字より大きい文字列　長さ5文字以上、30文字以内に違反する
	var product_name *productName     // 30文字以内の文字列を指定する
	// 前処理
	BeforeAll(func() {
		_, empty_str = NewProductName("")
		_, length_over = NewProductName("aaaaaaaaaabbbbbbbbbbccccccccccd")
		product_name, _ = NewProductName("水性ボールペン")
	})
	// 文字数の検証
	Context("文字数の検証", Label("無効な文字数"), func() {
		It("空文字列の場合、errs.DomainErrorが返る", func() {
			Expect(empty_str).To(Equal(errs.NewDomainError("商品名の長さは5文字以上、30文字以内です。")))
		})
		It("30文字より大きいの場合,errs.DomainErrorが返る", func() {
			Expect(length_over).To(Equal(errs.NewDomainError("商品名の長さは5文字以上、30文字以内です。")))
		})
	})
	// 文字数の検証
	Context("有効な文字数の検証", Label("有効な文字数"), func() {
		It("5文字以上30文字以内場合,ProductNameが返る", func() {
			Expect(product_name.Value()).To(Equal("水性ボールペン"))
		})
	})
})
