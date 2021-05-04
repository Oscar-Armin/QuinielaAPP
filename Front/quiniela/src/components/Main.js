
import React from 'react'
import { Switch, Route, Redirect } from 'react-router-dom';
import Home from './Home'
import LoadImage from './LoadImage';
import Login from './Login'
import Registrar from './Registrar';
import Usuario from './Usuario';
import Administrador from './Administrador';
import Perfil from './Perfil';
import Forget from './Forget';
import Reset from './Reset';
import Cargar from './Cargar';

const Main = ()=>{
    return(
        <Switch>

            <Route path="/home" component={ () => <Home  /> }  />
            <Route path="/login" component={ () => <Login  /> }  />
            <Route path="/registrar" component={ () => <Registrar  /> }  />
            <Route path="/load" component={ () => <LoadImage  /> }  />
            <Route path="/usuario" component={ () => <Usuario  /> }  />
            <Route path="/admin" component={ () => <Administrador  /> }  />
            <Route path="/perfil" component={ () => <Perfil  /> }  />
            <Route path="/forget" component={ () => <Forget  /> }  />
            <Route path="/reset/:id" component={ () => <Reset  /> }  />
            <Route path="/carga" component={ () => <Cargar  /> }  />
            
            <Redirect to="/home" />
        </Switch>
    );
}


export default Main;