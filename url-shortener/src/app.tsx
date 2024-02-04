import React from 'react';
import { Provider } from 'react-redux';
import Router from './Router';
import { store } from './store/store';

export default function App(): React.ReactElement {
  return (
    <Provider store={store}>
      <Router />
    </Provider>
  );
}
