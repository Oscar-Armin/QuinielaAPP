import "base-64";
const url_api = "http://localhost:3080";


export async function obtener_dinero(name) {
    
    return fetch(url_api + "/dinero_juego", {
      

      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: name
      }),
    });
}

export async function obtener_temporada() {
    
    return fetch(url_api + "/temporada_actual", {
      

      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        mensaje: "mensaje"
      }),
    });
}

export async function obtener_oros(name) {
    
    return fetch(url_api + "/numero_gold", {
      

      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: name
      }),
    });
}

export async function obtener_platas(name) {
    
    return fetch(url_api + "/numero_silver", {
      

      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: name
      }),
    });
}

export async function obtener_bronces(name) {
    
    return fetch(url_api + "/numero_bronze", {
      

      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: name
      }),
    });
}

export async function obtener__colores() {
    
    return fetch(url_api + "/colores", {
      

      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        mensaje: "mensaje"
      }),
    });
}

export var registro_deporte = async function (
  
  nombre,
  color,
  fotografias,
  
  
  
) {
  console.log(nombre,color)
  let file = fotografias.item(0);
  
  let foto_base64 = await pFileReader(file);
  foto_base64 = foto_base64.split(",")[1];
  
  return fetch(url_api + "/crearDeporte", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      
      nombre: nombre,
      color:color,
      
      foto: foto_base64
      
    }),
  });
};

export async function obtener_deportes() {
    
  return fetch(url_api + "/deportes", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      mensaje: "mensaje"
    }),
  });
}

export async function deleteDeporte(id) {
    
  return fetch(url_api + "/borraDeporte", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      id: parseInt(id)
    }),
  });
}

export async function obtener_ARecom() {
    
  return fetch(url_api + "/Arecompensas", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      mensaje: "mensaje"
    }),
  });
}

export async function obtener_Apuntos() {
    
  return fetch(url_api + "/Apuntaje", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      mensaje: "mensaje"
    }),
  });
}

export async function obtener_prediccion(user,temp) {
    
  return fetch(url_api + "/AUPred", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      id_user: user,
      id_temp:parseInt(temp)
    }),
  });
}

export async function obtener_temporadas() {
    
  return fetch(url_api + "/obtener_temporadas", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      mensaje: "mensaje"
    }),
  });
}

export async function obtener_nombre(ide) {
    
  return fetch(url_api + "/name_user", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      ID: parseInt(ide)
    }),
  });
}


export async function terminar_temporada() {
    
  return fetch(url_api + "/end_season", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      mensaje: "mensaje"
    }),
  });
}

export async function mensaje() {
    
  return fetch(url_api + "/mensaje", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      mensaje: "mensaje"
    }),
  });
}

export async function cambio_fase2() {
    
  return fetch(url_api + "/fase2", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      mensaje: "mensaje"
    }),
  });
}

export async function cambio_fase3() {
    
  return fetch(url_api + "/fase3", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      mensaje: "mensaje"
    }),
  });
}


export async function end_jornada() {
    
  return fetch(url_api + "/end_jornada", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      mensaje: "mensaje"
    }),
  });
}

export async function set_res(id,loc,vis) {
    
  return fetch(url_api + "/set_res", {
    

    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      id: parseInt(id),
      resultadol:parseInt(loc),
      resultadov:parseInt(vis)
    }),
  });
}



function pFileReader(file) {
  return new Promise((resolve, reject) => {
    var fr = new FileReader();
    fr.onload = () => {
      resolve(fr.result);
    };
    fr.readAsDataURL(file);
  });
}

