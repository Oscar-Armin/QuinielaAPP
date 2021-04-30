import "base-64";
const md5 = require("md5");
const url_api = "http://localhost:3080";

export async function loginUsuario(usuario, password) {
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

/*
	Username         string `json:"username"`
	Nombre           string `json:"nombre"`
	Apellido         string `json:"apellido"`
	Fecha_nacimiento string `json:"nacimiento"`
	Correo   string `json:"correo"`
	Foto     string `json:"foto"`
	Password string `json:"password"`
*/

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
    
    
  ) {
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
        username: usuario,
        nombre: nombre,
        apellido:apellido,
        
        
        correo: correo,
        foto: foto_base64,
        password: password,
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