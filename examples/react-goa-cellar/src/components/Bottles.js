import React, {Component, PropTypes} from 'react';
import {connect} from 'react-redux';
import Bottle from './Bottle';

@connect(state => ({data: state.bottles.data}))
class Bottles extends Component{
	constructor(props){
		super(props);
	}
  
	static propTypes = {
		data: PropTypes.array.isRequired,
    dispatch: PropTypes.func
	};
  
	render() {

		const {data} = this.props;

		return (
			<div className='container'>
				<table className='table table-bordered table-striped'>
					<thead>
						<tr>
							<th>id</th>
							<th>account_id</th>
							<th>name</th>
							<th>varietal</th>
							<th>vineyard</th>
							<th>vintage</th>
						</tr>
					</thead>
					<tbody>
					{data.map(bottle => {
						return (
							<Bottle key={bottle.id} bottle={bottle} />
						)
					})}
					</tbody>
				</table>
			</div>
		)
	}
}

export default Bottles;
