import React from 'react';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';
import Home from './pages/Home/Home';
import UrlNotFound from './pages/UrlNotFound/UrlNotFound';
import Urls from './pages/Urls/Urls';


function Router(): React.ReactElement {
  const router = createBrowserRouter([
    {
      path: '/',
      element: <Home />,
    },
    {
      path: '/urls',
      element: <Urls />,
    },
    {
      path: '/not-found',
      element: <UrlNotFound />,
    },
  ]);

  return <RouterProvider router={router} />;
}

export default Router;
