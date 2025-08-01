import { createContext, useContext, useEffect, useMemo, useState } from "react";
import { api } from "../lib/api";

type AuthProviderProps = {
  children: React.ReactNode;
};

type AuthProviderState = {
  access_token: string | null;
  setToken: (token: string) => void;
};

const AuthContext = createContext<AuthProviderState | undefined>(undefined);

const AuthProvider = ({ children, ...props }: AuthProviderProps) => {
  const [access_token, setToken_] = useState<string | null>(
    localStorage.getItem("access_token") || "",
  );

  const setToken = (token: string) => {
    localStorage.setItem("access_token", token);
    setToken_(token);
  };

  useEffect(() => {
    if (access_token) {
      api.defaults.headers.common["Authorization"] = "Bearer " + access_token;
    } else {
      delete api.defaults.headers.common["Authorization"];
    }
  }, [access_token]);

  const contextValue = useMemo(
    () => ({
      access_token: access_token,
      setToken: setToken,
    }),
    [access_token],
  );

  return (
    <AuthContext.Provider {...props} value={contextValue}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined)
    throw new Error("useAuth must be used inside AuthProvider");
  return context;
};

export default AuthProvider;
