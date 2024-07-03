package repository

import (
	"context"
	"database/sql"

	"github.com/shion0625/grpc/msa/internal/app/commandservice/domain/models/products"
)

type Product interface {
	Exists(ctx context.Context, tx *sql.Tx, product *products.Product) error
	Create(ctx context.Context, tx *sql.Tx, product *products.Product) error
	UpdateById(ctx context.Context, tx *sql.Tx, product *products.Product) error
	DeleteById(ctx context.Context, tx *sql.Tx, product *products.Product) error
}
