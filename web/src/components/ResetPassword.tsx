import React from 'react'
import { StoreContainer } from '../store'
import {Form, Button, Container, Grid} from 'semantic-ui-react'
import { NavLink } from 'react-router-dom'
import {Redirect} from 'react-router-dom'

interface Props {}

const ResetPassword: React.FC<Props> = () => {
    const store = StoreContainer.useContainer()
    if (!store.isLogin){
        return <Redirect to= '/LoginForm' />
    }
    return (  
    <div className="reset-pw">
      
        <Grid verticalAlign="middle" > 
            <Grid.Column style={{width:500}}>  
                <Form>
                     <h2>Reset Password</h2>
                     <div className={"login-submit-error" + (store.isError && store.errorMsg ? 'login-input-error' : '')}>            
                    {store.isError && store.errorMsg &&
                    <p className="login-error">{store.errorMsg}</p>}      
                            </div>
                    <Form.Field>
                        <div className={'field' + (store.isSubmit && !store.enteredCurrentPw ? 'login-input-error' : '')}>
                             <label>Current password:</label>
                            <input placeholder="current password" type="password" value={store.enteredCurrentPw} onChange={store.handleEnteredCurrentPw}/>
                            {store.isSubmit && !store.enteredCurrentPw && 
                                <div className="login-error"> Please enter your current password</div>}                        
                        </div>
                    </Form.Field>
                    <Form.Field>
                        <div className={'field' + (store.isSubmit && !store.enteredNewPw ? 'login-input-error' : '')}>
                            <label>New password: </label>
                            <input placeholder="new password" type="password" value={store.enteredNewPw} onChange={store.handleEnteredNewPw}/>
                            {store.isSubmit && !store.enteredNewPw && 
                                <div className="login-error"> Please enter your new password</div>}                        
                        </div>
                    </Form.Field>
                    <Button type="submit" className="ui primary button" onClick={store.handleResetPassword}>Save changes</Button>
                    <Button type="cancel"><NavLink to="/"></NavLink>Cancel</Button>
                 </Form>
             </Grid.Column>
        </Grid>
    </div> 
)
}
export default ResetPassword