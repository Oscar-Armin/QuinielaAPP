import React,{  useState,useEffect }   from 'react';
import Footer from './Footer'
import NavbarAdmin from './NavbarAdmin'
import { useHistory } from 'react-router-dom';
import '../../node_modules/react-bootstrap-table/css/react-bootstrap-table.css'



function ARecompensa() {
    const [total, setTotal] = useState([]);  
     
    
    useEffect(() => {
        
        fetch("http://localhost:3080/Arecompensas", { method : 'POST', 
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


    let history = useHistory();
    var a = localStorage.getItem("usuarioActual")
    
    if (a === null){
        history.push('/home')
        
    }
    var obj = JSON.parse(a)
    
    if(obj.id_usuario !==1){
        history.push('/usuario')
        
    }


    return (
        <div>
            <NavbarAdmin />
{/*
            <BootstrapTable data={data}  >
          <TableHeaderColumn isKey dataField='id' onclick ={console.log("asdasd")}>
            Usuario
          </TableHeaderColumn>
          <TableHeaderColumn dataField='name'>
            Nombre
          </TableHeaderColumn>
          <TableHeaderColumn dataField='value'>
            Apellido
          </TableHeaderColumn>
        </BootstrapTable>
*/}
                <table className = "table table-sm">
                    <thead>
                        <tr>
                            <th>Username</th>
                            <th>Nombre</th>
                            <th>Apellido</th>
                            <th>Tier</th>
                            <th>Total</th>
                            <th>Ultimo</th>
                            <th>Incremento</th>
                        </tr>
                    </thead>
                    <tbody>
                    {total.map((option) => (
                        <tr bgcolor={option.puesto}> 
                            
                            {/*<th ><button  id={option.id}  onClick ={handleDelete}>{option.id}</button></th>*/}
                            <th>{option.username}</th>
                            <th>{option.nombre}</th>
                            <th>{option.apellido}</th>
                            <th>{option.membresia}</th>
                            <th>Q{option.total}</th>
                            <th>+Q{option.ultimo}</th>
                            <th>{option.incremento}%</th>
                            
                        </tr>
                          ))}
                    </tbody>

                </table>
            <Footer />
        </div>
    )
}



export default ARecompensa
