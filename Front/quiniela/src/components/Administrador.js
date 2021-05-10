import React from 'react'
import Footer from './Footer'
import NavbarAdmin from './NavbarAdmin'
import { useHistory } from 'react-router-dom';
import { obtener_temporada,obtener_dinero,obtener_oros,obtener_platas,obtener_bronces } from "../api/api-admin";

var nomberTemporada;
var dinero=0;
var oros = 0;
var platas = 0;
var bronces = 0;

const Administrador = ()=>{
    
    let history = useHistory();
    var a = localStorage.getItem("usuarioActual")
    
    if (a === null){
        history.push('/home')
        
    }
    var obj = JSON.parse(a)
    
    if(obj.id_usuario !==1){
        history.push('/usuario')
        
    }
    evento();
    
    nomberTemporada = JSON.parse(localStorage.getItem('temporada')).nombre;
    dinero = parseInt(JSON.parse(localStorage.getItem('temporada')).dinero);
    oros = parseInt(JSON.parse(localStorage.getItem('temporada')).gold);
    platas = parseInt(JSON.parse(localStorage.getItem('temporada')).silver);
    bronces = parseInt(JSON.parse(localStorage.getItem('temporada')).bronze);
    //localStorage.removeItem('temporada');




    
    return (
        <div>
            <NavbarAdmin />
            <p></p>
            <br></br>
            
            <div className="border border-primary" style={{border: '1px solid black',borderRadius: '10px', float: 'center'}}>
                <label style={{background:'skyblue',color:'black'}} >Temporada: {nomberTemporada} </label>
                <br />
                <br />
                <br />
                En juego: Q{dinero}.00
                
                <br />
                <br />
                Participantes: 
                <br />
                <br />
                
                <center>
               
                </center>
                <div style = {{display: 'flex'}}> 
                    <div className="circuloG">
                        <h2> {oros}</h2>
                    </div>
                    <div className="circuloS">
                        <h2> {platas}</h2>
                    </div>
                    <div className="circuloB">
                        <h2> {bronces}</h2>
                    </div>
                </div>
                <br />
                <br />
                <br />
                <br />
                <br />
                <br />
                
            
            </div>
            <p></p>
            <br></br>
            

            <Footer />
            
        </div>
    )
}
async function evento(nombre) {
    let rawResponse = await obtener_temporada();
    if (rawResponse.status === 201) {
        let respuesta = await rawResponse.json();
        nomberTemporada= respuesta.name;
        
        
       rawResponse = await obtener_dinero(nomberTemporada);
        respuesta = await rawResponse.json();
        dinero = respuesta.cantidad;
        

        rawResponse = await obtener_oros(nomberTemporada);
        respuesta = await rawResponse.json();
        oros = respuesta.cantidad;
        

        rawResponse = await obtener_platas(nomberTemporada);
        respuesta = await rawResponse.json();
        platas = respuesta.cantidad;
        

        rawResponse = await obtener_bronces(nomberTemporada);
        respuesta = await rawResponse.json();
        bronces = respuesta.cantidad;
        
        localStorage.setItem(
            "temporada",
            JSON.stringify({
              nombre: nomberTemporada,
              dinero: dinero,
              gold: oros,
              silver:platas,
              bronze:bronces
              
              
            })
          );


    
    } else {
        console.log('No hay temporada actual')
        localStorage.setItem(
            "temporada",
            JSON.stringify({
              nombre: "",
              dinero: "0",
              gold: "0",
              silver:"0",
              bronze:"0"
              
              
            })
          );

    }
}
    
export default Administrador;