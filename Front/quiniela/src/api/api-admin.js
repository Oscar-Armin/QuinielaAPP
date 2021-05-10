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


function pFileReader(file) {
  return new Promise((resolve, reject) => {
    var fr = new FileReader();
    fr.onload = () => {
      resolve(fr.result);
    };
    fr.readAsDataURL(file);
  });
}

