import React,{  useState,useEffect ,useRef}   from 'react';
import Footer from './Footer'
import NavbarAdmin from './NavbarAdmin'
import FullCalendar from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import interactionPlugin from "@fullcalendar/interaction";
import listPlugin from '@fullcalendar/list';
import time from '@fullcalendar/timegrid';
import { useHistory } from 'react-router-dom';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import Swal from "sweetalert2";
import { cambio_fase2,cambio_fase3 ,end_jornada,set_res} from "../api/api-admin";



function AJornadas() {
  const puntosV = useRef();
    const puntosL = useRef();
    const [total, setTotal] = useState([]);  
    const [estado, setEstado] = useState(3);  
    const [open3, setOpen3] = React.useState(false);
    const [open2, setOpen2] = React.useState(false);
    const [open1, setOpen1] = React.useState(false);

    const [titulo, setTitle] = useState(" - ");  
    const [mesage1, setMensaje1] = useState(" - ");  
    const [mesage2, setMensaje2] = useState(" - ");  
    const [mesage3, setMensaje3] = useState(" - ");  

    const [local, setLocal] = useState(0);  
    const [visitante, setVisitante] = useState(0);  

    const [partido, setPartido] = useState(0);  


    let history = useHistory();
    var a = localStorage.getItem("usuarioActual")
    const handleClose3 = () => {
        setOpen3(false);
      };

      const handleClose2 = () => {
        setOpen2(false);
      };
      const handleClose1 = () => {
        setOpen1(false);
      };

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
        
        fetch("http://localhost:3080/partido_ja", { method : 'POST', 
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                mensaje: "mensaje"
            }) })
            .then((response) => response.json())
            .then(data => {
               
                    setTotal(data) 
                
                
                });
        },[])

      useEffect(() => {
      
          fetch("http://localhost:3080/estado_fase", { method : 'POST', 
              headers: {
                  Accept: "application/json",
                  "Content-Type": "application/json",
              },
              body: JSON.stringify({
                  mensaje: "mensaje"
              }) })
              .then((response) => response.json())
              .then(data => {
                  
                  setEstado(data.ID) 
                  
                  
                  });
          })


    if (a === null){
        history.push('/home')
        
    }
    var obj = JSON.parse(a)
    
    if(obj.id_usuario !==1){
        history.push('/usuario')
        
    }
    console.log(estado)
    async function handleFase2(e) {
      e.preventDefault();
      if (estado === 1){
        const rawResponse = await cambio_fase2();
        if (rawResponse.status === 201) {
          Toast.fire({
            icon: "success",
            title: `¡Se cambio a fase 2!`,
          });
        }else{
          Toast.fire({
            icon: "error",
            title: `No se pudo cambiar`,
          });
        }
      }else{
        Toast.fire({
          icon: "error",
          title: `No se pudo cambiar`,
        });
      }
      
    }

    async function handleFase3(e) {
      e.preventDefault();
      if (estado === 2){
        const rawResponse = await cambio_fase3();
        if (rawResponse.status === 201) {
          Toast.fire({
            icon: "success",
            title: `¡Se cambio a fase 3!`,
          });
        }else{
          Toast.fire({
            icon: "error",
            title: `No se pudo cambiar`,
          });
        }
      }else{
        Toast.fire({
          icon: "error",
          title: `No se pudo cambiar`,
        });
      }
      
    }

    async function handleJornada(e) {
      e.preventDefault();
      if (estado === 2){
        const rawResponse = await end_jornada();
        if (rawResponse.status === 201) {
          Toast.fire({
            icon: "success",
            title: `¡Se cambio a fase 3!`,
          });
        }else{
          Toast.fire({
            icon: "error",
            title: `No se pudo cambiar`,
          });
        }
      }else{
        Toast.fire({
          icon: "error",
          title: `No se pudo cambiar`,
        });
      }
      
    }

    async function handleResultado(e) {
      e.preventDefault();
      console.log(partido)
      console.log(parseInt(puntosL.current.value))
      console.log(parseInt(puntosV.current.value))
      if (estado === 2){
        const rawResponse = await set_res(
          parseInt(partido),
          parseInt(puntosL.current.value),
          parseInt(puntosV.current.value)

        );
        if (rawResponse.status === 201) {
          
          Toast.fire({
            icon: "success",
            title: `¡Se guardaron los puntos!`,
          });
        }else{
          Toast.fire({
            icon: "error",
            title: `No se pudo cambiar`,
          });
        }
      }else{
        Toast.fire({
          icon: "error",
          title: `No se pudo cambiar`,
        });
      }
      
    }




    async function handleDateSelect(e) {
        if (estado === 1){
          setTitle(e.event.title)
          setOpen1(true)
        } else if (estado === 2){
          setPartido(e.event.id)
          setLocal(e.event.extendedProps.resultadoL)
          setVisitante(e.event.extendedProps.resultadoV)

          setTitle(e.event.title)
          setMensaje2(" vs ")
          var mensaje3 = e.event.start.toString();

          setMensaje3(mensaje3)
          setOpen2(true)
        }else if (estado === 3){
            
            
            var mensaje1 = ""
            setTitle(e.event.title)
            if ( parseInt(e.event.extendedProps.prediccion) === 0){
              mensaje1 = "/ - /"
            setMensaje1(mensaje1)
            }else{
              mensaje1 = e.event.extendedProps.resultadoL+" - " +e.event.extendedProps.resultadoV 
            setMensaje1(mensaje1)
            }
            
            var mensaje2 = e.event.start.toString();
            setMensaje2(mensaje2)
            setOpen3(true)
        }


        
      }
    return (
        <div>
            <NavbarAdmin />
            <center>
            <div style = {{display : 'flex'}}>
              
              <button type="button" className="btn btn-primary" onClick ={handleFase2}>Fase 2</button>
              <button type="button" className="btn btn-primary" onClick ={handleFase3}>Fase 3</button>
              <button type="button" className="btn btn-primary" onClick ={handleJornada}>Terminar Jornada</button>

            </div>
            </center>

            <br></br>
            <div id="Calendario">
                    <br/> <br/>
                    <FullCalendar
                        plugins={[ dayGridPlugin, listPlugin, time, interactionPlugin]}
                        initialView="dayGridMonth"
                        headerToolbar={{
                            left: 'dayGridMonth,timeGridWeek,listYear',
                            center: 'title,today',
                            right: 'prevYear,prev,next,nextYear'
                        }}
                        editable={true}
                        selectable={true}
                        selectMirror={true}
                        dayMaxEvents={true}
                        locale = {'es'}
                        events={ total}
                          eventClick ={handleDateSelect}
                        //select={handleDateSelect}
                        //initialEvents = {Ingresar_eventos()}
                        //events={this.state.calendarEvents}
                        //select={handleDateSelect}
                        //eventContent={renderEventContent} // custom render function
                        //eventClick={this.handleEventClick}
                    />
                </div>
                {/*3 */}
                <Dialog
        open={open3}
        onClose={handleClose3}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">{titulo}</DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            {mesage1}
            <br></br>
            {mesage2}
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose3} color="primary">
            Ok
          </Button>
         
        </DialogActions>
      </Dialog>
                        {/*2 */}
                     <   Dialog
        open={open2}
        onClose={handleClose2}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">{titulo}</DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
          
          <br></br>
          <input  ref={puntosL} type="number"   min="0" max="50" step="1" defaultValue ={local}/>
            
            
            {mesage2}
            
            <input ref={puntosV} type="number" name="numero"  min="0" max="50" step="1" defaultValue ={visitante}/>
            <br></br>
            <br />
            {mesage3}
            
          </DialogContentText>
        </DialogContent>
        <DialogActions>
        <Button onClick={handleResultado} color="primary">
            guardar cambios
          </Button>
          <Button onClick={handleClose2} color="primary">
            cancelar
          </Button>
          
        </DialogActions>
      </Dialog>
      {/*1*/}
      <   Dialog
        open={open1}
        onClose={handleClose1}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">{titulo}</DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            {mesage1}
            <br></br>
            {mesage2}
            <br></br>
            {mesage3}

          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose1} color="primary">
            Ok
          </Button>
         
        </DialogActions>
      </Dialog>
            
            <Footer />
        </div>
    )
}






export default AJornadas
