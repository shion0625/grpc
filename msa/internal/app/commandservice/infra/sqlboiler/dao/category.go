package dao

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/shion0625/grpc/msa/internal/app/commandservice/domain/models/categories"
	"github.com/shion0625/grpc/msa/internal/app/commandservice/domain/repository"
	"github.com/shion0625/grpc/msa/internal/app/commandservice/infra/sqlboiler/handler"
	repoModels "github.com/shion0625/grpc/msa/internal/app/commandservice/infra/sqlboiler/models"
	"github.com/shion0625/grpc/msa/internal/pkg/errs"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CategoryRepositorySQLBoiler struct{}

func NewCategoryRepositorySQLBoiler() repository.Category {
	repoModels.AddCategoryHook(boil.AfterInsertHook, categoryAfterInsertHook)
	repoModels.AddCategoryHook(boil.AfterUpdateHook, categoryAfterUpdateHook)
	repoModels.AddCategoryHook(boil.AfterDeleteHook, categoryAfterDeleteHook)

	return &CategoryRepositorySQLBoiler{}
}

func (r *CategoryRepositorySQLBoiler) Exists(ctx context.Context, tx *sql.Tx, category *categories.Category) error {
	condition := repoModels.CategoryWhere.Name.EQ(category.Name().Value())

	if exists, err := repoModels.Categories(condition).Exists(ctx, tx); err != nil {
		return handler.DBErrHandler(err)
	} else if exists {
		return errs.NewCRUDError(fmt.Sprintf("category name %s already exists", category.Name().Value()))
	}

	return nil
}

func (r *CategoryRepositorySQLBoiler) Create(ctx context.Context, tx *sql.Tx, category *categories.Category) error {
	newCategory := repoModels.Category{
		ID:    0,
		ObjID: category.Id().Value(),
		Name:  category.Name().Value(),
	}

	if err := newCategory.Insert(ctx, tx, boil.Whitelist("obj_id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}

	return nil
}

func (r *CategoryRepositorySQLBoiler) UpdateById(ctx context.Context, tx *sql.Tx, category *categories.Category) error {
	oldCategory, err := repoModels.Categories(repoModels.CategoryWhere.ObjID.EQ(category.Id().Value())).One(ctx, tx)

	if err != nil {
		return handler.DBErrHandler(err)
	}

	if oldCategory == nil {
		return errs.NewCRUDError(fmt.Sprintf("category id %s not found", category.Id().Value()))
	}

	updateCategory := repoModels.Category{
		ID:    oldCategory.ID,
		ObjID: category.Id().Value(),
		Name:  category.Name().Value(),
	}

	if _, err := updateCategory.Update(ctx, tx, boil.Whitelist("obj_id", "name")); err != nil {
		return handler.DBErrHandler(err)
	}

	return nil
}

func (r *CategoryRepositorySQLBoiler) DeleteById(ctx context.Context, tx *sql.Tx, category *categories.Category) error {
	oldCategory, err := repoModels.Categories(repoModels.CategoryWhere.ObjID.EQ(category.Id().Value())).One(ctx, tx)

	if err != nil {
		return handler.DBErrHandler(err)
	}

	if oldCategory == nil {
		return errs.NewCRUDError(fmt.Sprintf("category id %s not found", category.Id().Value()))
	}

	if _, err = oldCategory.Delete(ctx, tx); err != nil {
		return handler.DBErrHandler(err)
	}

	return nil
}

// 登録処理後に実行されるフック
func categoryAfterInsertHook(ctx context.Context, exec boil.ContextExecutor, category *repoModels.Category) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを登録しました。\n", category.ObjID, category.Name)
	return nil
}

// 変更処理後に実行されるフック
func categoryAfterUpdateHook(ctx context.Context, exec boil.ContextExecutor, category *repoModels.Category) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを変更しました。\n", category.ObjID, category.Name)
	return nil
}

// 削除処理後に実行されるフック
func categoryAfterDeleteHook(ctx context.Context, exec boil.ContextExecutor, category *repoModels.Category) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを削除しました。\n", category.ObjID, category.Name)
	return nil
}
