"use client";

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { api } from "@/lib/api"; // axios instance
import { toast } from "sonner";

function Redeem() {
  const [couponCode, setCouponCode] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!couponCode.trim()) {
      toast.error("Invalid input", {
        description: "Please enter a valid coupon code.",
      });
      return;
    }

    setLoading(true);
    try {
      const res = await api.post("/coupon/redeem", {
        coupon_code: couponCode.trim(),
      });

      toast.success("Coupon Redeemed", {
        description: res.data.message || "Success!",
      });
      setCouponCode("");
    } catch (err: any) {
      toast.error("Redemption Failed", {
        description: err.response?.data?.message || "Something went wrong.",
      });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="grid place-items-center h-screen">
      <form
        onSubmit={handleSubmit}
        className="space-y-4 max-w-sm flex flex-col justify-center"
      >
        <Input
          type="text"
          placeholder="Enter coupon code"
          value={couponCode}
          onChange={(e) => setCouponCode(e.target.value)}
          disabled={loading}
        />
        <Button type="submit" disabled={loading}>
          {loading ? "Redeeming..." : "Redeem"}
        </Button>
      </form>
    </div>
  );
}

export default Redeem;
