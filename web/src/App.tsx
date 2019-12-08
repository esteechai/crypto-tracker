import React from 'react';
import './App.css';
import LoginForm from "./components/LoginForm"
import Dashboard from "./components/Dashboard"
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import NavBar from './components/NavBar';
import CryptoList from './components/CryptoList';
import ProductDetails from './components/ProductDetails';
import FavCrypto from './components/FavCrypto';
import SignupForm from './components/SignupForm';

const App: React.FC = () => {
  return (
    <Router>   
        <NavBar />
        <Switch>
          <Route exact path="/" component={LoginForm} />
          <Route path="/Dashboard" component={Dashboard} />
          <Route path="/NavBar" component={NavBar} />
          <Route path="/CryptoList" component={CryptoList} />
          <Route path="/ProductDetails" component={ProductDetails} />   
          <Route path="/FavCrypto" component={FavCrypto} /> 
          <Route path="/SignupForm" component={SignupForm} />      
        </Switch>
    </Router>   
  );
}

export default App;

