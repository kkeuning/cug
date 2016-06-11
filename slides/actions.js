//START OMIT
import * as types from './actionTypes';
import axios from 'axios';
const requestBottleData = () => ({
  type: types.REQ_BOTTLE_DATA
});
const receiveBottleData = (json) => ({
  type: types.RECV_BOTTLE_DATA, data: json
});
const receiveBottleError = (json) => ({
  type: types.RECV_BOTTLE_ERROR, data: json
});
export const fetchBottleData = (url) => {
  return (dispatch) => {
      dispatch(requestBottleData());
      return axios({url: url, timeout: 20000, method: 'get', responseType: 'json'})
        .then((response) => {
	        dispatch(receiveBottleData(response.data));
	      })
        .catch((response) => {
	        dispatch(receiveBottleError(response.data));
	      });
    };
};
//END OMIT
