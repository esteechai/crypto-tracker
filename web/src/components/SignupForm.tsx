import React from 'react'
import { StoreContainer } from '../store'
import {Container, Grid, Segment} from 'semantic-ui-react'
import {Link, Redirect} from 'react-router-dom'

interface Props {}

const SignupForm : React.FC<Props> = () => {
    const store = StoreContainer.useContainer()

    if (!store.isLogin && store.signupVerification){
        return <Redirect to= '/signup-success' />
    }
    
    return (
        <div className="login-signup-container">
        <Container className="ui grid login-signup-container" style={{backgroundColor: ''}} textAlign="center">
        <Grid verticalAlign="middle">                       
          <Grid.Column style={{width:450}}>                      
             <form className="ui form">
                 <Segment stacked>
                     <h2> Signup</h2>
                           {(store.isError && store.errorMsg && store.isSubmit) ?
                      <div className="login-error"><i className="times circle outline icon"></i>{store.errorMsg}</div> : ''} 
                         <div className="login-signup-form">
                             <Container textAlign="left">
                                 <div className={'field required' + (!store.enteredUsername ? 'login-input-error' : '')}>
                                 <label><b>Username</b></label>
                                     <input placeholder="Email Address" id="email" value={store.enteredUsername} onChange={store.handleEnteredUsername} />
                                     {store.isSubmit && !store.enteredUsername && 
                                         <div className="login-error"> Required </div>}
                                 </div>
                             </Container>
                         </div>

                         <div className="login-signup-form">
                             <Container textAlign="left">
                                 <div className={'field required' + (!store.enteredEmail ? 'login-input-error' : '')}>
                                 <label><b>Email Address</b></label>
                                     <input placeholder="Email Address" id="email" value={store.enteredEmail} onChange={store.handleEnteredEmail} />
                                     {store.isSubmit && !store.enteredEmail && 
                                         <div className="login-error"> Required </div>}
                                 </div>
                             </Container>
                         </div>

                         <div className="login-signup-form">
                             <Container textAlign="left">
                                 <div className={'field required' + (store.isSubmit && !store.enteredPassword ? 'login-input-error' : '')}>
                                     <Container textAlign="left"><b>Password</b></Container>
                                 <input placeholder="Password" id="password" type="password" value={store.enteredPassword} onChange={store.handleEnteredPassword}/>
                                 {store.isSubmit && !store.enteredPassword && 
                                     <div className="login-error"> Required </div>}
                                 </div>
                             </Container>
                         </div>
                         <button type="submit" className="ui primary button" onClick={store.handleSignup}>Signup</button>
                         <p>Got an account? <Link to="/login" onClick={store.ResetFormInput}>Login</Link> </p>
                 </Segment>
             </form>
        </Grid.Column>
        </Grid>
    </Container>
    </div> 
    )
}
export {SignupForm}