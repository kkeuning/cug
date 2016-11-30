import React, {Component, PropTypes} from 'react';
import {connect} from 'react-redux';
import Bottle from './Bottle';

@connect(state => ({data: state.bottles.data}))
class MBottles extends Component{
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
					{data.map(bottle => {
						return (
						<table className='table table-bordered table-striped'>
							<thead>
								<tr>
									<th>Column</th>
									<th>Value</th>
								</tr>
							</thead>
							<tbody>
									<tr>
										<td>id</td>
										<td>{bottle.id}</td>
									</tr>
									<tr>
										<td>account_id</td>
										<td>{bottle.account.id}</td>
									</tr>
									<tr>
										<td>name</td>
										<td>{bottle.name}</td>
									</tr>
									<tr>
										<td>varietal</td>
										<td>{bottle.varietal}</td>
									</tr>
									<tr>
										<td>vineyard</td>
										<td>{bottle.vineyard}</td>
									</tr>
									<tr>
										<td>vintage</td>
										<td>{bottle.vintage}</td>
									</tr>
							</tbody>
						</table>
						)
					})}
			</div>
		)
	}
}

export default MBottles;
