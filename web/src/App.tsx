import React from 'react';
import './App.css';
import {Login} from "./components/LoginForm"
import {Dashboard} from "./components/Dashboard"
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import {NavBar} from './components/NavBar';
import {LoginNavBar} from './components/LoginNavBar'; 
import {CryptoList} from './components/CryptoList';
import {FavCrypto} from './components/FavCrypto';
import {SignupForm} from './components/SignupForm';
import { StoreContainer } from './store';
import {ResetPassword} from './components/ResetPassword';
import {ForgotPassword} from './components/ForgotPassword';
import {ForgotPasswordVerified} from './components/ForgotPassVerified';
import {SignupSuccess} from './components/SignupSuccess';

const App: React.FC = () => {
  const store = StoreContainer.useContainer()

  return (
   <div>
      <Router> 
      {(store.isLogin)? <NavBar /> : <LoginNavBar />}
          <Switch>
            <Route exact path="/" component={Login} />
            <Route path="/dashboard" component={Dashboard} />
            <Route path="/navBar" component={NavBar} />
            <Route path="/crypto-list" component={CryptoList} />
            <Route path="/fav" component={FavCrypto} /> 
            <Route path="/signup" component={SignupForm} />   
            <Route path="/login" component={Login} />  
            <Route path="/reset-password" component={ResetPassword} />  
            <Route path="/forgot-password" component={ForgotPassword} />  
            <Route path="/forgot-pass-verified" component={ForgotPasswordVerified} /> 
            <Route path="/signup-success" componenet={SignupSuccess} />
          </Switch>
      </Router>   
    </div>
  );
}

export default App;



