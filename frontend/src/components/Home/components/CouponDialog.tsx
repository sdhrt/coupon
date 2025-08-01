import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useRef, useState } from "react";
import { api } from "@/lib/api";
import { toast } from "sonner";

function CouponDialog() {
  const [dialogOpen, setDialogOpen] = useState(false);
  const codeRef = useRef<HTMLInputElement>(null);
  const [couponType, setCouponType] = useState("fixed");
  const [visibility, setVisibility] = useState("public");
  const valueRef = useRef<HTMLInputElement>(null);
  const daysRef = useRef<HTMLInputElement>(null);

  const couponForm = async (
    // @ts-ignore
    event: React.FormEvent<HTMLFormElement>,
  ) => {
    event.preventDefault();
    const code = codeRef.current?.value;
    const type = couponType;
    const value = valueRef.current?.value;
    const days = daysRef.current?.value;
    if (code == "" || !value || Number(days) < 1) {
      toast.error("Missing Values");
      return;
    }
    try {
      const CouponResponse = await api.post(
        "/coupon/create",
        {
          coupon_code: code,
          coupon_type: type,
          coupon_value: value,
          coupon_visibility: visibility,
          coupon_duration: days,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        },
      );
      console.log(CouponResponse.data);
      toast.message(CouponResponse.data.message || "Added coupon");
      setDialogOpen(false);
    } catch (error: any) {
      toast.error(error.response.data.message || "Error adding coupon");
      console.log(error);
    }
  };
  return (
    <Dialog open={dialogOpen}>
      <DialogTrigger asChild>
        <Button onClick={() => setDialogOpen(true)} variant={"outline"}>
          Add Coupon
        </Button>
      </DialogTrigger>
      <DialogContent className="flex flex-col items-center">
        <DialogHeader>
          <DialogTitle>Add coupon?</DialogTitle>
        </DialogHeader>
        <form
          onSubmit={couponForm}
          className="w-[300px] flex flex-col gap-2 items-center"
        >
          <Input ref={codeRef} id="code" placeholder="e.g. GET10OFF" />
          <Label htmlFor="code" className="text-xs text-muted-foreground">
            leave blank for generated code
          </Label>
          <div className="flex flex-row gap-2">
            <Input ref={valueRef} placeholder="Value" />
            <Select
              onValueChange={(type) => setCouponType(type)}
              defaultValue={couponType}
            >
              <SelectTrigger className="w-[50%]">
                <SelectValue placeholder="Type" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="percentage">Percentage</SelectItem>
                <SelectItem value="fixed">Fixed</SelectItem>
              </SelectContent>
            </Select>
          </div>
          <div className="flex w-full justify-between *:w-[50%]">
            <Label htmlFor="visibility">Visibility: </Label>
            <Select
              onValueChange={(visibility) => setVisibility(visibility)}
              defaultValue={visibility}
            >
              <SelectTrigger className="" id="visibility">
                <SelectValue placeholder="public" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="public">Public</SelectItem>
                <SelectItem value="private">Private</SelectItem>
              </SelectContent>
            </Select>
          </div>
          <Input
            ref={daysRef}
            placeholder="expires after ? days"
            type="number"
          />
          <div className="flex gap-2">
            <Button type="reset" variant={"destructive"}>
              Reset
            </Button>
            <Button type="submit">Add</Button>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  );
}

export default CouponDialog;
