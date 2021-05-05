import "base-64";
const md5 = require("md5");
const url_api = "http://localhost:3080";

export async function cargar_usuario(usuario,nombre,apellido, password) {
    
    return fetch(url_api + "/cargaUser", {
      

      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: usuario,
        nombre: nombre,
        apellido:apellido,
        password: md5(password),
      }),
    });
}

export async function cargar_temporadas(nombre,inicio,fin) {
    
  return fetch(url_api + "/cargaTMP", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      
      nombre: nombre,
      inicio:inicio,
      fin:fin
    }),
  });
}

export async function carga_rmembresia(id_user,temporada,membresia) {
    
  return fetch(url_api + "/cargaRMEM", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      
      id_user: id_user,
      temporada:temporada,
      membresia:membresia
    }),
  });
}

export async function carga_jornada(id_jornada,temporada) {
    
  return fetch(url_api + "/cargaJOR", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      
      njornada: id_jornada,
      temporada:temporada
      
    }),
  });
}