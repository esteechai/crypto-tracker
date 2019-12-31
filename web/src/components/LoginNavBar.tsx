import React from 'react'
import { Menu, Input, Search, Dropdown, Icon } from 'semantic-ui-react'
import { NavLink } from 'react-router-dom'
interface Props {}

const LoginNavBar: React.FC<Props> = () => {
    return (
      <div>
        <Menu attached="top" tabular className="login-navbar"> 
            <Menu.Item id="logo" name="Crypto Tracker" />
        </Menu>
      </div>
    )
}

export {LoginNavBar}