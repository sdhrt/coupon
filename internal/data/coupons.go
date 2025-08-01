package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type CouponModel struct {
	DB *sql.DB
}

type Coupon struct {
	Coupon_id         string
	User_id           string
	Created_at        string
	Expires_at        string
	Coupon_code       string
	Coupon_type       string
	Coupon_visibility string
	Coupon_usage      int
	Coupon_limit      int
	Coupon_value      float32
}

// CouponCode function generates a coupon from uppercase and numbers of length n
func CouponCode(n int) string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, n)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}

// Create_user method take email, name and password and registers it to the database
// returns the id of the coupon
func (m *CouponModel) Create_coupon(
	user_id, coupon_code, coupon_type, coupon_value, coupon_visibility string, expiry_date time.Time,
) (string, error) {
	id := uuid.New()

	if coupon_code == "" {
		coupon_code = CouponCode(8)
	}

	query := `
	INSERT INTO coupons(coupon_id, user_id, coupon_code, coupon_type, coupon_visibility, coupon_value, expires_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{id.String(), user_id, coupon_code, coupon_type, coupon_visibility, coupon_value, expiry_date.UTC()}
	_, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

// Get_all_coupons returns all the coupons associated with the supplied user_id
func (m *CouponModel) Get_all_coupons(user_id string) ([]Coupon, error) {
	var coupons []Coupon
	query := "SELECT * FROM coupons WHERE user_id = $1"
	rows, err := m.DB.Query(query, user_id)
	if err != nil {
		return coupons, err
	}
	defer rows.Close()

	for rows.Next() {
		var coupon Coupon
		if err := rows.Scan(
			&coupon.Coupon_id,
			&coupon.User_id,
			&coupon.Created_at,
			&coupon.Expires_at,
			&coupon.Coupon_code,
			&coupon.Coupon_type,
			&coupon.Coupon_visibility,
			&coupon.Coupon_usage,
			&coupon.Coupon_limit,
			&coupon.Coupon_value,
		); err != nil {
			return coupons, err
		}
		coupons = append(coupons, coupon)
	}
	if err = rows.Err(); err != nil {
		return coupons, err
	}
	return coupons, nil
}

// Redeem coupon redeems a coupon that is associated with the user_id.
// The function also increments the usage of the token, given that it has not crossed the limit
//
//	Takes
//	   user_id string
//	   coupon_code string
func (m *CouponModel) Redeem_coupon(userID string, couponCode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := m.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Get coupon details and lock row
	var couponID string
	var expiresAt time.Time
	var usage, limit int

	query := `
	SELECT coupon_id, expires_at, coupon_usage, coupon_limit
	FROM coupons
	WHERE coupon_code = $1
	FOR UPDATE`

	err = tx.QueryRowContext(ctx, query, couponCode).Scan(&couponID, &expiresAt, &usage, &limit)
	if err == sql.ErrNoRows {
		return errors.New("coupon not found")
	} else if err != nil {
		return err
	}

	if time.Now().After(expiresAt) {
		return errors.New("coupon has expired")
	}

	if usage >= limit {
		return errors.New("coupon usage limit reached")
	}

	var count int
	checkQuery := `
	SELECT COUNT(*) FROM redemptions
	WHERE user_id = $1 AND coupon_id = $2`

	err = tx.QueryRowContext(ctx, checkQuery, userID, couponID).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user has already redeemed this coupon")
	}

	// Insert redemption record
	redemption_id := uuid.New()
	insertQuery := `
	INSERT INTO redemptions(redemption_id, user_id, coupon_id, redeemed_at, status)
	VALUES ($1, $2, $3, NOW(), 'success')`

	_, err = tx.ExecContext(ctx, insertQuery, redemption_id.String(), userID, couponID)
	if err != nil {
		return err
	}

	// Increment coupon usage
	updateQuery := `
	UPDATE coupons
	SET coupon_usage = coupon_usage + 1
	WHERE coupon_id = $1`

	_, err = tx.ExecContext(ctx, updateQuery, couponID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		fmt.Printf("transaction commit failed: %s\n", err.Error())
		return err
	}

	return nil
}

// delete_coupon deletes coupons which matches user_id and coupon_id
func (m *CouponModel) Delete_coupon(user_id, coupon_id string) error {
	query := `
	DELETE FROM coupons WHERE user_id = $1 AND coupon_id = $2`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{user_id, coupon_id}
	row, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("No record deleted")
	}
	return nil
}
