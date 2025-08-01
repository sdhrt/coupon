import { createBrowserRouter } from "react-router";
import Home from "./components/Home/Home";
import Landing from "./components/Landing/Landing";
import HomeLayout from "./layouts/HomeLayout";
import Redeem from "./components/Redeem/Redeem";
import Redemptions from "./components/Redemptions/Redemptions";

export const router = createBrowserRouter([
  {
    path: "/",
    Component: Landing,
  },
  {
    path: "home",
    Component: HomeLayout,
    children: [
      {
        path: "",
        Component: Home,
      },
      {
        path: "redeem",
        Component: Redeem,
      },
      {
        path: "redemptions",
        Component: Redemptions,
      },
    ],
  },
  {
    path: "*",
    element: (
      <div className="font-bold tracking-wider text-xl">404 not found</div>
    ),
  },
]);
