import React from "react";
import Footer from './Footer'
import NavbarAdmin from './NavbarAdmin'
import { useHistory } from 'react-router-dom';
import Swal from "sweetalert2";
import YAML from 'yaml'
import { cargar_usuario } from "../api/api-carga";
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
          for (let i in entrada) {
            //console.log("--------------------")
            //console.log(entrada[i].nombre)
            //console.log(entrada[i].apellido)
            //console.log(entrada[i].password)
            //console.log(entrada[i].username)

            
            const rawResponse = await cargar_usuario(
                
                entrada[i].username,
                entrada[i].nombre,
                
                entrada[i].apellido,
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

          }
            
          //console.log(respuesta.ID)
          
          //if (rawResponse.status === 201) {
              if (bandera){

              //console.log(selectedFiles[0].result)
            //const respuesta = await rawResponse.json();
            //console.log(respuesta)

                    //console.log(respuesta.ID)
                  Toast.fire({
                    icon: "success",
                    title: `¡Se le cargo correctamente la informacion !`,
                  });
/*
            
          } else if(rawResponse.status === 500){
            Toast.fire({
              icon: 'error',
              title: 'Usuario no existe'
            })
          } else {
            Toast.fire({
              icon: "error",
              title: "No se pudo iniciar sesión",
            });*/
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
