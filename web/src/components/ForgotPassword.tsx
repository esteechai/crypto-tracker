import React, { useReducer } from 'react'
import { StoreContainer } from '../store'
import {Form, Button, Grid} from 'semantic-ui-react'
import { NavLink, Redirect } from 'react-router-dom'

interface Props {}

const ForgotPassword: React.FC<Props> = () => {
    const store = StoreContainer.useContainer()

    if(store.verifiedEmail){
        return <Redirect to ="/forgot-pass-verified" /> 
    }

    return (  
    <div className="reset-pw">
        <Grid verticalAlign="middle" > 
            <Grid.Column style={{width:500}}>  
                <Form>
                     <h2>Forgot Password</h2>
                     <p>Forgotten your password? Enter your email address below, and we'll email you instructions for setting a new one.</p>
                    {(store.isError && store.errorMsg && store.isSubmit) ?
                    <div className="login-error"><i className="times circle outline icon"></i>{store.errorMsg}</div> : 
                    ''}                          
                    <Form.Field>
                        <div>
                             <label><b>Email address:</b></label>
                            <input placeholder="current password" value={store.enteredEmail} onChange={store.handleEnteredEmail}/>
                            {store.isSubmit && !store.enteredEmail && 
                            <div className="login-error"> Required </div>}                        
                        </div>
                    </Form.Field>
                    <Button type="submit" className="ui primary button" onClick={store.HandleForgotPassword}>Reset my password</Button>
                    <Button type="cancel"><NavLink to="/login" onClick={store.ResetForgotPassInput}>Cancel</NavLink></Button>
                 </Form>
             </Grid.Column>
        </Grid>
    </div> 
)
}
export {ForgotPassword}