package application

import (
	"context"

	"github.com/shion0625/grpc/msa/internal/app/commandservice/domain/models/products"
	"github.com/shion0625/grpc/msa/internal/app/commandservice/domain/repository"
)

type ProductService interface {
	Add(ctx context.Context, product *products.Product) error
	Update(ctx context.Context, product *products.Product) error
	Delete(ctx context.Context, product *products.Product) error
}

type productServiceImpl struct {
	rep         repository.Product
	transaction // transaction構造体のエンベデッド
}

// コンストラクタ
func NewproductServiceImpl(rep repository.Product) ProductService {
	return &productServiceImpl{rep: rep}
}

func (ins *productServiceImpl) Add(ctx context.Context, product *products.Product) error {
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = ins.complete(tran, err)
	}()

	if err = ins.rep.Exists(ctx, tran, product); err != nil {
		return err
	}

	if err = ins.rep.Create(ctx, tran, product); err != nil {
		return err
	}
	return err
}

func (ins *productServiceImpl) Update(ctx context.Context, product *products.Product) error {
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = ins.complete(tran, err)
	}()

	if err = ins.rep.UpdateById(ctx, tran, product); err != nil {
		return err
	}

	return err
}

func (ins *productServiceImpl) Delete(ctx context.Context, product *products.Product) error {
	tran, err := ins.begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = ins.complete(tran, err)
	}()

	if err = ins.rep.DeleteById(ctx, tran, product); err != nil {
		return err
	}

	return err
}
