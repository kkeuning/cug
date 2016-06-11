import * as types from '../reduxa/cellarActionTypes';
import { combineReducers } from 'redux'; 


const bottleReducer = (state = {
	isLoading: false,
	data: false,
	message: false,
	code: false,
	error: false}
, action = null) => {
	switch(action.type) {
		case types.RECV_LIST_BOTTLE_ERROR:
			return {...state, isLoading: false, data: false, message: action.message, code: action.code, error: true};
		case types.RECV_LIST_BOTTLE_SUCCESS:
			return {...state, isLoading: false, data: action.data, message: false, error: false};
		case types.REQ_LIST_BOTTLE:
			return {...state, isLoading: true, message: false, code: action.code, error: false};
		default:
			return state;
	}
};

const rootReducer = combineReducers({
	bottles: bottleReducer
});

export default rootReducer;
