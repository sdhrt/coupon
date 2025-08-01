import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import LoginTabContent from "./LoginTabContent";
import RegisterTabContent from "./RegisterTabContent";
import { useEffect } from "react";
import { useAuth } from "@/providers/auth-provider";

function Landing() {
  const { access_token, setToken } = useAuth();
  useEffect(() => {
    if (access_token !== "") {
      localStorage.clear();
      setToken("");
    }
  }, []);
  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1 className="text-xl font-bold my-8">Coupon management system</h1>
      <Tabs
        defaultValue={"login"}
        className="w-[400px] *:w-full flex items-center"
      >
        <TabsList className="">
          <TabsTrigger value="login">Login</TabsTrigger>
          <TabsTrigger value="register">Register</TabsTrigger>
        </TabsList>
        <TabsContent value="login">
          <LoginTabContent />
        </TabsContent>
        <TabsContent value="register">
          <RegisterTabContent />
        </TabsContent>
      </Tabs>
    </div>
  );
}

export default Landing;
