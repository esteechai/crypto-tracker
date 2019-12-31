import React, {useEffect} from 'react'
import { StoreContainer } from '../store'
import {Container, Grid, Segment} from 'semantic-ui-react'
import {Link, Redirect} from 'react-router-dom'

interface Props{}

const LoginForm: React.FC<Props> = () => {
    const store = StoreContainer.useContainer()
    if(store.isLogin){
        store.fetchDataFromAPI("/api/get/products","product")
       return <Redirect to= '/CryptoList' />
    }


// useEffect(() => {
   
// },[])

    
return(
    <div className="login-signup-container">
   <Container className="ui grid" style={{backgroundColor: ''}} textAlign="center">
           <Grid verticalAlign="middle">                       
             <Grid.Column style={{width:450}}>                      
                <form className="ui form">
                    <Segment stacked>
                        <h2>Login</h2>
                            <div className={"login-submit-error" + (store.isError && store.errorMsg ? 'login-input-error' : '')}>            
                                {store.isError && store.errorMsg &&
                                    <p className="login-error">{store.errorMsg}</p>}      
                            </div>
                            <div className="login-signup-form">
                                <Container textAlign="left">
                                    <div className={'field' + (store.isSubmit && !store.enteredEmail ? 'login-input-error' : '')}>
                                    <label><b>Email Address</b></label>
                                        <input placeholder="Email Address" id="email" value={store.enteredEmail} onChange={store.handleEnteredEmail} />
                                        {store.isSubmit && !store.enteredEmail && 
                                            <div className="login-error">Required </div>}
                                    </div>
                                </Container>
                            </div>

                            <div className="login-signup-form">
                                <Container textAlign="left">
                                    <div className={'field' + (store.isSubmit && !store.enteredPassword ? 'login-input-error' : '')}>
                                        <Container textAlign="left"><b>Password</b></Container>
                                    <input placeholder="Password" id="password" type="password" value={store.enteredPassword} onChange={store.handleEnteredPassword}/>
                                    {store.isSubmit && !store.enteredPassword && 
                                        <div className="login-error"> Required </div>}
                                    </div>
                                </Container>
                            </div>

                            <button type="submit" className="ui primary button" onClick={store.handleLogin}>Log in</button>
                            <p><Link to = "/ForgotPassword" onClick={store.ResetFormInput}>Forgot password?</Link></p>
                            <p>Don't have an account yet?<Link to="/SignupForm" onClick={store.ResetFormInput}>Signup</Link></p> 
                    </Segment>
                </form>
         </Grid.Column>
         </Grid>
    </Container> 
    </div>
)
}

export default LoginForm

