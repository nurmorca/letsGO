package infra

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

var INSERT_PRODUCTS = `INSERT INTO products (name, price, discount, store)
VALUES
('batmobile', 8888.8, 0.0, 'wayne enterprises'),
('glasses', 100.0, 8.0, 'kents'),
('cat food (from the owner)', 9000.0, 0.0, 'selina kyle'),
('typewriter', 500.0, 12.0, 'lois lane');`

func TestDataInsert(ctx context.Context, dbPool *pgxpool.Pool) {
	insertProductResult, insertProductsErr := dbPool.Exec(ctx, INSERT_PRODUCTS)
	if insertProductsErr != nil {
		log.Error(insertProductsErr)
	} else {
		log.Info(fmt.Sprintf("products data crated with %d rows", insertProductResult.RowsAffected()))
	}
}
