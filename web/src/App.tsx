import React from 'react';
import './App.css';
import LoginForm from "./components/LoginForm"
import Dashboard from "./components/Dashboard"
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import NavBar from './components/NavBar';
import LoginNavBar from './components/LoginNavBar'; 
import CryptoList from './components/CryptoList';
import FavCrypto from './components/FavCrypto';
import SignupForm from './components/SignupForm';
import { StoreContainer } from './store';
import ResetPassword from './components/ResetPassword';
import ForgotPassword from './components/ForgotPassword';

const App: React.FC = () => {
  const store = StoreContainer.useContainer()

  return (
   <div>
      <Router> 
      {(store.isLogin)? <NavBar /> : <LoginNavBar />}
          <Switch>
            <Route exact path="/" component={LoginForm} />
            <Route path="/Dashboard" component={Dashboard} />
            <Route path="/NavBar" component={NavBar} />
            <Route path="/CryptoList" component={CryptoList} />
            <Route path="/FavCrypto" component={FavCrypto} /> 
            <Route path="/SignupForm" component={SignupForm} />   
            <Route path="/LoginForm" component={LoginForm} />  
            <Route path="/ResetPassword" component={ResetPassword} />  
            <Route path="/ForgotPassword" component={ForgotPassword} />  
          </Switch>
      </Router>   
    </div>
  );
}

export default App;



