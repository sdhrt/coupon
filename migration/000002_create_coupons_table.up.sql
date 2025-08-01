CREATE TABLE IF NOT EXISTS coupons (
    coupon_id CHAR(36) NOT NULL PRIMARY KEY,
    user_id CHAR(36) NOT NULL REFERENCES users ON DELETE CASCADE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    expires_at timestamp(0) with time zone NOT NULL,
    coupon_code VARCHAR(20) UNIQUE NOT NULL,
    coupon_type VARCHAR(20) NOT NULL,
    coupon_visibility VARCHAR(20),
    coupon_usage int NOT NULL DEFAULT 0,
    coupon_limit int NOT NULL DEFAULT 50,
    coupon_value float NOT NULL DEFAULT 0.0
);
