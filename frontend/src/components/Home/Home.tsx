import CouponDialog from "./components/CouponDialog";
import CouponsTable from "./components/CouponsTable";

function Home() {
  return (
    <div className="flex flex-col justify-center items-center w-full mt-2">
      <CouponDialog />
      <CouponsTable />
    </div>
  );
}

export default Home;
