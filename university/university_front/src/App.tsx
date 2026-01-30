import { createBrowserRouter } from "react-router"
import { RouterProvider } from "react-router/dom"
import { LoginPage } from "./pages/login/page.tsx"
import { RegisterPage } from "./pages/register/page.tsx"

const router = createBrowserRouter([
    {
        path: "/",
        element: <div>Hello World!</ div >,
    },
    {
        path: "/login",
        element: <LoginPage />,
    },
    {
        path: "register",
        element: <RegisterPage />,
    },
]);


function App() {
    return (
        <RouterProvider router={router} />
    );
}

export default App
