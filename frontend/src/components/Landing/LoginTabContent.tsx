import { useRef } from "react";
import { Input } from "../ui/input";
import { Label } from "../ui/label";
import { Button } from "../ui/button";
import axios from "axios";
import { useNavigate } from "react-router";
import { toast } from "sonner";
import { config } from "@/lib/config";
import { useAuth } from "@/providers/auth-provider";

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
      setToken(access_token);
      localStorage.setItem("access_token", access_token);
      return navigate("/home", {
        flushSync: true,
      });
    } catch (error: any) {
      const { message } = error.response.data.message;
      toast.error(message ? message : "error occured");
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
