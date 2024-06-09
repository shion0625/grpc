package products

import (
	"fmt"
	"github.com/shion0625/grpc/msa/internal/pkg/errs"
	"regexp"
	"unicode/utf8"
)

type productId struct {
	value string
}

func NewProductId(value string) (*productId, error) {
	const Length int = 36
	const REGEXP string = "([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})"

	if utf8.RuneCountInString(value) != Length {
		return nil, errs.NewDomainError(fmt.Sprintf("商品IDは%d文字でなければなりません。", Length))
	}

	// 引数の正規表現(UUID)チェック
	if !regexp.MustCompile(REGEXP).Match([]byte(value)) {
		return nil, errs.NewDomainError("商品IDはUUIDの形式でなければなりません。")
	}
	return &productId{value: value}, nil
}

func (pi *productId) Value() string {
	return pi.value
}

func (pi *productId) Equals(other *productId) bool {
	// ポインターが同じかどうか
	if pi == other {
		return true
	}
	// 値が同じかどうか
	return pi.Value() == other.Value()
}
