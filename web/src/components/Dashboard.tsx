import React from 'react'
import { StoreContainer } from '../store'

interface Props{}

const Dashboard: React.FC<Props> = () => {
    const store = StoreContainer.useContainer()
    return (
        <div>
            <p>You're logged in</p>
            
        </div>
    )
}

export default Dashboard


