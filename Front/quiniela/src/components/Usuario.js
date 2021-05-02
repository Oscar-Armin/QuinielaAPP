import React from 'react'
import Footer from './Footer'
import NavbarUser from './NavbarUser'
import { useHistory } from 'react-router-dom';
import { consultUser } from "../api/api-user";
import Swal from "sweetalert2";

export default function Usuario() {
    let history = useHistory();
    var a = localStorage.getItem("usuarioActual")
    //console.log(a)
    if (a === null){
        history.push('/home')
        return(<></>);
    }
    var obj = JSON.parse(a)
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
    async function evento() {
        
    
        
    
        try {
          const rawResponse = await consultUser(
            obj.id_usuario
          );
    
          
          
          //console.log(respuesta.ID)
          
          if (rawResponse.status === 201) {
            
            let respuesta = await rawResponse.json();
            localStorage.setItem(
                "editar_user",
                JSON.stringify({
                    id_usuario: respuesta.id_usuario,
                    nombre:respuesta.nombre,
                    usuario: respuesta.username,
                    apellido:respuesta.apellido,
                    nacimiento:respuesta.nacimiento,
                    registro:respuesta.registro,
                    correo:respuesta.correo,
                    foto:respuesta.foto,
                    password:respuesta.password
                
                })
              );
            
            
            
 
            
          } else if(rawResponse.status === 500){
            Toast.fire({
              icon: 'error',
              title: 'Algo ocurrio mal'
            })
          } else {
            Toast.fire({
              icon: "error",
              title: "Algo ocurrio mal",
            });
            history.push('/registrar')
          }
        } catch (error) {
          console.log(error);
          Toast.fire({
            icon: "error",
            title: "No se pudo iniciar sesión",
          });
          history.push('/registrar')
        }
        
        
    }
    
    evento()

    return (
        <div>
            <NavbarUser />
            pagina principal user
            <Footer />
        </div>
    )
}
