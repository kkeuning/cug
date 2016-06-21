import * as types from '../reduxa/cellarActionTypes';
import { combineReducers } from 'redux'; 

const bottleReducer = (state = {
	isLoading: false,
	data: false,
	message: false,
	status: false,
	error: false}
, action = null) => {
	switch(action.type) {
		case types.RECV_LIST_BOTTLES_ERROR:
			return {...state,
				isLoading: false,
				data: false,
				message: action.message,
				code: action.status,
				error: true};
		case types.RECV_LIST_BOTTLES_SUCCESS:
			return {...state,
				isLoading: false,
				data: action.data,
				code: action.status,
				message: false,
				error: false};
		case types.REQ_LIST_BOTTLES:
			return {...state,
        isLoading: true,
        message: false,
        code: false,
        error: false};
		default:
			return state;
	}
};

const rootReducer = combineReducers({
	bottles: bottleReducer
});

export default rootReducer;
