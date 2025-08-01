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

function Redemptions() {
  const [redemptions, setRedemptions] = useState([]);
  useEffect(() => {
    const redemptionFetch = async () => {
      try {
        const redemptionResponse = await api.post("/coupon/redemptions");
        console.log(redemptionResponse);
        if (redemptionResponse.status == 200) {
          setRedemptions(redemptionResponse.data.redemptions);
        }
      } catch (error) {
        console.log("error");
        console.log(error);
        toast.error("Error occured");
      }
    };
    redemptionFetch();
  }, []);

  return (
    <div className="flex flex-col h-full m-2 gap-4">
      <Table>
        <TableCaption>A list of your recent redemptions.</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[10px]">S. No.</TableHead>
            <TableHead>Code</TableHead>
            <TableHead>Email</TableHead>
            <TableHead>Name</TableHead>
            <TableHead>Redeemed at</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {redemptions &&
            redemptions.map((redemption: any, index: number) => {
              return (
                <TableRow key={redemption.Coupon_id}>
                  <TableCell>{index + 1}</TableCell>
                  <TableCell>{redemption.coupon_code}</TableCell>
                  <TableCell>{redemption.user_email}</TableCell>
                  <TableCell>{redemption.user_name}</TableCell>
                  <TableCell>
                    {new Date(redemption.redeemed_at).toLocaleDateString()}
                  </TableCell>
                </TableRow>
              );
            })}
        </TableBody>
      </Table>
    </div>
  );
}

export default Redemptions;
