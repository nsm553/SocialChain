import React from 'react';
import { Link } from "react-router-dom";
import { Button } from 'semantic-ui-react';
import { Grid } from 'semantic-ui-react';
import { Image, List } from 'semantic-ui-react'
//import image from './images/bill-melinda-gates.jpeg';

export default () => (
	<div>	
		<Grid.Row>
			<Link to="/"><img className="back-button" src='/images/left-arrow.svg'/></Link>
			<Link to="/user"><img className="profile-button" src='/images/profile.svg'/></Link>
			<div className="hero">	
				<div className="dark-frame"></div>
				<Image className="image-hero" src='/images/bill-melinda-gates.jpg' size='massive' />
	    	<h1 className="image-hero-text">Bill & Melinda Gates Foundation</h1>
    	</div>
    </Grid.Row>
		<Grid centered columns={2}>
	    <Grid.Column>
	    	<p className="body">Bill & Melinda Gates Foundation (BMGF), also known as the Gates Foundation, 
	    	is a private foundation founded by Bill and Melinda Gates. It was launched in 2000, 
	    	and is said to be the largest private foundation in the US, holding $38 billion in assets.
	    	The primary aims of the foundation are, globally, to enhance healthcare and reduce extreme poverty, and in America, 
	    	to expand educational opportunities and access to information technology. The foundation, based in Seattle, Washington, 
	    	is controlled by its three trustees: Bill and Melinda Gates, and Warren Buffett. Other principal officers include Co-Chair 
	    	William H. Gates, Sr. and Chief Executive Officer Susan Desmond-Hellmann.</p>
	    	<List divided verticalAlign='middle' className="margin-bottom-120">
			    <List.Item>
			      <List.Content floated='right'>
			        <Button primary onClick={()=>{
                        alert("You have successfully donated 2 ETH to Bill and Melinda Gates Foundation!")
                        return fetch('http://localhost:8080/api/completerequest/0', {
                            method: 'POST',
                                headers: {
                                    Accept: 'application/json',
                                    'Content-Type': 'application/json',
                                },
                                body: JSON.stringify({
                                    foo: 'bar',
                                }),
                        })
                            .then((response) => response.json())
                            .then((responseJson) => {
                                console.log(responseJson)
                            })
                            .catch((error) => {
                                console.error(error);
                            });

                    }}>Donate</Button>
			      </List.Content>
			      <List.Content floated='right'>
			        <p className="price">2 ETH</p>
			      </List.Content>
			      <Image className="donation-image" avatar src='/images/books.jpg' />
			      <List.Content>Books <a target="_blank" href="https://www.amazon.com/">(Amazon)</a></List.Content>
			    </List.Item>
			    <List.Item>
			      <List.Content floated='right'>
			        <Button primary>Donate</Button>
			      </List.Content>
			      <List.Content floated='right'>
			        <p className="price">5 ETH</p>
			      </List.Content>
			      <Image className="donation-image" avatar src='/images/tables.jpeg' />
			      <List.Content>Tables <a target="_blank" href="https://www.ikea.com/">(IKEA)</a></List.Content>
			    </List.Item>
			    <List.Item>
			      <List.Content floated='right'>
			        <Button primary>Donate</Button>
			      </List.Content>
			      <List.Content floated='right'>
			        <p className="price">1 ETH</p>
			      </List.Content>
			      <Image className="donation-image" avatar src='/images/paper.jpg' />
			      <List.Content>School stationery <a target="_blank" href="https://www.staples.com/">(Staples)</a></List.Content>
			    </List.Item>
			    <List.Item>
			      <List.Content floated='right'>
			        <Button primary>Donate</Button>
			      </List.Content>
			      <List.Content floated='right'>
			        <p className="price">3 ETH</p>
			      </List.Content>
			      <Image className="donation-image" avatar src='/images/pens.jpeg' />
			      <List.Content>Pens <a target="_blank" href="https://www.homedepot.com/">(Staples)</a></List.Content>
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
