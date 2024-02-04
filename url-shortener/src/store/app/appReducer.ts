import { createSlice } from '@reduxjs/toolkit';
import { RootState } from '../store';

export type IAppMeta = {
  meta: {};
};

const initialState: IAppMeta = {
  meta: {},
};

export const appReducer = createSlice({
  name: 'app',
  initialState,
  reducers: {
    setAppMeta: (state, data) => {
    },
  },
});

export const { setAppMeta } = appReducer.actions;

export const selectAppState = (state: RootState) => state.app;

export default appReducer.reducer;
