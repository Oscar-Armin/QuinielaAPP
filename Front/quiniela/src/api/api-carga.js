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

export async function carga_deporte(deporte) {
    
  return fetch(url_api + "/cargaDEP", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      
      nombre: deporte
      
      
    }),
  });
}

export async function carga_equipo(deporte) {
    
  return fetch(url_api + "/cargaEQ", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      
      nombre: deporte
      
      
    }),
  });
}

export async function carga_partido(equipo_local,equipo_visitante,puntos_local,puntos_visitante,fecha_partido,deporte,jornada,temporada) {
    
  return fetch(url_api + "/cargaPAR", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      
      equipo_local: equipo_local,
      equipo_visitante:equipo_visitante,
      puntos_local:puntos_local,
      puntos_visitante:puntos_visitante,
      fecha_partido:fecha_partido,
      deporte:deporte,
      jornada:jornada,
      temporada:temporada

      
    }),
  });
}

export async function carga_prediccion(username,equipo_local,equipo_visitante,deporte,fecha_partido,jornada,temporada,puntos_local,puntos_visitante,puntos) {
    
  return fetch(url_api + "/cargaPRE", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username:username,
      equipo_local: equipo_local,
      equipo_visitante:equipo_visitante,
      deporte:deporte,
      fecha_partido:fecha_partido,
      
      jornada:jornada,
      temporada:temporada,
      puntos_local:puntos_local,
      puntos_visitante:puntos_visitante,
      puntos:puntos

      
    }),
  });
}

export async function carga_puntos() {
    
  return fetch(url_api + "/carga_puntos", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      mensaje:"Funciono",

    }),
  });
}

export async function carga_recompensa() {
    
  return fetch(url_api + "/carga_recompensa", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      mensaje:"Funciono",

    }),
  });
}