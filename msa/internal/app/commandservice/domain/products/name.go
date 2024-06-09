package products

import (
	"fmt"
	"unicode/utf8"

	"github.com/shion0625/grpc/msa/internal/pkg/errs"
)

type productName struct {
	value string
}

func NewProductName(value string) (*productName, error) {
	const MIN_LENGTH int = 5
	const MAX_LENGTH int = 30

	length := utf8.RuneCountInString(value)

	if length < MIN_LENGTH || length > MAX_LENGTH {
		return nil, errs.NewDomainError(fmt.Sprintf("product name must be between %d and %d characters", MIN_LENGTH, MAX_LENGTH))
	}
	return &productName{value: value}, nil
}

func (pn *productName) Value() string {
	return pn.value
}
