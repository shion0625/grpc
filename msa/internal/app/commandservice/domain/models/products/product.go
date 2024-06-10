package products

import (
	"github.com/google/uuid"
	"github.com/shion0625/grpc/msa/internal/app/commandservice/domain/models/categories"
	"github.com/shion0625/grpc/msa/internal/pkg/errs"
)

type Product struct {
	id       *productId
	name     *productName
	price    *productPrice
	category *categories.Category
}

func NewProduct(name *productName, price *productPrice, category *categories.Category) (*Product, *errs.DomainError) {
	uid, err := uuid.NewRandom()
	if err != nil { // UUIDを生成する
		return nil, errs.NewDomainError(err.Error())
	}

	// 商品ID用値オブジェクトを生成する
	id, domainErr := NewProductId(uid.String())
	if domainErr != nil {
		return nil, domainErr
	}

	// 商品エンティティのインスタンスを生成して返す
	return &Product{
		id:       id,
		name:     name,
		price:    price,
		category: category,
	}, nil
}

func (p *Product) Equals(other *Product) (bool, *errs.DomainError) {
	if other == nil {
		return false, errs.NewDomainError("引数でnilが渡されました。")
	}
	return p.id.Equals(other.id), nil
}

func (p *Product) Id() *productId {
	return p.id
}

func (p *Product) Name() *productName {
	return p.name
}

func (p *Product) Price() *productPrice {
	return p.price
}
