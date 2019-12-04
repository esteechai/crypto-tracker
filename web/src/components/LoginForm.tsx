import React, { useState } from 'react'
import { StoreContainer } from '../store'
import { render} from 'react-dom'
import {Input, Container, Grid, GridColumn, Segment} from 'semantic-ui-react'
import {Link, Redirect} from 'react-router-dom'

interface Props{}

const LoginForm: React.FC<Props> = () => {
    const store = StoreContainer.useContainer()
    if(store.isLogin){
       // console.log("sucessfully login")
       return <Redirect to= '/dashboard' />
    }

  
    store.useFetchProducts("http://localhost:8080/api/get/products")
    
return(
   <Container className="ui grid" style={{backgroundColor: ''}} textAlign="center">
           <Grid verticalAlign="middle">                       
             <Grid.Column style={{width:450}}>                      
                <form className="ui form">
                    <Segment stacked>
                    <h2> Login</h2>
                    <div className={'ui erorr message' + (store.isError && store.errorMsg ? 'has-error' : '')}>            
                        {store.isError && store.errorMsg &&
                            <p>{store.errorMsg}</p>}      
                    </div>

                    <div className={'field' + (store.isSubmit && !store.enteredEmail ? 'has-error' : '')}>
                        <label> Email Address</label>
                        <input placeholder="Email Address" id="email" value={store.enteredEmail} onChange={store.handleEnteredEmail} />
                        {store.isSubmit && !store.enteredEmail && 
                            <div className="input-feedback"> Email is required </div>}
                    </div>
                
                    <div className={'field' + (store.isSubmit && !store.enteredPassword ? 'has-error' : '')}>
                    <label> Password</label>
                    <input placeholder="Password" id="password" type="password" value={store.enteredPassword} onChange={store.handleEnteredPassword}/>
                     {store.isSubmit && !store.enteredPassword && 
                        <div className="input-feedback"> Password is required </div>}
                    </div>
                
                    <button type="submit" className="ui primary button" onClick={store.handleLogin}>Log in</button>
                    <Link to = "/NavBar">Forgot password?</Link>
                    <Link to ="/cryptoList">Crypto List</Link>
                    <Link to="/signupForm"><p>Register</p></Link> 
                    
                    </Segment>
                </form>
         </Grid.Column>
         </Grid>
    </Container> 
)
}

export default LoginForm

