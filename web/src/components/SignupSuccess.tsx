import React from 'react'
import { StoreContainer } from '../store'
import {Form, Button, Grid, Container} from 'semantic-ui-react'
import { NavLink } from 'react-router-dom'

interface Props {}

const SignupSuccess: React.FC<Props> = () => {
    console.log("sign up sucess.tsx")
    const store = StoreContainer.useContainer()
    return (  
    <div className="login-signup-container">
        <Container className="ui grid login-signup-container" style={{backgroundColor: ''}} textAlign="center">
        <Grid verticalAlign="middle" > 
            <Grid.Column style={{width:500}}>  
                <Form>
                    <Form.Field>
                     <h2>Email Verification</h2>
                             <p className="success-msg"><i className="check circle icon"></i>A verification link has been sent to your email. Please check your email and confirm your email address.</p> 
                         <Button type="cancel"><NavLink to="/login" onClick={store.ResetFormInput}>OK</NavLink></Button>
                    </Form.Field>
                 </Form>
             </Grid.Column>
        </Grid>
        </Container>
    </div> 
    
)
}
export {SignupSuccess}