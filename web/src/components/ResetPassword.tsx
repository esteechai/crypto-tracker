import React from 'react'
import { StoreContainer } from '../store'
import {Form, Button, Grid} from 'semantic-ui-react'
import { NavLink } from 'react-router-dom'
import {Redirect} from 'react-router-dom'

interface Props {}

const ResetPassword: React.FC<Props> = () => {
    const store = StoreContainer.useContainer()
    if (!store.isLogin){
        return <Redirect to= '/login' />
    }

    if (store.passReset && store.isLogin){
        return <Redirect to = '/reset-pass-success'/>
    }

    return (  
    <div className="reset-pw">
        <Grid verticalAlign="middle" > 
            <Grid.Column style={{width:500}}>  
                <Form>
                     <h2>Reset Password</h2>
                    {(store.isError && store.errorMsg && store.isSubmit) ?
                    <div className="login-error"><i className="times circle outline icon"></i>{store.errorMsg}</div> : '' }
                    <Form.Field>
                        <div>
                             <label><b>Current password:</b></label>
                            <input placeholder="current password" type="password" value={store.enteredCurrentPw} onChange={store.handleEnteredCurrentPw}/>
                            {store.isSubmit && !store.enteredCurrentPw && 
                                <div className="login-error"> Please enter your current password</div>}                        
                        </div>
                    </Form.Field>
                    <Form.Field>
                         <div>
                            <label><b>New password: </b></label>
                            <input placeholder="new password" type="password" value={store.enteredNewPw} onChange={store.handleEnteredNewPw}/>
                            {store.isSubmit && !store.enteredNewPw && 
                                <div className="login-error"> Please enter your new password</div>}                        
                        </div>
                    </Form.Field>
                    <Button type="submit" className="ui primary button" onClick={store.handleResetPassword}>Save changes</Button>
                    <Button type="cancel"><NavLink to="/crypto-list" onClick={store.ResetResetPwInput}>Cancel</NavLink></Button>
                 </Form>
             </Grid.Column>
        </Grid>
    </div> 
)
}
export {ResetPassword}