import React from 'react';
import './App.css';
import LoginForm from "./components/LoginForm"
import Dashboard from "./components/Dashboard"
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import NavBar from './components/NavBar';
import CryptoList from './components/CryptoList';

const App: React.FC = () => {
  return (
    // <div>
    //   <LoginForm/>
    // </div>
    <Router>   
       <Switch>
        <Route exact path="/" component={LoginForm} />
        <Route path="/Dashboard" component={Dashboard} />
        <Route path="/NavBar" component={NavBar} />
        <Route path="/CryptoList" component={CryptoList} />
      </Switch>
    </Router>   
  );
}

export default App;

