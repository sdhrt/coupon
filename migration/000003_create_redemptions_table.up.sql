CREATE TABLE IF NOT EXISTS redemptions (
    redemption_id CHAR(36) NOT NULL PRIMARY KEY,
    user_id CHAR(36) NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    coupon_id CHAR(36) NOT NULL REFERENCES coupons(coupon_id) ON DELETE CASCADE,
    redeemed_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    status VARCHAR(20) NOT NULL DEFAULT 'success', -- success, failed, expired
    UNIQUE(user_id, coupon_id)
);


