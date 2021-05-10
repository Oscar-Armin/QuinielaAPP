import React, { useRef, useState } from "react";
import Footer from './Footer';
import NavbarAdmin from './NavbarAdmin';
import { obtener__colores ,obtener_deportes,registro_deporte,deleteDeporte } from "../api/api-admin";
import { useHistory } from 'react-router-dom';

import Swal from "sweetalert2";
  

function Admin_deportes() {
  let history = useHistory();
  var a = localStorage.getItem("usuarioActual")
  
  if (a === null){
      history.push('/home')
      
  }
  var obj = JSON.parse(a)
  
  if(obj.id_usuario !==1){
      history.push('/usuario')
      
  }

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
  

    const [selectedFiles, setSelectedFiles] = useState([]);
    const deporteRef = useRef();
    let [color, setFruit] = useState("⬇️ Select a color ⬇️")
    let [sport, setSport] = useState("⬇️ Select a sport ⬇️")

    let handleFruitChange = (e) => {
        setFruit(e.target.value)
      }
    let handleSportChange = (e) => {
      setSport(e.target.value)
      console.log(sport)
      }

    
    
    evento();
    async function handleSubmit(e) {
        e.preventDefault();
        
        const rawResponse = await registro_deporte(
          
          deporteRef.current.value,
          color,
          selectedFiles,
          
          
          
          
          
        );
        if(rawResponse.status === 201){
          
          
          Toast.fire({
            icon: 'success',
            title: `¡Se registro@ ${deporteRef.current.value}!`
            
          })
          history.push('/Adeportes')
          //history.push('/admin/perfil')
        } else{
          Toast.fire({
            icon: 'error',
            title: `No se pudo registrar`
          })
        }
    }
    async function handleDelete(e) {
      e.preventDefault();
      console.log("----")
      console.log(sport)
      console.log("----")
      const rawResponse = await deleteDeporte(
        
        sport
        
        
        
        
      );
      console.log( rawResponse.status )
      if(rawResponse.status === 201){
        
        
        Toast.fire({
          icon: 'success',
          title: `¡Se elimino !`
        })
        history.push('/Adeportes')
        //history.push('/admin/perfil')
      } else{
        Toast.fire({
          icon: 'error',
          title: `Ocurrio un error`
        })
      }
  }
    var entrada = JSON.parse(localStorage.getItem("colores"))
    var jsonDeporte = JSON.parse(localStorage.getItem("deportes"))
    
    return (


        <div>
            <NavbarAdmin />
            
            <br />
                <div className="border"  style={{background: 'lightgrey'}} >
                    <h2> Creacion </h2>
                    <form onSubmit={handleSubmit}>
                    
                    <div className="contenedor">
                        <div className="relative w-full mb-3">
                            <label
                                className="block uppercase text-gray-700 text-xs font-bold mb-2"
                                htmlFor="grid-password"
                            >
                                Deporte:
                            </label>
                            
                            
                            <input align ="right"
                                type="text"
                                ref={deporteRef}
                                className="px-3 py-3 placeholder-gray-400 text-gray-700 bg-white rounded text-sm shadow focus:outline-none focus:shadow-outline w-full ease-linear transition-all duration-150"
                                placeholder="Deporte"
                                required
                            />
                            
                        </div>
                    
                    
                    
                    <div className="relative w-full mb-3">
                    <label
                        className="block uppercase text-gray-700 text-xs font-bold mb-2"
                        htmlFor="grid-password"
                      >
                        Color:
                      </label>
                      <br></br>
                        <select onChange={handleFruitChange}>
                            <option value="⬇️ Select a color ⬇️">---Select a color---</option>
                            {entrada.arreglo.map((option) => (
                            <option value={option.color}>{option.color}</option>
                          ))}
                        </select>
                    </div>
                    <div className="relative w-full mb-3">
                      <label
                        className="block uppercase text-gray-700 text-xs font-bold mb-2"
                        htmlFor="grid-password"
                      >
                        Deporte:
                      </label>
                        <br></br>
                        <input
                            type="file"
                            name="myFile" 
                            className="hidden"
                            
                            onChange={(e) => setSelectedFiles(e.target.files)}
                            required
                          />
                          </div>
                          <div className="relative w-full mb-3">
                          <button
                        className="bg-green-900 text-black active:bg-gray-700 text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 w-full ease-linear transition-all duration-150"
                        type="submit"
                        
                      >
                        Registrarme
                      </button>
                          </div>
                          <div></div>
                          <div></div>
                          </div>
                    </form>
                    <br></br>
                </div>
                <div className="border"  style={{background: 'lightgrey'}} >
                    <h2> Eliminacion/Editar </h2>
                    <select onChange={handleSportChange}>
                            <option value="⬇️ Select a color ⬇️">---Select a deporte---</option>
                            {jsonDeporte.arreglo.map((option) => (
                            <option value={option.id}>{option.name}</option>
                          ))}
                        </select>
                    <button
                        
                        onClick ={handleDelete}                       
                      >
                        Eliminar
                      </button>
                      <button
                        className="bg-green-900 text-black active:bg-gray-700 text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 w-full ease-linear transition-all duration-150"
                        type="submit"
                        
                      >
                        Editar
                      </button>
                    
                    
                </div>
                <div className="border"  style={{background: 'lightgrey'}} >
                    <h2> Editar </h2>
                    
                </div>
            <Footer />
        </div>
    )
}

async function evento() {
    let rawResponse = await obtener__colores();
    if (rawResponse.status === 201) {
        const rColores = await rawResponse.json()
        
        localStorage.setItem(
            "colores",
            JSON.stringify({
              arreglo:rColores
              
              
            })
          );
    
    } else {
        console.log(rawResponse.status)
        localStorage.setItem(
            "colores",
            JSON.stringify({
              
              
              
            })
          );

    }
    rawResponse = await obtener_deportes();
    if (rawResponse.status === 201) {
      const rdeportes = await rawResponse.json()
      
      localStorage.setItem(
          "deportes",
          JSON.stringify({
            arreglo:rdeportes
            
            
          })
        );
  
  } else {
      console.log(rawResponse.status)
      localStorage.setItem(
          "deportes",
          JSON.stringify({
            
            
            
          })
        );

  }
}

export default Admin_deportes
