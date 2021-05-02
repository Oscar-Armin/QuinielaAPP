import React, { useRef, useState  } from "react";
import Footer from './Footer'
import NavbarUser from './NavbarUser'
import { useHistory } from 'react-router-dom';

import Swal from "sweetalert2";



import { actualizarUsuario } from "../api/api-user";


const Perfil = ()=>{
    
    var md5 = require('md5');
    
    const usuarioRef = useRef();
    
    const nombreRef = useRef();
    const apellidoRef = useRef();
    
    const correoRef = useRef();
    const fotoRef = useRef();
    const passwordRef = useRef();
    const confpasswordRef = useRef();

    const [loading, setLoading] = useState(false);
    const [selectedFiles, setSelectedFiles] = useState([]);

    let history = useHistory();
    var a = localStorage.getItem("usuarioActual")
    //console.log(a)
    if (a === null){
        history.push('/home')
        return(<></>);
    }
    
    //console.log(obj.id_usuario)
    
    
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
        var contra = md5(passwordRef.current.value)
          //verifico contraseñas
        if (contra !== JSON.parse(localStorage.getItem('editar_user')).password) {
          Swal.fire({
            title: "Error",
            text: "Las contraseñas no coinciden",
            icon: "error",
            confirmButtonText: "Ok",
          });
          //window.location.reload(false);
          return;
        }
  
        var contras = ""
        
        
        if (confpasswordRef.current.value){
            if ( !password_validate(confpasswordRef.current.value)) {
                Swal.fire({
                  title: "Error",
                  text: "contraseña nueva invalida*",
                  icon: "error",
                  confirmButtonText: "Ok",
                });
                //window.location.reload(false);
                
                return;
            } else{
                contras = md5(confpasswordRef.current.value)
            }


        }else{
            contras = JSON.parse(localStorage.getItem('editar_user')).password
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
  
        var photo 
        var auxilar 
        //console.log(JSON.parse(localStorage.getItem('editar_user')).id_usuario)
        //console.log(usuarioRef.current.value)
        //console.log(nombreRef.current.value)
        //console.log(apellidoRef.current.value)
        //console.log(correoRef.current.value)
        //console.log(selectedFiles.length)
        if(selectedFiles.length === 0){
          auxilar = false
            photo =JSON.parse(localStorage.getItem('editar_user')).foto
            //console.log("NO cambie de foto")
            //console.log(photo)
        }else{
            photo = selectedFiles
            auxilar = true
            //console.log("cambie de foto")
        }
        
        //console.log(photo)
        //console.log(contras)
        
        try {
            
            //console.log(fecha)
          const rawResponse = await actualizarUsuario(
            JSON.parse(localStorage.getItem('editar_user')).id_usuario,
            usuarioRef.current.value,
            nombreRef.current.value,
            apellidoRef.current.value,
            
            correoRef.current.value,
            photo,
            contras,
            auxilar,
            JSON.parse(localStorage.getItem('usuarioActual')).usuario
            
            
            
          );
    
          
    
          
    
          if(rawResponse.status === 201){
        //if(true){
//            const respuesta = await rawResponse.json()  
            
            
            Toast.fire({
              icon: 'success',
              title: `¡Se actualizo@ ${nombreRef.current.value}!`
            })
            //history.push('/admin/perfil')
          } else{
            //console.log(respuesta);
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
      if(!JSON.parse(localStorage.getItem('editar_user')).usuario){
        window.location.reload(false);
      }
      
    return (

        <div>
            

            <NavbarUser />
            
            
            
            
            <div className="container mx-auto px-4 h-full">
          <div className="flex content-center items-center justify-center h-full">
            <div className="w-full lg:w-6/12 px-4">
              <div className="relative flex flex-col min-w-0 break-words w-full mb-6 shadow-lg rounded-lg bg-gray-300 border-0">
                <div className="rounded-t mb-0 px-6 py-6">
                  <div className="text-center mb-3">
                    <h6 className="text-gray-600 text-sm font-bold">Informacion {}</h6>
                  </div>
  
                  <hr className="mt-6 border-b-1 border-gray-400" />
                </div>
                <div className="flex-auto px-4 lg:px-10 py-10 pt-0">
                  <div className="text-gray-500 text-center mb-3 font-bold">
                    <small>Cambie si desea actualizar</small>
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
                      placeholder={JSON.parse(localStorage.getItem('editar_user')).usuario}
                        type="text"
                        defaultValue={JSON.parse(localStorage.getItem('editar_user')).usuario}
                        ref={usuarioRef}
                        
                        className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                        
                        
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
                      defaultValue={JSON.parse(localStorage.getItem('editar_user')).nombre}
                        type="text"
                        ref={nombreRef}
                        placeholder={JSON.parse(localStorage.getItem('editar_user')).nombre}
                        className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                        
                        
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
                        defaultValue={JSON.parse(localStorage.getItem('editar_user')).apellido}
                        ref={apellidoRef}
                        
                        className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                        placeholder={JSON.parse(localStorage.getItem('editar_user')).apellido}
                        
                      />
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
                        defaultValue={JSON.parse(localStorage.getItem('editar_user')).correo}
                        ref={correoRef}
                        placeholder={JSON.parse(localStorage.getItem('editar_user')).correo}
                        className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                        
                        
                      />
                    </div>

                    <div className="relative w-full mb-3">
                      <label
                        className="block uppercase text-gray-700 text-xs font-bold mb-2"
                        htmlFor="grid-password"
                      >
                        Foto de perfil
                      </label>
                      <div className="img-holder">
						<img src={JSON.parse(localStorage.getItem('editar_user')).foto} alt="" id="img" className="img" width="500" height="600" />
					</div>
                      <div className="flex items-center justify-center bg-grey-lighter">
                        <label className="w-64 flex flex-col items-center px-4 py-6 bg-white text-blue rounded-lg shadow-lg tracking-wide uppercase border border-blue cursor-pointer hover:bg-blue hover:text-white">
                  
                          <span className="mt-2 text-base leading-normal">
                          
                            Seleccionar archivo
                            <p></p>  
                            <small> Para actulizar </small>
                            <p></p>
                          </span>
                          <input
                            type="file"
                            name="myFile" 
                            className="hidden"
                            ref={fotoRef}
                            onChange={(e) => setSelectedFiles(e.target.files) }
                            
                          />
                        </label>
                      </div>
                    </div>
                    <div className="relative w-full mb-3">
                      <label
                        className="block uppercase text-gray-700 text-xs font-bold mb-2"
                        htmlFor="grid-password"
                      >
                        Contraseña Actual
                        
                      </label>
                      <p />
                      <small>Ingrese contraseña actual si desea actualizar</small>

                      <p/>
                      <input
                        type="password"
                        ref={passwordRef}
                        className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                        placeholder="Contraseña"
                        required
                      />
                      <p />
                      
                    </div>
  
                    <div className="relative w-full mb-3">
                      <label
                        className="block uppercase text-gray-700 text-xs font-bold mb-2"
                        htmlFor="grid-password"
                      >
                        Confirmación nueva
                      </label>
                      <p/>
                      <small>Ingrese contraseña nueva si la desea cambiar</small>
                      <p />
                      <input
                        type="password"
                        ref={confpasswordRef}
                        className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                        placeholder="Confirmar contraseña"
                        
                      />
                    </div>
  
                    <div className="text-center mt-6">
                      <button
                        className="bg-green-900 text-black active:bg-gray-700 text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 w-full ease-linear transition-all duration-150"
                        type="submit"
                        disabled={loading}
                      >
                        Actualizar
                      </button>
                    </div>
                  </form>
                  
                </div>
              </div>
            </div>
          </div>
          
        </div>
            <Footer />
            
            
        </div>
    )
}


  
function password_validate(p) {
    return /[A-Z]/.test(p)&& /[a-z]/.test(p) && /[0-9]/.test(p) && /^[A-Za-z0-9]{8,20}$/.test(p) ;
}


  
export default Perfil;