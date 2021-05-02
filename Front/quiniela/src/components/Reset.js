import React, { useRef } from "react";
import Navbar from './Navbar'
import Footer from './Footer'
import Swal from "sweetalert2";
import {withRouter} from 'react-router-dom';
//import { useHistory } from 'react-router-dom';
import { resetPassword } from "../api/api-user";

function Reset(props) {
    const passwordRef = useRef();
    const confpasswordRef = useRef();
    let identificador = parseInt(props.match.params.id)
    

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

      async function handleSubmit(e) {
        e.preventDefault();
    
        //setLoading(true);
        if (passwordRef.current.value !== confpasswordRef.current.value) {
            Swal.fire({
              title: "Error",
              text: "Las contraseñas no coinciden",
              icon: "error",
              confirmButtonText: "Ok",
            });
            //window.location.reload(false);
            return;
          }
    
          
          
          
          if ( !password_validate(passwordRef.current.value)) {
            Swal.fire({
              title: "Error",
              text: "contraseña invalida*",
              icon: "error",
              confirmButtonText: "Ok",
            });
            //window.location.reload(false);
            
            return;
          } 
    
    
        try {
            
          const rawResponse = await resetPassword(
            identificador,
            passwordRef.current.value
            
          );
          
          
          
          //console.log(respuesta.ID)
          
          if (rawResponse.status === 201) {
            const respuesta = await rawResponse.json();
            console.log(respuesta)

                    //console.log(respuesta.ID)
                  Toast.fire({
                    icon: "success",
                    title: `¡Se actualizo la contraseña!`,
                  });

            
          } else if(rawResponse.status === 500){
            Toast.fire({
              icon: 'error',
              title: 'Usuario no existe'
            })
          } else {
            Toast.fire({
              icon: "error",
              title: "No se pudo iniciar sesión",
            });
          }
        } catch (error) {
          console.log(error);
          Toast.fire({
            icon: "error",
            title: "No se pudo iniciar sesión",
          });
        }
    
        //setLoading(false);
      }
    

    return (
        <div>

            <Navbar />
            <>
      <div className="container mx-auto px-4 h-full">
        <div className="flex content-center items-center justify-center h-full">
          <div className="w-full lg:w-4/12 px-4">
            <div className="relative flex flex-col min-w-0 break-words w-full mb-6 shadow-lg rounded-lg bg-gray-300 border-0">
              <div className="rounded-t mb-0 px-6 py-6">
                <div className="text-center mb-3">
                  <h6 className="text-gray-600 text-sm font-bold">
                    Reestrablecer contraseña
                  </h6>
                </div>

                <hr className="mt-6 border-b-1 border-gray-400" />
              </div>
              <div className="flex-auto px-4 lg:px-10 py-10 pt-0">
                <div className="text-gray-500 text-center mb-3 font-bold">
                  <small>Ingrese su nueva contraseña </small>
                </div>
                <form onSubmit={handleSubmit}>
      
                <div className="relative w-full mb-3">
                      <label
                        className="block uppercase text-gray-700 text-xs font-bold mb-2"
                        htmlFor="grid-password"
                      >
                        Contraseña
                      </label>
                      <p/>
                      <input
                        type="password"
                        ref={passwordRef}
                        className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                        placeholder="Contraseña"
                        required
                      />
                    </div>
  
                    <div className="relative w-full mb-3">
                      <label
                        className="block uppercase text-gray-700 text-xs font-bold mb-2"
                        htmlFor="grid-password"
                      >
                        Confirmación contraseña
                      </label>
                      <p/>
                      <input
                        type="password"
                        ref={confpasswordRef}
                        className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                        placeholder="Confirmar contraseña"
                        required
                      />
                    </div>

                  <div align="center"> 
                    <button
                      className="bg-gray-900 text-black active:bg-gray-700 text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 w-full ease-linear transition-all duration-150"
                      type="submit"
                    >
                      Enviar
                    </button>
       
                  </div>
                </form>

              </div>
            </div>

          </div>
        </div>
      </div>
    </>
            < Footer />
            
        </div>
    )
}

function password_validate(p) {
    return /[A-Z]/.test(p)&& /[a-z]/.test(p) && /[0-9]/.test(p) && /^[A-Za-z0-9]{8,20}$/.test(p) ;
  }
  
export default withRouter(Reset);
