import HomeSidebar from "@/components/Home/home-sidebar";
import { SidebarProvider } from "@/components/ui/sidebar";
import { useAuth } from "@/providers/auth-provider";
import { useEffect } from "react";
import { Outlet, useNavigate } from "react-router";

function HomeLayout() {
  const { access_token, setToken } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (access_token === "") {
      setToken("");
      navigate("/");
      return;
    }
  }, []);

  return (
    <SidebarProvider>
      <HomeSidebar />
      <main className="w-full">
        <Outlet />
      </main>
    </SidebarProvider>
  );
}

export default HomeLayout;
