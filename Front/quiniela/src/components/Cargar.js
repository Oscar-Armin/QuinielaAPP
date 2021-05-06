import React from "react";
import Footer from './Footer'
import NavbarAdmin from './NavbarAdmin'
import { useHistory } from 'react-router-dom';
import Swal from "sweetalert2";
import YAML from 'yaml'
import { cargar_usuario,cargar_temporadas,carga_rmembresia,carga_jornada,carga_deporte,carga_equipo,carga_partido } from "../api/api-carga";

import { forgetUsuario } from "../api/api-user";
//import { useHistory } from 'react-router-dom';


function Cargar() {
    
    let fileReader;
    //const [selectedFiles, setSelectedFiles] = useState([]);
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
    var cadena = "";
    const handleFileRead = (e) =>{
        cadena= fileReader.result;
        //cadena = fileReader.result;
        //console.log(content)
    }

    const handleFileChosen = (file) =>{
        fileReader = new FileReader();
        fileReader.onloadend = handleFileRead;
        fileReader.readAsText(file);
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



      async function handleSubmit(e) {
        e.preventDefault();
        let bandera = false;
        //setLoading(true);
    
        try {
          /*const rawResponse = await forgetUsuario(
            selectedFiles
            
          );*/
    
          //console.log(cadena)
          const resa = YAMLtoJSON(cadena)
          const entrada = JSON.parse(resa)
          let id_user =0
          let rawResponse
          let respuesta
          let mes
          let fecha_inicio
          let fecha_fin
          for (let i in entrada) {
            //console.log("--------------------")
            //console.log(entrada[i].nombre)
            //console.log(entrada[i].apellido)
            //console.log(entrada[i].password)
            //console.log(entrada[i].username)

            //carga de usuario            
            rawResponse = await cargar_usuario(//--------------------------
                
                entrada[i].username.replace("'", ""),
                entrada[i].nombre.replace("'", ""),
                
                entrada[i].apellido.replace("'", ""),
                entrada[i].password,
                
              );
              if(rawResponse.status === 201){
                  bandera = true;
                console.log("Se cargo usuario "+entrada[i].username)

              }else if(rawResponse.status === 500){
                bandera = false;
                console.log("Error al cargo usuario "+entrada[i].username)
              }else{
                bandera = false;
                console.log("Error al cargo usuario "+entrada[i].username)
              }

             rawResponse = await forgetUsuario(entrada[i].username);
              if(rawResponse.status === 201){
                respuesta = await rawResponse.json();
                id_user = respuesta.ID;
              console.log(id_user)

            }else if(rawResponse.status === 500){
              bandera = false;
              console.log("Error al cargo usuario "+entrada[i].username)
            }else{
              bandera = false;
              console.log("Error al cargo usuario "+entrada[i].username)
            }

            
            for (let j in entrada[i].resultados) {
                let strTMP = entrada[i].resultados[j].temporada
                let anio = strTMP.substring(0,4)
                
                if(strTMP.length === 7){
                  mes = strTMP.substring(6,7)
                }else{
                  mes = strTMP.substring(6,8)
                }
                
                fecha_inicio = "01/"+mes+"/"+anio
                fecha_fin = "30/"+mes+"/"+anio
                
                //ingreso de temporadas
                rawResponse = await cargar_temporadas(//------------------------
                  strTMP,
                  fecha_inicio,
                  fecha_fin
                );
                //console.log(entrada[i].resultados[j].tier)
                //Ingreso de las membresias que tenia el usuario
                
                
                rawResponse = await carga_rmembresia(//---------------------------------------------
                  id_user,
                  strTMP,
                  entrada[i].resultados[j].tier
                );
                
                for (let k in entrada[i].resultados[j].jornadas) {
                  
                  //carga_jornada
                  rawResponse = await carga_jornada(//---------------------------------------------
                    parseInt(entrada[i].resultados[j].jornadas[k].jornada.substring(1,2)),
                    strTMP
                    
                  );
                  
                  for (let l in entrada[i].resultados[j].jornadas[k].predicciones) {
                    //carga_deporte 
                    
                    rawResponse = await carga_deporte(//---------------------------------------------
                      entrada[i].resultados[j].jornadas[k].predicciones[l].deporte
                      
                    );


                    //carga equioos
                    
                    rawResponse = await carga_equipo(//---------------------------------------------
                      entrada[i].resultados[j].jornadas[k].predicciones[l].visitante
                      
                    );
                    rawResponse = await carga_equipo(//---------------------------------------------
                      entrada[i].resultados[j].jornadas[k].predicciones[l].local
                      
                    );
                    
                    
                    rawResponse = await carga_partido(//---------------------------------------------
                      entrada[i].resultados[j].jornadas[k].predicciones[l].local,
                      entrada[i].resultados[j].jornadas[k].predicciones[l].visitante,
                      parseInt(entrada[i].resultados[j].jornadas[k].predicciones[l].resultado.local),
                      parseInt(entrada[i].resultados[j].jornadas[k].predicciones[l].resultado.visitante),
                      entrada[i].resultados[j].jornadas[k].predicciones[l].fecha,
                      entrada[i].resultados[j].jornadas[k].predicciones[l].deporte,
                      parseInt(entrada[i].resultados[j].jornadas[k].jornada.substring(1,2)),
                      strTMP
                    );
                      console.log("----------------------")
                  }
                  
                }
            }

          }
            
          
          if (true){
                  console.log(bandera)
                  Toast.fire({
                    icon: "success",
                    title: `¡Se le cargo correctamente la informacion !`,
                  });
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
            title: "No se pudo cargar datos, error en archivo",
          });
        }
    
        //setLoading(false);
      }
    
    return (
        <div>
          
            <NavbarAdmin />
            
            
            <div className="container mx-auto px-4 h-full">
        <div className="flex content-center items-center justify-center h-full">
          <div className="w-full lg:w-4/12 px-4">
            <div className="relative flex flex-col min-w-0 break-words w-full mb-6 shadow-lg rounded-lg bg-gray-300 border-0">
              <div className="rounded-t mb-0 px-6 py-6">
                <div className="text-center mb-3">
                  <h6 className="text-gray-600 text-sm font-bold">
                    Carga masiva de datos
                  </h6>
                  
                </div>

                <hr className="mt-6 border-b-1 border-gray-400" />
              </div>
              <div className="flex-auto px-4 lg:px-10 py-10 pt-0">
                <div className="text-gray-500 text-center mb-3 font-bold">
                  <small>Seleccione archivo yaml</small>
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
                    {/*<input
                            type="file"
                            name="myFile" 
                            className="hidden"
                            
                            onChange={(e) => setSelectedFiles(e.target.files)}
                            required
                          />*/}

                <input
                            type="file"
                            id="file" 
                            className="input-file"
                            accept = '.yaml'
                            
                            onChange={e => handleFileChosen(e.target.files[0])}
                            required
                          />
                  </div>


                  <div align="center"> 
                    <button
                      className="bg-gray-900 text-black active:bg-gray-700 text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 w-full ease-linear transition-all duration-150"
                      type="submit"
                    >
                      Cargar
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

function YAMLtoJSON(yamlStr) { 
    
	var obj = YAML.parse(yamlStr); 
	var jsonStr = JSON.stringify(obj); 
	return jsonStr; 
} 

export default Cargar
