package products

import (
	"fmt"
	"unicode/utf8"

	"github.com/shion0625/grpc/msa/internal/pkg/errs"
)

type productName struct {
	value string
}

func NewProductName(value string) (*productName, *errs.DomainError) {
	const (
		MIN_LENGTH int = 5
		MAX_LENGTH int = 30
	)

	length := utf8.RuneCountInString(value)

	if length < MIN_LENGTH || length > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("商品名の長さは%d文字以上、%d文字以内です。", MIN_LENGTH, MAX_LENGTH))
	}

	return &productName{value: value}, nil
}

func (pn *productName) Value() string {
	return pn.value
}
