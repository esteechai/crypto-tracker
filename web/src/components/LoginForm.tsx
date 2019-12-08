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
       return <Redirect to= '/CryptoList' />
    }

    store.useFetchProducts("http://localhost:8080/api/get/products")
    
return(
   <Container className="ui grid" style={{backgroundColor: ''}} textAlign="center">
           <Grid verticalAlign="middle">                       
             <Grid.Column style={{width:450}}>                      
                <form className="ui form">
                    <Segment stacked>
                        <h2> Login</h2>
                            <div className={"login-submit-error" + (store.isError && store.errorMsg ? 'has-error' : '')}>            
                                {store.isError && store.errorMsg &&
                                    <p className="login-error">{store.errorMsg}</p>}      
                            </div>

                          
                            <div className={'field' + (store.isSubmit && !store.enteredEmail ? 'has-error' : '')}>
                            <label><b>Email Address</b></label>
                                <input placeholder="Email Address" id="email" value={store.enteredEmail} onChange={store.handleEnteredEmail} />
                                {store.isSubmit && !store.enteredEmail && 
                                    <div className="login-error"> Email is required </div>}
                            </div>
                
                            <div className={'field' + (store.isSubmit && !store.enteredPassword ? 'has-error' : '')}>
                                <label><b>Password</b></label>
                            <input placeholder="Password" id="password" type="password" value={store.enteredPassword} onChange={store.handleEnteredPassword}/>
                            {store.isSubmit && !store.enteredPassword && 
                                <div className="login-error"> Password is required </div>}
                            </div>
                
                            <button type="submit" className="ui primary button" onClick={store.handleLogin}>Log in</button>
                            <p><Link to = "/">Forgot password?</Link></p>
                            <p>Don't have an account yet?   <Link to="/SignupForm">Signup</Link> </p>
                            
                    </Segment>
                </form>
         </Grid.Column>
         </Grid>
    </Container> 
)
}

export default LoginForm

