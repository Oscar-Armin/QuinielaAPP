import React from 'react'

import { useHistory } from 'react-router-dom';


function Navbar_user() {
    let history = useHistory();
    const redirect = () => {

        history.push('/home')
      }
    

    return (
        <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <div className="container-fluid">
        <a className="navbar-brand" href="/usuario">
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
                        <a className="nav-link active" aria-current="page" href="/home">Perfil</a>
                    </li>
                </ul>
                <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                    <li className="nav-item">
                        <a className="nav-link active" aria-current="page" href="/home">Membresia</a>
                    </li>
                </ul>
                <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                    <li className="nav-item">
                        <a className="nav-link active" aria-current="page" href="/home">Eventos </a>
                    </li>
                </ul>
                <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                    <li className="nav-item">
                        <a className="nav-link active" aria-current="page" href="/home">Recompensas </a>
                    </li>
                </ul>


            </div>
            
                    
                    <form className="d-flex" align="right" id="flex-end">
                    
                        <button class="btn btn-danger" type="submit" align="right"  onClick={redirect}>Log out </button>
                    </form>
                    
        </div>
        </nav>
    );
}
export default Navbar_user;