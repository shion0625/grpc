package products

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/shion0625/grpc/msa/internal/pkg/errs"
)

type productId struct {
	value string
}

func NewProductId(value string) (*productId, *errs.DomainError) {
	const (
		LENGTH int    = 36                                                                       // フィールドの長さ
		REGEXP string = "([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})" // UUIDの正規表現
	)

	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("商品IDの長さは%d文字でなければなりません。", LENGTH))
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
