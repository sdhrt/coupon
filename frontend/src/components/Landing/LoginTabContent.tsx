import { useRef } from "react";
import { Input } from "../ui/input";
import { Label } from "../ui/label";
import { Button } from "../ui/button";
import axios from "axios";
import { config } from "@/lib/config";
import { useNavigate } from "react-router";
import { useAuth } from "@/providers/auth-provider";
import { toast } from "sonner";

function LoginTabContent() {
  const emailRef = useRef<HTMLInputElement>(null);
  const passwordRef = useRef<HTMLInputElement>(null);
  const { setToken } = useAuth();
  const navigate = useNavigate();

  const loginForm = async (
    // @ts-ignore
    event: React.FormEvent<HTMLFormElement>,
  ) => {
    event.preventDefault();
    const email = emailRef.current?.value;
    const password = passwordRef.current?.value;
    try {
      const LoginResponse = await axios.post(
        `${config.server_url}/login`,
        {
          email: email,
          password: password,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        },
      );
      const { access_token } = LoginResponse.data;
      localStorage.setItem("access_token", access_token);
      setToken(access_token);
      return navigate("/home", {
        flushSync: true,
      });
    } catch (error: any) {
      toast.error(error.response.data.message || "error occured");
    }
  };

  return (
    <form className="flex flex-col gap-2 rounded-sm" onSubmit={loginForm}>
      <div className="font-semibold text-lg flex justify-center">
        Login to your account
      </div>
      <Label htmlFor="email">Email</Label>
      <Input ref={emailRef} id="email" />
      <Label htmlFor="password">password</Label>
      <Input ref={passwordRef} id="password" type={"password"} />
      <Button type="submit">Login</Button>
    </form>
  );
}

export default LoginTabContent;
