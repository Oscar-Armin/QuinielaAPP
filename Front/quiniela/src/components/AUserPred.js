import React,{  useState,useEffect }   from 'react';
import Footer from './Footer'
import NavbarAdmin from './NavbarAdmin'
import {withRouter} from 'react-router-dom';
import {obtener_prediccion} from "../api/api-admin";

import Swal from "sweetalert2";

function AUserPred(props) {
    let identificador = parseInt(props.match.params.id)
    const [nombre, setName] = useState("name"); 
    const [temporadas, setTemp] = useState([]);  
    const [predic, setPredi] = useState([{"deporte":"foot","local":"local","visitante":"visitante","plocal":0,"pvisitante":0,"rlocal":0,"rvisitante":0,"puntos":0,"fecha":"11/05/2021"}]);  
    let [temp, setT] = useState(1801)

    let handleSportChange = (e) => {
        setT(e.target.value)
        
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
    
    
    useEffect(() => {
        
        fetch("http://localhost:3080/name_user", { method : 'POST', 
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                ID: parseInt(identificador)
            }) })
            .then((response) => response.json())
            .then(data => {
                setName(data.nombre) 
                });


        
    
},)

    useEffect(() => {
            
        fetch("http://localhost:3080/obtener_temporadas", { method : 'POST', 
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                mensaje: "mensaje"
            }) })
            .then((response) => response.json())
            .then(data => {
                setTemp(data) 
                });


        

    },[])

    async function handleSubmit(e) {
        console.log(typeof(identificador))
            console.log(typeof(temp))
          
        e.preventDefault();
        
        const rawResponse = await obtener_prediccion(
          
            identificador,
            temp
          
          
          
          
          
        );
        if(rawResponse.status === 201){
        //    if(true){

            const respuesta = await rawResponse.json();
            //setPredi(respuesta);
            if(respuesta === null){
                console.log("esta nulll")
            }else{
                console.log("cargue")
                setPredi(respuesta);
            }
            
          Toast.fire({
            icon: 'success',
            title: `¡Se oltuvo la info!`
            
          })
          console.log(respuesta)
          
          //history.push('/admin/perfil')
        } else{
            setPredi([{"deporte":"foot","visitante":"visitante","plocal":0,"pvisitante":0,"rlocal":0,"rvisitante":0,"puntos":0,"fecha":"11/05/2021"}])
          Toast.fire({
            icon: 'error',
            title: `No se pudo registrar`
          })
          console.log(predic)
          console.log("fallo")
        }
    }

    console.log(predic)

    return (
        <div>
            
            <NavbarAdmin />
            <br></br>
            <right><label style={{background :'skyblue' }} align ='right'>{nombre}</label></right>
            <br></br>
            <div >
                
                <center>
            <select class="form-select form-select-lg mb-3" aria-label=".form-select-lg example" onChange={handleSportChange}>
            
                            <option value="⬇️ Select a season ⬇️" >---Select a color---</option>
                            {temporadas.map((option) => (
                            <option value={option.ID}  >{option.correo}</option>
                          ))}
                        </select>

                        <button type="button" class="btn btn-info" onClick ={handleSubmit}>Cargar Predicciones</button>
                        </center>
                        </div>
                        
                        <table className = "table table-sm">
                    <thead>
                        <tr>
                            <th>Deporte</th>   
                            <th>Local</th>
                            <th>Visitante</th>
                            <th>Prediccion</th>
                            
                            <th>Resultado</th>
                            
                            <th>Puntos</th>
                            <th>Fecha</th>
                        </tr>
                    </thead>
                    <tbody>
                    {predic.map((option) => (
                        <tr > 
                            
                            
                            <td>{option.deporte}</td>
                            <th>{option.local}</th>
                            <th>{option.visitante}</th>
                            
                            <th>{option.plocal}-{option.pvisitante}</th>
                            <th>{option.rlocal}-{option.rvisitante}</th>
                            <th>{option.puntos}</th>
                            <th>{option.fecha}</th>
                            
                            
                        </tr>
                          ))}
                    </tbody>

                </table>
            <Footer /> 
        </div>
    )
}


export default withRouter(AUserPred);
