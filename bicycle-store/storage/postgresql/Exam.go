package postgresql

import (
	"app/api/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type R_Repo struct {
	db *pgxpool.Pool
}

func NewCodeRepo(db *pgxpool.Pool) *R_Repo {
	return &R_Repo{db: db}
}

func (r R_Repo) PromoView(ctx context.Context, req *models.StigmaApi) (res models.SigmaSql, err error) {

	query := `select order_id, sum(list_price) AS "list_price" , sum(discount) AS "discount"
from order_items
WHERE order_id = $1 GROUP BY  order_id`

	err = r.db.QueryRow(ctx, query, req.Order_id).Scan(
		&res.Order_id,
		&res.List_price,
		&res.Discount,
	)

	if err != nil {
		return res, err
	}

	if req.Promo_Code == "" {
		return res, nil
	}

	res.List_price -= res.Discount

	return res, nil
}
