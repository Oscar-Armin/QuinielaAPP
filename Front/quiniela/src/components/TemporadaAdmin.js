import {React}  from 'react';
import NavbarAdmin from './NavbarAdmin';
import Footer from './Footer';
import { useHistory } from 'react-router-dom';


function TemporadaAdmin() {

    //const [selectedFiles, setSelectedFiles] = useState([]);
    let history = useHistory();
    var a = localStorage.getItem("usuarioActual")
    
    console.log(a)
    if (a === null){
        history.push('/home')
        
    }
    var obj = JSON.parse(a)
    
    if(obj.id_usuario !==1){
        history.push('/home')
        
    }
    return (
        <div>
            <NavbarAdmin />
           
            <Footer />
        </div>
    )
}



export default TemporadaAdmin
