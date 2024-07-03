//nolint:dupl
package application

import (
	"context"
	"database/sql"

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

func (ps *productServiceImpl) Add(ctx context.Context, product *products.Product) error {
	return ExecuteTransaction(ctx, &ps.transaction, func(tran *sql.Tx) error {
		if err := ps.rep.Exists(ctx, tran, product); err != nil {
			return err
		}

		if err := ps.rep.Create(ctx, tran, product); err != nil {
			return err
		}

		return nil
	})
}

func (ps *productServiceImpl) Update(ctx context.Context, product *products.Product) error {
	return ExecuteTransaction(ctx, &ps.transaction, func(tran *sql.Tx) error {
		if err := ps.rep.UpdateById(ctx, tran, product); err != nil {
			return err
		}

		return nil
	})
}

func (ps *productServiceImpl) Delete(ctx context.Context, product *products.Product) error {
	return ExecuteTransaction(ctx, &ps.transaction, func(tran *sql.Tx) error {
		if err := ps.rep.DeleteById(ctx, tran, product); err != nil {
			return err
		}

		return nil
	})
}
