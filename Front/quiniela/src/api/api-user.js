import "base-64";
const md5 = require("md5");
const url_api = "http://localhost:3080";

export async function loginUsuario(usuario, password) {
    console.log(md5(password));
    return fetch(url_api + "/login", {
      

      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: usuario,
        password: md5(password),
      }),
    });
}
  
export async function consultUser(id) {
  return fetch(url_api + "/consultarUser", {
    
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      ID: id
      
    }),
  });
}

export var registrarUsuario = async function (
  usuario,
  nombre,
  apellido,
  nacimiento,
  correo,
  fotografias,
  password,
  
  
) {
  let file = fotografias.item(0);
  
  let foto_base64 = await pFileReader(file);
  foto_base64 = foto_base64.split(",")[1];
  
  return fetch(url_api + "/creaUser", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username: usuario,
      nombre: nombre,
      apellido:apellido,
      nacimiento: nacimiento,
      
      correo: correo,
      foto: foto_base64,
      password: md5(password),
    }),
  });
};

export var actualizarUsuario = async function (
  id,
  usuario,
  nombre,
  apellido,
  
  correo,
  fotografias,
  password,
  cambioFoto,
  antiguo
  
) {
  
  if (cambioFoto){
    let file = fotografias.item(0);
  
    let foto_base64 = await pFileReader(file);
    foto_base64 = foto_base64.split(",")[1];
    
    return fetch(url_api + "/actualizarUser", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id_usuario:id,
        username: usuario,
        nombre: nombre,
        apellido:apellido,
        
        
        correo: correo,
        foto: foto_base64,
        password: password,
        bandera:cambioFoto,
        anterior:antiguo
      }),
    });
  }else{
    return fetch(url_api + "/actualizarUser", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id_usuario:id,
        username: usuario,
        nombre: nombre,
        apellido:apellido,
        
        
        correo: correo,
        foto: fotografias,
        password: password,
        bandera:cambioFoto,
        anterior:antiguo
      }),
    });
  }

};

export async function forgetUsuario(usuario) {
  
  return fetch(url_api + "/consultID", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      usermail: usuario
      
    }),
  });
};

export var resetPassword = async function(idUser,pass) {
  console.log(idUser)
  return fetch(url_api + "/resetP", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      id_usuario: idUser,
      password: md5(pass)
      
    }),
  });
};

function pFileReader(file) {
  return new Promise((resolve, reject) => {
    var fr = new FileReader();
    fr.onload = () => {
      resolve(fr.result);
    };
    fr.readAsDataURL(file);
  });
}




  /*
    puntaje
    recompensa
    registro_membresia
    membresia
    prediccion
    usuario

  */