//START OMIT
// List all bottles in account optionally filtering by year
// path is the request path, the format is "/cellar/accounts/:accountID/bottles"
// years is used to build the request query string.
// This function returns a promise which dispatches an error if the HTTP response is a 4xx or 5xx.
export const listBottles = (path, years) => {
  return (dispatch) => {
    dispatch(actions.requestListBottles());
    return axios({
      timeout: 20000,
      url: 'http://cellar.goa.design' + path,
      method: 'get',
      params: {
        years: years
      },
      responseType: 'json'
    })
      .then((response) => {
        dispatch(actions.receiveListBottlesSuccess(response.data, response.status));
      })
      .catch((response) => {
        dispatch(actions.receiveListBottlesError(response.data, response.status));
      });
  };
};
//END OMIT
