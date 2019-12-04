import React from 'react'
import { StoreContainer } from '../store'
import { Menu, Input, Search, Dropdown, Icon } from 'semantic-ui-react'
import { NavLink } from 'react-router-dom'
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
                    <div className="ui icon input">
                        <Input type="text "
                            value={store.searchKey}
                            placeholder="Search"
                            onChange={store.handleSearch}
                        />
                    <i aria-hidden="true" className="search icon"></i>
                    </div>          
                </Menu.Item>
                <Menu.Menu position="right">
                <Menu.Item>
                    <NavLink to ="/CryptoList">Cryptocurrency</NavLink>
                </Menu.Item>

                <Menu.Item>
                    <NavLink to ="/">Favourites</NavLink>
                    </Menu.Item>
                
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