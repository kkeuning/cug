import 'babel-core/polyfill';
import React from 'react';
import ReactDOM from 'react-dom';
import {Provider} from 'react-redux';
import {Router, Route} from 'react-router';
import {browserHistory} from 'react-router';
import configureStore from './store';
import App from './containers/App';
import Bottles from './components/Bottles';
import MBottles from './components/MBottles';
import Error from './containers/Error';
import Home from './components/Home';
import {listBottles} from './reduxa/cellarActionCreators';

//const history = new browserHistory();
const store = configureStore();

const loadBottles = () => {
  store.dispatch(listBottles('/cellar/accounts/1/bottles'));
};
let bottlesView = Bottles;
if (screen.width < 700) {
    bottlesView = MBottles;
}
ReactDOM.render(
  <Provider store={store}>
    <Router history={browserHistory}>
      <Route component={App}>
        <Route path='/' component={Home}/>
        <Route path='/bottles' component={bottlesView} onEnter={loadBottles}/>
        <Route path='/error' component={Error}/>
      </Route>
    </Router>
  </Provider>,
  document.getElementById('root')
);
