package categories

import (
	"fmt"
	"github.com/shion0625/grpc/msa/internal/pkg/errs"
	"unicode/utf8"
)

// カテゴリ名を保持する値オブジェクト
type categoryName struct {
	value string
}

// コンストラクタ
func NewCategoryName(value string) (*categoryName, *errs.DomainError) {
	const MIN_LENGTH int = 2  // フィールドの最小文字数
	const MAX_LENGTH int = 20 // フィールドの最大文字数
	// 引数の文字数チェック
	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("カテゴリ名の長さは%d文字以上、%d文字以内です。", MIN_LENGTH, MAX_LENGTH))
	}
	return &categoryName{value: value}, nil
}

// valueフィールドのゲッター
func (cn *categoryName) Value() string {
	return cn.value
}
