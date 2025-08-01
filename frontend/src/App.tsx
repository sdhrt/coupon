import { RouterProvider } from "react-router";
import { router } from "./routes";
import AuthProvider from "./providers/auth-provider";
import { Toaster } from "sonner";

function App() {
  return (
    <AuthProvider>
      <RouterProvider router={router} />
      <Toaster />
    </AuthProvider>
  );
}

export default App;
