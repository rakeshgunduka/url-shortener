import React from 'react';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';
import Analytics from './pages/Analytics/Analytics';
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
      path: '/404',
      element: <UrlNotFound />,
    },
    {
      path: '/analytics',
      element: <Analytics />,
    },
  ]);

  return <RouterProvider router={router} />;
}

export default Router;
