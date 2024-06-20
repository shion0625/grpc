package repository

import (
	"context"
	"database/sql"

	"github.com/shion0625/grpc/msa/internal/app/commandservice/domain/models/categories"
)

type Category interface {
	Exists(ctx context.Context, tx *sql.Tx, category *categories.Category) error
	Create(ctx context.Context, tx *sql.Tx, category *categories.Category) error
	UpdateById(ctx context.Context, tx *sql.Tx, category *categories.Category) error
	DeleteById(ctx context.Context, tx *sql.Tx, category *categories.Category) error
}
