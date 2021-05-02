import React from 'react'
import Footer from './Footer'
import NavbarAdmin from './NavbarAdmin'
import { useHistory } from 'react-router-dom';

const Administrador = ()=>{
    let history = useHistory();
    var a = localStorage.getItem("usuarioActual")
    console.log(a)
    if (a === null){
        history.push('/home')
        return(<></>);
    }
    var obj = JSON.parse(a)
    
    if(obj.id_usuario !==1){
        history.push('/home')
        return(<></>);
    }
    
    return (
        <div>
            <NavbarAdmin />
            Pagina admin principal
            <Footer />
            
        </div>
    )
}
export default Administrador;