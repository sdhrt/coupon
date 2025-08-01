package data

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Models struct {
	Users       UserModel
	Coupons     CouponModel
	Redemptions RedemptionModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:       UserModel{DB: db},
		Coupons:     CouponModel{DB: db},
		Redemptions: RedemptionModel{DB: db},
	}
}
