package products

import (
	"fmt"

	"github.com/shion0625/grpc/msa/internal/pkg/errs"
)

type productPrice struct {
	value uint32
}

func NewProductPrice(value uint32) (*productPrice, error) {
	const MIN_VALUE uint32 = 50
	const MAX_VALUE uint32 = 10000

	if value < MIN_VALUE || value > MAX_VALUE {
		return nil, errs.NewDomainError(fmt.Sprintf("product price must be between %d and %d", MIN_VALUE, MAX_VALUE))
	}
	return &productPrice{value: value}, nil
}

func (pp *productPrice) Value() uint32 {
	return pp.value
}
