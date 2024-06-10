package categories

import (
	"fmt"
	"github.com/shion0625/grpc/msa/internal/pkg/errs"
	"regexp"
	"unicode/utf8"
)

// カテゴリ番号を保持する値オブジェクト(UUIDを保持する)
type categoryId struct {
	value string
}

func NewCategoryId(value string) (*categoryId, *errs.DomainError) {
	// フィールドの長さ
	const LENGTH int = 36
	// UUIDの正規表現
	const REGEXP string = "([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})"
	// 引数の文字数チェック
	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("カテゴリIDの長さは%d文字でなければなりません。", LENGTH))
	}
	// 引数の正規表現(UUID)チェック
	if !regexp.MustCompile(REGEXP).Match([]byte(value)) {
		return nil, errs.NewDomainError("カテゴリIDはUUIDの形式でなければなりません。")
	}
	return &categoryId{value: value}, nil
}

// valueフィールドのゲッター
func (ci *categoryId) Value() string {
	return ci.value
}

// 同一性検証
func (ci *categoryId) Equals(value *categoryId) bool {
	if ci == value { // アドレスが同じ?
		return true
	}
	// 値が同じ
	return ci.value == value.Value()
}
