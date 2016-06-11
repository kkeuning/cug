import {compose, createStore, applyMiddleware } from 'redux';

import rootReducer from '../reducers';
import thunkMiddleware from 'redux-thunk';

const createAppStore = compose(
	applyMiddleware(thunkMiddleware),
  window.devToolsExtension ? window.devToolsExtension() : f => f
)(createStore);

export default function configureStore(initialState) {
	return createAppStore(rootReducer, initialState);
};
