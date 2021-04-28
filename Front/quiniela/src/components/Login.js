import React, { useRef, useState } from "react";
import Footer from './Footer'
import Navbar from './Navbar'
import { Link } from "react-router-dom";
import Swal from "sweetalert2";
import { useHistory } from 'react-router-dom';
import { loginUsuario } from "../api/api-user";


const Login = ({titulo})=>{

    let history = useHistory();
    const redirect = () => {

        history.push('/registrar')
      }
    
    const usuarioRef = useRef();
    const passwordRef = useRef();
  
    const [loading, setLoading] = useState(false);
    console.log(loading)

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
    
        setLoading(true);
    
        try {
          const rawResponse = await loginUsuario(
            usuarioRef.current.value,
            passwordRef.current.value
          );
    
          
          
          //console.log(respuesta.ID)
          
          if (rawResponse.status === 201) {
            const respuesta = await rawResponse.json();
                localStorage.setItem(
                    "usuarioActual",
                    JSON.stringify({
                      usuario: usuarioRef.current.value,
                      
                      id_usuario: respuesta.ID
                    })
                  );
                    
                  Toast.fire({
                    icon: "success",
                    title: `¡Bienvenid@ ${usuarioRef.current.value}!`,
                  });
            


            
          } else if(rawResponse.status === 500){
            Toast.fire({
              icon: 'error',
              title: 'Credenciales inválidas'
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
    
        setLoading(false);
      }
    
    
    return(
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
                    Iniciar sesión
                  </h6>
                </div>

                <hr className="mt-6 border-b-1 border-gray-400" />
              </div>
              <div className="flex-auto px-4 lg:px-10 py-10 pt-0">
                <div className="text-gray-500 text-center mb-3 font-bold">
                  <small>Ingresa tus datos</small>
                </div>
                <form onSubmit={handleSubmit}>
                  <div className="relative w-full mb-3">
                    <label
                      className="block uppercase text-gray-700 text-xs font-bold mb-2"
                      htmlFor="grid-password"
                    >
                      Usuario:
                    </label>
                    <br/>
                    <input
                      type="text"
                      ref={usuarioRef}
                      className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                      placeholder="Usuario"
                      required
                    />
                  </div>

                  <div className="relative w-full mb-3">
                    <label
                      className="block uppercase text-gray-700 text-xs font-bold mb-2"
                      htmlFor="grid-password"
                    >
                      Contraseña:
                    </label>
                    <br/>
                    <input
                      type="password"
                      ref={passwordRef}
                      className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                      placeholder="Password"
                      required
                    />
                  </div>
                  <div align="center"> 
                    <button
                      className="bg-gray-900 text-black active:bg-gray-700 text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 w-full ease-linear transition-all duration-150"
                      type="submit"
                    >
                      Iniciar sesión
                    </button>
       
                  </div>
                </form>
                <center>
                <button 
                      className="bg-gray-900 text-black active:bg-gray-700 text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 w-full ease-linear transition-all duration-150"
                      onClick={redirect}
                    >
                      Registrar
                    </button>
                    </center>
              </div>
            </div>
            <div className="flex flex-wrap mt-6 relative">
              <div className="w-1/2">
                <Link to="/Forget" className="text-gray-300">
                  <small>Olvidaste la contraseña</small>
                </Link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
            < Footer />
            
        </div>
        
    );
}

export default Login;