import React from 'react'
import { Switch, Route, Redirect } from 'react-router-dom';
import Home from './Home'
import Login from './Login'
import Registrar from './Registrar';

const Main = ()=>{
    return(
        <Switch>

            <Route path="/home" component={ () => <Home  /> }  />
            <Route path="/login" component={ () => <Login  /> }  />
            <Route path="/registrar" component={ () => <Registrar  /> }  />
            
            <Redirect to="/home" />
        </Switch>
    );
}

export default Main;