import React from 'react'

import { useHistory } from 'react-router-dom';


function Navbar() {
    let history = useHistory();
    const redirect = () => {

        history.push('/login')
      }
    

    return (
        <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <div className="container-fluid">
        <a class="navbar-brand" href="/home">
      <img src="./logo.png" alt="" width="90" height="30" />
    </a>

            <div className="collapse navbar-collapse" id="navbarSupportedContent">
                {/*<ul className="navbar-nav me-auto mb-2 mb-lg-0">
                    <li className="nav-item">
                        <a className="nav-link active" aria-current="page" href="#">Home</a>
                    </li>
                </ul>*/}
                <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                    <li className="nav-item">
                        <a className="nav-link active" aria-current="page" href="/home">Equipos</a>
                    </li>
                </ul>
                <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                    <li className="nav-item">
                        <a className="nav-link active" aria-current="page" href="/home">Deportes</a>
                    </li>
                </ul>
                <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                    <li className="nav-item">
                        <a className="nav-link active" aria-current="page" href="/home">Acerca de nosotros</a>
                    </li>
                </ul>

            </div>
            <right>
                    
                    <form className="d-flex" align="right" id="flex-end">
                    
                        <button className="btn btn-outline-success" type="submit" align="right"  onClick={redirect}>Login</button>
                    </form>
                    </right>
        </div>
        </nav>
    );
}
export default Navbar;