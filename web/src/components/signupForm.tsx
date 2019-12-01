import React from 'react'
import { StoreContainer } from '../store'

interface Props{}

const SignupForm : React.FC<Props> = () => {
    const store = StoreContainer.useContainer()
    return (
        <div>
            <p>sign up </p>
        </div>
    )

}

export default SignupForm