import React from 'react'
import { StoreContainer } from '../store'
import { Redirect } from 'react-router'

interface DashboardProps {
}

const Dashboard: React.FC<DashboardProps> = () => {
    const store = StoreContainer.useContainer() 
    if (!store.isLogin){
        return <Redirect to= '/login' />
    }
    return (
        <div></div>
    )
}

export {Dashboard}


