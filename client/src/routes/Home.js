import React from 'react';
import { Link } from "react-router-dom";
import { Button } from 'semantic-ui-react';
import { Grid } from 'semantic-ui-react';
import { Image, List } from 'semantic-ui-react'

export default () => (
	<div className="block-blackground">	
		<Grid centered columns={2} className="principal">
	    <div className="gray-background"></div>
	    <Grid.Column>
	      <Image src='/images/principal2.png'/>
	    </Grid.Column>
	  </Grid>
		<Grid centered columns={2}>
	    <Grid.Column>
				<Grid.Row>
					<Link to="/company">	
						<div className="hero card">	
							<div className="dark-frame"></div>
							<Image className="image-hero" src='/images/bill-melinda-gates.jpg' size='massive'/>
				    	<h1 className="image-hero-text">Bill & Melinda Gates Foundation</h1>
			    	</div>
		    	</Link>
		    	<div className="hero card">	
						<div className="dark-frame"></div>
						<Image className="image-hero" src='/images/american-civil-liberties-union.jpg' size='massive' />
			    	<h1 className="image-hero-text">American Civil Liberties Union</h1>
		    	</div>
		    	<div className="hero card">	
						<div className="dark-frame"></div>
						<Image className="image-hero" src='/images/teach-for-america.jpg' size='massive' />
			    	<h1 className="image-hero-text">Teach for America</h1>
		    	</div>
		    	<div className="hero card">	
						<div className="dark-frame"></div>
						<Image className="image-hero" src='/images/women-who-code.jpg' size='massive' />
			    	<h1 className="image-hero-text">Women Who code</h1>
		    	</div>
		    	<div className="hero card">	
						<div className="dark-frame"></div>
						<Image className="image-hero" src='/images/raices.jpg' size='massive' />
			    	<h1 className="image-hero-text">Raices</h1>
		    	</div>
		    </Grid.Row>
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

