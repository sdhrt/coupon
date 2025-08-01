import { useRef } from "react";
import { Input } from "../ui/input";
import { Label } from "../ui/label";
import { Button } from "../ui/button";
import axios from "axios";
import { toast } from "sonner";

function RegisterTabContent() {
    const formRef = useRef<HTMLFormElement>(null);
    const emailRef = useRef<HTMLInputElement>(null);
    const nameRef = useRef<HTMLInputElement>(null);
    const passwordRef = useRef<HTMLInputElement>(null);

    const registerForm = async (
        // @ts-ignore
        event: React.FormEvent<HTMLFormElement>,
    ) => {
        event.preventDefault();

        const name = nameRef.current?.value;
        const email = emailRef.current?.value;
        const password = passwordRef.current?.value;

        try {
            const registerResponse = await axios.post(
                `http://localhost:8000/register`,
                {
                    email: email,
                    name: name,
                    password: password,
                },
                {
                    headers: {
                        "Content-Type": "application/json",
                    },
                },
            );
            console.log("registerResponse: ", registerResponse);
            toast.message("Successfully registered", {
                description: "Please continue to login",
            });
            if (formRef.current) {
                formRef.current.reset();
            }
        } catch (error: any) {
            console.log("error");
            console.log(error.response);
            toast.error(error.response.data.message || "Couldn't register user")
        }
    };

    return (
        <form
            className="flex flex-col gap-2 rounded-sm"
            onSubmit={registerForm}
            ref={formRef}
        >
            <div className="font-semibold text-lg flex justify-center">
                Register your account
            </div>
            <Label htmlFor="email">Email</Label>
            <Input ref={emailRef} id="email" type="email" autoComplete="off" />
            <Label htmlFor="name">Name</Label>
            <Input ref={nameRef} id="name" autoComplete="off" />
            <Label htmlFor="password">Password</Label>
            <Input ref={passwordRef} id="password" type={"password"} />
            <Button type="submit">Register</Button>
        </form>
    );
}

export default RegisterTabContent;
