import React from 'react'
import { Menu} from 'semantic-ui-react'

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