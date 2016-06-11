import React, {Component, PropTypes} from 'react';
import {connect} from 'react-redux';
import NavBar from '../components/NavBar';
import Spinner from '../components/Spinner';

import 'bootstrap/dist/css/bootstrap.css';
import 'app.css';

@connect(state => ({routerState: state.router, bottles: state.bottles }))
class App extends Component {
	constructor(props) {
		super(props);
	}

	static propTypes = {
		children: PropTypes.object.isRequired,
		bottles: PropTypes.object.isRequired
	};
	
	render() {
		const {children} = this.props;
		const {bottles} = this.props;

		return (
			<div>
				<NavBar />
				<div className='container'>
				{bottles.isLoading ? <Spinner /> : children}
				</div>
			</div>
		);
	}
}

export default App;