import React from 'react'
import { StoreContainer } from '../store'
import { Menu, Input, Search, Dropdown, Confirm } from 'semantic-ui-react'
import { NavLink, Link } from 'react-router-dom'
interface Props {}

const NavBar: React.FC<Props> = () => {
    const store = StoreContainer.useContainer()
    return (
        <div className="navbar"> 
            <Menu attached="top" tabular> 
                <Menu.Item 
                    name="Cyrpto Tracker"
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
                <Menu.Item name="Cryptocurrency">
                    <NavLink to ="/CryptoList">Cryptocurrency</NavLink>
                </Menu.Item>

                <Menu.Item>
                    <NavLink to ="/FavCrypto">Favourites</NavLink>
                    </Menu.Item>
                
            <Dropdown item icon="user"simple>
                <Dropdown.Menu>
                    <Dropdown.Item onClick={store.handleLogoutMsg}>Logout</Dropdown.Item>
                    <Confirm 
                        open={store.openLogoutMsg}
                        content ="Are you sure you want to logout?"
                        onCancel={store.CancelLogout}
                        onConfirm={store.ConfirmLogout}
                    />
                    <Dropdown.Divider />
                    <Dropdown.Item><NavLink to = "/ResetPassword" onClick={store.ResetResetPwInput}>Reset Password</NavLink></Dropdown.Item>
                </Dropdown.Menu>
            </Dropdown>
            </Menu.Menu> 
            </Menu>
         </div>
    )
}

export default NavBar