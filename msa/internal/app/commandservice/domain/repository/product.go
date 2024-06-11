package repository

import (
	"context"
	"database/sql"

	"github.com/shion0625/grpc/msa/internal/app/commandservice/domain/models/products"
)

type Product interface {
	Exists(ctx context.Context, tx *sql.Tx, product *products.Product) error
	Create(ctx context.Context, tx *sql.Tx, product *products.Product) error
	UpdateByID(ctx context.Context, tx *sql.Tx, product *products.Product) error
	DeleteByID(ctx context.Context, tx *sql.Tx, product *products.Product) error
}
