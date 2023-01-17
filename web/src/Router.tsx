import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import BaseLayout from './pages/BaseLayout';
import ErrorPage from './pages/errors/NotFound';

const router = createBrowserRouter([
  {
    path: "/",
    element: <BaseLayout />,
    errorElement: <ErrorPage />,
  }
]);

export default function Router() {
  return <RouterProvider router={router} />
}
