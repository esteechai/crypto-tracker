import React from 'react'
import { StoreContainer } from '../store'
import { Menu, Input, Search, Dropdown, Icon } from 'semantic-ui-react'
interface Props {}

const NavBar: React.FC<Props> = () => {
    const store = StoreContainer.useContainer()
    return (
        <div> 
            <Menu attached="top" tabular> 
                <Menu.Item 
                    name="Cypto Tracker"
                />

                <Menu.Item>
                    <Input
                        transparent
                        icon={{name:"search", link:true}}
                        placeholder="Search"
                    />
                </Menu.Item>
                <Menu.Menu position="right">
                <Menu.Item
                    name="Cryptocurrency"
                    //active={activeItem === 'bio'}
                    //onClick={this.handleItemClick}
                />

                <Menu.Item
                    name="Favourites"
                    //onClick="{}
                />
           
                  
            <Dropdown item icon="user"simple>
                <Dropdown.Menu>
                    <Dropdown.Item>Logout</Dropdown.Item>
                    <Dropdown.Divider />
                    <Dropdown.Item>Reset Password</Dropdown.Item>
                </Dropdown.Menu>
            </Dropdown>
            </Menu.Menu> 
            </Menu>
        </div>
    )

}

    export default NavBar