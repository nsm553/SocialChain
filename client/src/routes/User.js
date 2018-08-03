import React from 'react';
import { Link } from "react-router-dom";
import { Button } from 'semantic-ui-react';
import { Grid } from 'semantic-ui-react';
import { Image, List } from 'semantic-ui-react'
//import image from './images/bill-melinda-gates.jpeg';

export default () => (
	<div>
		<Grid.Row>
			<Link to="/company"><img className="back-button" src='/images/left-arrow-black.svg'/></Link>
			<div className="hero">	
				<div className="profile-container">
					<Image className="profile-black" src='/images/profile-black2.png'/>
				</div>
	    	</div>
	    </Grid.Row>
		<Grid centered columns={2}>
	    <Grid.Column>
	    	<h1 className="profile-title">Causes you have supported</h1>
            <h2> Current Balance: 8.42 ETH</h2>
			<h3> Bill and Melinda Gates Foundation</h3>
	    	<List divided verticalAlign='middle' className="margin-bottom-120">
			    <List.Item>
			      <List.Content floated='right'>
			        <p className="price">2 ETH</p>
			      </List.Content>
			      <Image className="donation-image" avatar src='/images/books.jpg' />
                    <List.Content>Books <a target="_blank" href="https://www.amazon.com/">(Amazon)</a></List.Content>
			    </List.Item>
			    <List.Item>
			      <List.Content floated='right'>
			        <p className="price">5 ETH</p>
			      </List.Content>
			      <Image className="donation-image" avatar src='/images/tables.jpeg' />
                    <List.Content>Tables <a target="_blank" href="https://www.ikea.com/">(IKEA)</a></List.Content>
			    </List.Item>

			  </List>
	    </Grid.Column>
	  </Grid>
	</div>
);


/*
	<div>	
		<div>lala</div>
		<ul>
		  <li>
		    <Link to="/">Home</Link>
		  </li>
		  <li>
		    <Link to="/company">Company</Link>
		  </li>
		</ul>
	</div>
*/



// export function Welcome(props) {
// 	alert("FOO");
//     // return <h1>Hello, {props.name}</h1>;
// }

// import React from 'react';

// export default () => (
// 	<div>	
// 		<div>auehauheauehauehauheuha</div>
// 	</div>
// );
