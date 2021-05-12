import React,{  useState,useEffect }   from 'react';
import NavbarAdmin from './NavbarAdmin';
import Footer from './Footer';
import { useHistory } from 'react-router-dom';
import {terminar_temporada,mensaje} from "../api/api-admin";
import Swal from "sweetalert2";




function TemporadaAdmin() {
    
    const [total, setTotal] = useState([]);  
    //const [selectedFiles, setSelectedFiles] = useState([]);
    let history = useHistory();
    var a = localStorage.getItem("usuarioActual")
    const Toast = Swal.mixin({
        toast: true,
        position: "top-end",
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        didOpen: (toast) => {
          toast.addEventListener("mouseenter", Swal.stopTimer);
          toast.addEventListener("mouseleave", Swal.resumeTimer);
        },
      });
    
    
    console.log(a)
    if (a === null){
        history.push('/home')
        
    }
    var obj = JSON.parse(a)
    
    if(obj.id_usuario !==1){
        history.push('/usuario')
        
    }
    useEffect(() => {
        
            fetch("http://localhost:3080/Apuntaje", { method : 'POST', 
                headers: {
                    Accept: "application/json",
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    mensaje: "mensaje"
                }) })
                .then((response) => response.json())
                .then(data => {
                    if (data === null ){

                    }else{
                        setTotal(data) 
                    }
                    
                    });


            
        
    },[])
    async function handleDelete(e) {
        
        console.log(e.target.id)
        
        history.push('/UserPred/'+e.target.id)
        
    }

    async function handleEND(e) {
        
        var rawResponse = await terminar_temporada();
        if(rawResponse.status === 201){
            console.log("realizo el cambio")
        }

        rawResponse = await mensaje();
        if(rawResponse.status === 201){
            const respuesta = await rawResponse.json();
          
            Toast.fire({
              icon: 'success',
              title: `¡${respuesta.mensaje}!`
              
            })
            
           
          } else{
            Toast.fire({
              icon: 'error',
              title: `No se pudo registrar`
            })
          }
        
    }


    return (
        <div>
            <NavbarAdmin />
            
            <div>

                <br></br>
                <center>
                <button type="button" onClick = {handleEND} className="btn btn-warning">Terminar Temporada</button><small>&nbsp;&nbsp;&nbsp;&nbsp;Termina y crea automaticamente las temporadas</small>
                
                </center>
                <br></br>
                <br></br>
                <br></br>
                
            <table className = "table table-sm">
                    <thead>
                        <tr>
                            <th>Posicion</th>   
                            <th>Username</th>
                            <th>Nombre</th>
                            <th>Apellido</th>
                            
                            <th>P10</th>
                            <th>P5</th>
                            <th>P3</th>
                            <th>P0</th>
                            <th>Puntos</th>
                        </tr>
                    </thead>
                    <tbody>
                    {total.map((option) => (
                        <tr > 
                            
                            {/*<th ><button  id={option.id}  onClick ={handleDelete}>{option.id}</button></th>*/}
                            <th>{option.puesto}</th>
                            <th><button  id={option.id} onClick={handleDelete} >{option.username}</button></th>
                            <th>{option.nombre}</th>
                            <th>{option.apellido}</th>
                            <th>{option.p5}</th>
                            <th>{option.p5}</th>
                            <th>{option.p3}</th>
                            <th>{option.p0}</th>
                            <th>{option.puntos}</th>
                            
                        </tr>
                          ))}
                    </tbody>

                </table>
                </div>
            <Footer />
        </div>
    )
}



export default TemporadaAdmin
