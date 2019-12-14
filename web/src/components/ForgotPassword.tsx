import React from 'react'
import { StoreContainer } from '../store'
import {Form, Button, Grid} from 'semantic-ui-react'
import { NavLink } from 'react-router-dom'

interface Props {}

const ForgotPassword: React.FC<Props> = () => {
    const store = StoreContainer.useContainer()
    return (  
    <div className="reset-pw">
        <Grid verticalAlign="middle" > 
            <Grid.Column style={{width:500}}>  
                <Form>
                     <h2>Forgot Password</h2>
                     <p>A reset password link will be sent to your email.</p>
                    {(store.isError && store.errorMsg && store.isSubmit) ?
                    <div className="login-error"><i className="times circle outline icon"></i>{store.errorMsg}</div> : 
                    <div className="success-msg"><i className="check circle icon"></i>{store.successMsg}</div>}                          <Form.Field>
                        <div>
                             <label><b>Email address:</b></label>
                            <input placeholder="current password" value={store.enteredEmail} onChange={store.handleEnteredEmail}/>
                            {store.isSubmit && !store.enteredEmail && 
                            <div className="login-error"> Required </div>}                        
                        </div>
                    </Form.Field>
                    <Button type="submit" className="ui primary button" onClick={store.HandleForgotPassword}>Submit</Button>
                    <Button type="cancel"><NavLink to="/LoginForm" onClick={store.ResetForgotPassInput}>Cancel</NavLink></Button>
                 </Form>
             </Grid.Column>
        </Grid>
    </div> 
)
}
export default ForgotPassword