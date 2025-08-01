package data

import (
	"context"
	"database/sql"
	"time"
)

type RedemptionModel struct {
	DB *sql.DB
}

type RedemptionRecord struct {
	RedemptionID string    `json:"redemption_id"`
	RedeemedAt   time.Time `json:"redeemed_at"`
	UserName     string    `json:"user_name"`
	UserEmail    string    `json:"user_email"`
	CouponCode   string    `json:"coupon_code"`
	Status       string    `json:"status"`
}

// Returns all the coupons redemption status that is associated with the user_id
func (m *RedemptionModel) GetAllRedemptions(user_id string) ([]RedemptionRecord, error) {
	query := `
	SELECT 
		r.redemption_id,
		r.redeemed_at,
		u.name AS user_name,
		u.email AS user_email,
		c.coupon_code,
		r.status
	FROM redemptions r
	INNER JOIN users u ON r.user_id = u.user_id
	INNER JOIN coupons c ON r.coupon_id = c.coupon_id
	WHERE r.user_id = $1
	ORDER BY r.redeemed_at DESC;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var redemptions []RedemptionRecord

	for rows.Next() {
		var r RedemptionRecord
		err := rows.Scan(
			&r.RedemptionID,
			&r.RedeemedAt,
			&r.UserName,
			&r.UserEmail,
			&r.CouponCode,
			&r.Status,
		)
		if err != nil {
			return nil, err
		}
		redemptions = append(redemptions, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return redemptions, nil
}
