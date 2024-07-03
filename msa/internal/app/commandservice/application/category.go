//nolint:dupl
package application

import (
	"context"
	"database/sql"

	"github.com/shion0625/grpc/msa/internal/app/commandservice/domain/models/categories"
	"github.com/shion0625/grpc/msa/internal/app/commandservice/domain/repository"
)

type CategoryService interface {
	Add(ctx context.Context, category *categories.Category) error
	Update(ctx context.Context, category *categories.Category) error
	Delete(ctx context.Context, category *categories.Category) error
}

type categoryService struct {
	rep         repository.Category
	transaction // transaction構造体のエンベデッド
}

func NewCategoryService(rep repository.Category) CategoryService {
	return &categoryService{rep: rep}
}

func (cs *categoryService) Add(ctx context.Context, category *categories.Category) error {
	return ExecuteTransaction(ctx, &cs.transaction, func(tran *sql.Tx) error {
		if err := cs.rep.Exists(ctx, tran, category); err != nil {
			return err
		}

		if err := cs.rep.Create(ctx, tran, category); err != nil {
			return err
		}

		return nil
	})
}

func (cs *categoryService) Update(ctx context.Context, category *categories.Category) error {
	return ExecuteTransaction(ctx, &cs.transaction, func(tran *sql.Tx) error {
		if err := cs.rep.UpdateById(ctx, tran, category); err != nil {
			return err
		}

		return nil
	})
}

func (cs *categoryService) Delete(ctx context.Context, category *categories.Category) error {
	return ExecuteTransaction(ctx, &cs.transaction, func(tran *sql.Tx) error {
		if err := cs.rep.DeleteById(ctx, tran, category); err != nil {
			return err
		}

		return nil
	})
}
