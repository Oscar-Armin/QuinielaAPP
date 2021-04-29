import React, { useRef, useState } from "react";
import Swal from "sweetalert2";
//import { useHistory } from "react-router-dom";
import { registrarUsuario } from "../api/api-user";

import Footer from './Footer'
import Navbar from './Navbar'
import "react-datepicker/dist/react-datepicker.css";
import DatePicker from "react-datepicker";

const  Registrar =()=> {
    
    const [startDate, setStartDate] = useState(new Date());
    const usuarioRef = useRef();
    
    const nombreRef = useRef();
    const apellidoRef = useRef();
    const nacimientoRef = useRef();
    const correoRef = useRef();
    const fotoRef = useRef();
    const passwordRef = useRef();
    const confpasswordRef = useRef();

    const [loading, setLoading] = useState(false);
    const [selectedFiles, setSelectedFiles] = useState([]);
    //const history = useHistory();
  
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
        //verifico contraseñas
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

      //verifico correo
      var emailRegex = /^[-\w.%+]{1,64}@(?:[A-Z0-9-]{1,63}\.){1,125}[A-Z]{2,63}$/i;
      if (!emailRegex.test(correoRef.current.value)) {
        Swal.fire({
          title: "Error",
          text: "Correo invalido",
          icon: "error",
          confirmButtonText: "Ok",
        });
        //window.location.reload(false);
        return;
      } 


      //verifico edad
      //var fechaEnMiliseg = Date.now();
      var age = calcularEdad(startDate);
      //console.log(age)
      if (age < 18){
        Swal.fire({
          title: "Error",
          text: "No se pueden registrar menores de edad",
          icon: "error",
          confirmButtonText: "Ok",
        });
        //window.location.reload(false);
        return;
      }

      
      try {
          var fecha = startDate.getUTCDate() + "/"+(startDate.getUTCMonth() + 1) +"/"+ startDate.getUTCFullYear()
          //console.log(fecha)
        const rawResponse = await registrarUsuario(
          usuarioRef.current.value,
          nombreRef.current.value,
          apellidoRef.current.value,
          fecha,
          correoRef.current.value,
          selectedFiles,
          passwordRef.current.value,
          
          
          
          
        );
  
        const respuesta = await rawResponse.json()
  
        
  
        if(rawResponse.status === 201){
          
          
          Toast.fire({
            icon: 'success',
            title: `¡Se registro@ ${nombreRef.current.value}!`
          })
          //history.push('/admin/perfil')
        } else{
          console.log(respuesta);
          Toast.fire({
            icon: 'error',
            title: 'No se registro, valide informacion'
          })
        }
      } catch (error) {
        console.log(error);
        Toast.fire({
          icon: 'error',
          title: 'No se pudo registrar al usuario'
        })
      } 
      setLoading(false);
    }
  

    return (
        <>
        <Navbar />
        <div className="container mx-auto px-4 h-full">
          <div className="flex content-center items-center justify-center h-full">
            <div className="w-full lg:w-6/12 px-4">
              <div className="relative flex flex-col min-w-0 break-words w-full mb-6 shadow-lg rounded-lg bg-gray-300 border-0">
                <div className="rounded-t mb-0 px-6 py-6">
                  <div className="text-center mb-3">
                    <h6 className="text-gray-600 text-sm font-bold">Registro</h6>
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
                      <p/>
                      
                      <input align ="right"
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
                        Nombre:
                      </label>
                      <p/>
                      <input
                        type="text"
                        ref={nombreRef}
                        className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                        placeholder="Nombre"
                        required
                      />
                    </div>
                    <div className="relative w-full mb-3">
                      <label
                        className="block uppercase text-gray-700 text-xs font-bold mb-2"
                        htmlFor="grid-password"
                      >
                        Apellido:
                      </label>
                      <p/>
                      <input
                        type="text"
                        ref={apellidoRef}
                        className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                        placeholder="Apellido"
                        required
                      />
                    </div>
                    <div className="relative w-full mb-3">
                      <label
                        className="block uppercase text-gray-700 text-xs font-bold mb-2"
                        htmlFor="grid-password"
                      >
                        Fecha de nacimiento
                      </label>
                      <p/>
                      <div>
                 
                    <DatePicker ref={nacimientoRef} selected={startDate} onChange={date => setStartDate(date) } 
                     format={"dd-MM-yyyy"}
                    />
                    </div>
                    </div>
                    <div className="relative w-full mb-3">
                      <label
                        className="block uppercase text-gray-700 text-xs font-bold mb-2"
                        htmlFor="grid-password"
                      >
                        Correo
                      </label>
                      <p/>
                      <input
                        type="text"
                        ref={correoRef}
                        className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                        placeholder="Correo"
                        required
                      />
                    </div>

                    <div className="relative w-full mb-3">
                      <label
                        className="block uppercase text-gray-700 text-xs font-bold mb-2"
                        htmlFor="grid-password"
                      >
                        Foto de perfil
                      </label>
                      
                      <div className="flex items-center justify-center bg-grey-lighter">
                        <label className="w-64 flex flex-col items-center px-4 py-6 bg-white text-blue rounded-lg shadow-lg tracking-wide uppercase border border-blue cursor-pointer hover:bg-blue hover:text-white">
                  
                          <span className="mt-2 text-base leading-normal">
                            Seleccionar archivo
                          </span>
                          <input
                            type="file"
                            name="myFile" 
                            className="hidden"
                            ref={fotoRef}
                            onChange={(e) => setSelectedFiles(e.target.files)}
                            required
                          />
                        </label>
                      </div>
                    </div>
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
  
                    <div className="text-center mt-6">
                      <button
                        className="bg-green-900 text-black active:bg-gray-700 text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 w-full ease-linear transition-all duration-150"
                        type="submit"
                        disabled={loading}
                      >
                        Registrarme
                      </button>
                    </div>
                  </form>
                </div>
              </div>
            </div>
          </div>
        </div>
        <Footer/>
      </>
    )
}

function calcularEdad(fecha) {
  var hoy = new Date();
  var cumpleanos = new Date(fecha);
  var edad = hoy.getFullYear() - cumpleanos.getFullYear();
  var m = hoy.getMonth() - cumpleanos.getMonth();

  if (m < 0 || (m === 0 && hoy.getDate() < cumpleanos.getDate())) {
      edad--;
  }

  return edad;
}

function password_validate(p) {
  return /[A-Z]/.test(p)&& /[a-z]/.test(p) && /[0-9]/.test(p) && /^[A-Za-z0-9]{8,20}$/.test(p) ;
}

export default Registrar;
