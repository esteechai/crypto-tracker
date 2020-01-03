import React from 'react'
import { StoreContainer } from '../store'
import {Form, Button, Grid, Container} from 'semantic-ui-react'
import { NavLink } from 'react-router-dom'
import {Redirect} from 'react-router-dom'

interface Props {}

const ResetPassSuccess: React.FC<Props> = () => {
    const store = StoreContainer.useContainer()

    if (!store.isLogin){
        return <Redirect to= '/login' />
    }

    return ( 
        <div className="login-signup-container">
            <Container className="ui grid login-signup-container" textAlign="center">
            <Grid verticalAlign="middle" > 
                 <Grid.Column style={{width:500}}>  
                    <Form>
                        <Form.Field>
                        <h2>Password Reset Successful</h2>
                                <p className="success-msg"><i className="check circle icon"></i>Your password has been successfully changed.</p> 
                            <Button type="cancel"><NavLink to="/" onClick={store.ResetResetPwInput}>OK</NavLink></Button>
                        </Form.Field>
                    </Form>
                </Grid.Column>
            </Grid>
            </Container>
        </div> 
    )
}
export {ResetPassSuccess}