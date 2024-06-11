package repository

import (
	"context"
	"database/sql"

	"github.com/shion0625/grpc/msa/internal/app/commandservice/domain/models/categories"
)

type Category interface {
	Exist(ctx context.Context, tx *sql.Tx, category categories.Category) (bool, error)
	Create(ctx context.Context, tx *sql.Tx, category *categories.Category) error
	UpdateByID(ctx context.Context, tx *sql.Tx, category *categories.Category) error
	DeleteByID(ctx context.Context, tx *sql.Tx, category *categories.Category) error
}
