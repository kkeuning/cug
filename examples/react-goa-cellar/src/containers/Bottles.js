import React, {Component, PropTypes} from 'react';
import {connect} from 'react-redux';
import Bottles from '../components/Bottles';
import MBottles from '../components/mobile/Bottles';

@connect(state => ({routerState: state.router, bottles: state.bottles }))
class BottlesContainer extends Component {
	constructor(props) {
		super(props);
	}

	render() {
		console.log('inside the Bottles container');
		console.log(screen.width);
		const BottlesView = (screen.width < 700) ? MBottles : Bottles;
		console.log(BottlesView);
		return (
			<div>
				<BottlesView />
			</div>
		);
	}
}

export default BottlesContainer;
