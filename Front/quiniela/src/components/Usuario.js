import React from 'react'
import Footer from './Footer'
import NavbarUser from './NavbarUser'
import { useHistory } from 'react-router-dom';

export default function Usuario() {
    let history = useHistory();
    var a = localStorage.getItem("usuarioActual")
    //console.log(a)
    if (a === null){
        history.push('/home')
        return(<></>);
    }
    
    

    return (
        <div>
            <NavbarUser />
            pagina principal user
            <Footer />
        </div>
    )
}
