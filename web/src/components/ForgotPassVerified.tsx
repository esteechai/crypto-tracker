import React from 'react'
import { StoreContainer } from '../store'
import {Form, Button, Grid} from 'semantic-ui-react'
import { NavLink } from 'react-router-dom'

interface Props {}

const ForgotPasswordVerified: React.FC<Props> = () => {

    const store = StoreContainer.useContainer()
    return (  
    <div className="reset-pw">
        <Grid verticalAlign="middle" > 
            <Grid.Column style={{width:500}}>  
                <Form>
                    <Form.Field>
                     <h2>Forgot Password</h2>
                             <p>A confirmation email has been sent to <b>{store.enteredEmail}</b>.</p> 
                             <p>Please check your inbox to complete reset password process.</p>
                         <Button type="cancel"><NavLink to="/LoginForm" onClick={store.ResetFormInput}>OK</NavLink></Button>
                    </Form.Field>
                 </Form>
             </Grid.Column>
        </Grid>
    </div> 
)
}
export default ForgotPasswordVerified