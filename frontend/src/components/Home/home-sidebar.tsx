import { ChartNetwork, LogOutIcon, Ticket, TicketCheck } from "lucide-react";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarTrigger,
  useSidebar,
} from "@/components/ui/sidebar";
import { Link } from "react-router";
import { Button } from "../ui/button";
import { useAuth } from "@/providers/auth-provider";

const items = [
  {
    title: "Home",
    url: "/home",
    icon: Ticket,
  },
  {
    title: "Redeem",
    url: "/home/redeem",
    icon: TicketCheck,
  },
  {
    title: "Redemptions",
    url: "/home/redemptions",
    icon: ChartNetwork,
  },
];

function HomeSidebar() {
  const { state } = useSidebar();
  const { setToken } = useAuth();

  const handleLogout = () => {
    localStorage.clear();
    setToken("");
    window.location.reload();
  };
  return (
    <Sidebar collapsible="icon">
      <SidebarHeader>
        {state !== "collapsed" ? (
          <div className="flex justify-between">
            <span className="font-semibold text-xl">Coupons</span>
            <SidebarTrigger />
          </div>
        ) : (
          <SidebarTrigger />
        )}
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <Link to={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <Button onClick={handleLogout}>
          {state == "collapsed" ? (
            <div>
              <LogOutIcon />
            </div>
          ) : (
            <span>Logout</span>
          )}
        </Button>
      </SidebarFooter>
    </Sidebar>
  );
}

export default HomeSidebar;
