import { api } from "@/lib/api";
import { useEffect, useState } from "react";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Button } from "@/components/ui/button";
import { Trash } from "lucide-react";
import { toast } from "sonner";

function CouponsTable() {
  const [coupons, setCoupons] = useState([]);
  useEffect(() => {
    const couponResponse = async () => {
      try {
        const couponResponse = await api.post("/coupon/get");
        if (couponResponse.status == 200) {
          setCoupons(couponResponse.data.coupons);
        }
      } catch (error) {
        console.log("error");
        console.log(error);
      }
    };
    couponResponse();
  }, []);

  const handleCouponDelete = async (id: string) => {
    try {
      const deleteResponse = await api.post(
        "/coupon/delete",
        {
          coupon_id: id,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        },
      );
      toast.message(deleteResponse.data.message);
    } catch (error) {
      console.log(error);
      toast.error("error occured while deleting coupon");
    }
  };
  return (
    <div className="flex flex-col h-full m-2 gap-4">
      <Table>
        <TableCaption>A list of your recent coupons.</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[10px]">S. No.</TableHead>
            <TableHead>Code</TableHead>
            <TableHead>Type</TableHead>
            <TableHead>Value</TableHead>
            <TableHead>Visibility</TableHead>
            <TableHead>Created at</TableHead>
            <TableHead>Expires at</TableHead>
            <TableHead>Delete</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {coupons &&
            coupons.map((coupon: any, index: number) => {
              return (
                <TableRow key={coupon.Coupon_id}>
                  <TableCell>{index + 1}</TableCell>
                  <TableCell>{coupon.Coupon_code}</TableCell>
                  <TableCell>{coupon.Coupon_type}</TableCell>
                  <TableCell>{coupon.Coupon_value}</TableCell>
                  <TableCell>{coupon.Coupon_visibility}</TableCell>
                  <TableCell>
                    {new Date(coupon.Created_at).toLocaleDateString()}
                  </TableCell>
                  <TableCell>
                    {new Date(coupon.Expires_at).toLocaleDateString()}
                  </TableCell>
                  <TableCell>
                    <Button
                      onClick={() => handleCouponDelete(coupon.Coupon_id)}
                      variant={"destructive"}
                      size={"sm"}
                    >
                      <Trash />
                    </Button>
                  </TableCell>
                </TableRow>
              );
            })}
        </TableBody>
      </Table>
    </div>
  );
}

export default CouponsTable;
