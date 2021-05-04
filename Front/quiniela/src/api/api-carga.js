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