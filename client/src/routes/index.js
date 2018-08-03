import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import Home from './Home'
import Company from './Company'
import User from './User'

export default () => (
	<BrowserRouter>
		<Switch>
			<Route path="/" exact component={Home}/>
			<Route path="/company" component={Company}/>
			<Route path="/user" component={User}/>
		</Switch>
	</BrowserRouter>
);
