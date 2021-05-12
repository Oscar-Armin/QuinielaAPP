package main

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/quotedprintable"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/godror/godror"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Types MANEJO DE USUARIO
type task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type reset struct {
	Usermail string `json:"usermail"`
}

type pass struct {
	Id_usuario int    `json:"id_usuario"`
	Password   string `json:"password"`
}

type id_correo struct {
	ID     int    `json:"ID"`
	Correo string `json:"correo"`
}

type simpleID struct {
	ID int `json:"ID"`
}

type usuario struct {
	Username         string `json:"username"`
	Nombre           string `json:"nombre"`
	Apellido         string `json:"apellido"`
	Fecha_nacimiento string `json:"nacimiento"`
	Correo           string `json:"correo"`
	Foto             string `json:"foto"`
	Password         string `json:"password"`
}

type allUser struct {
	ID_user          int    `json:"id_usuario"`
	Username         string `json:"username"`
	Nombre           string `json:"nombre"`
	Apellido         string `json:"apellido"`
	Fecha_nacimiento string `json:"nacimiento"`
	Fecha_registro   string `json:"registro"`
	Correo           string `json:"correo"`
	Foto             string `json:"foto"`
	Password         string `json:"password"`
}

type actualizar struct {
	ID_user  int    `json:"id_usuario"`
	Username string `json:"username"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo   string `json:"correo"`
	Foto     string `json:"foto"`
	Password string `json:"password"`
	Bandera  bool   `json:"bandera"`
	Anterior string `json:"anterior"`
}

//structs carga de archivos
type insertUser struct {
	Username string `json:"username"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Password string `json:"password"`
}

type insertTMP struct {
	Name   string `json:"nombre"`
	Inicio string `json:"inicio"`
	Fin    string `json:"fin"`
}

type insertRM struct {
	Id_user   int    `json:"id_user"`
	Temporada string `json:"temporada"`
	Membresia string `json:"membresia"`
}

type ingresoJ struct {
	Njornada  int    `json:"njornada"`
	Temporada string `json:"temporada"`
}

type ingresoD struct {
	Nombre string `json:"nombre"`
}

type jsonPartido struct {
	Equipo_local     string `json:"equipo_local"`
	Equipo_visitante string `json:"equipo_visitante"`
	Puntos_local     int    `json:"puntos_local"`
	Puntos_visitante int    `json:"puntos_visitante"`
	Fecha_partido    string `json:"fecha_partido"`
	Deporte          string `json:"deporte"`
	Jornada          int    `json:"jornada"`
	Temporada        string `json:"temporada"`
}

type jsonPrediccion struct {
	Username         string `json:"username"`
	Equipo_local     string `json:"equipo_local"`
	Equipo_visitante string `json:"equipo_visitante"`
	Deporte          string `json:"deporte"`
	Fecha_partido    string `json:"fecha_partido"`

	Jornada          int    `json:"jornada"`
	Temporada        string `json:"temporada"`
	Puntos_local     int    `json:"puntos_local"`
	Puntos_visitante int    `json:"puntos_visitante"`
	Puntos           int    `json:"puntos"`
}

type jsonVacio struct {
	Mensaje string `json:"mensaje"`
}

var db *sql.DB

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome the my GO API!")
}

//MANEJO DE USUARIO
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser usuario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {

		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)

	dec, ero := base64.StdEncoding.DecodeString(newUser.Foto)
	if ero != nil {
		panic(ero)
	}
	f, err := os.Create("usuarios/" + newUser.Username + ".jpeg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	//currentTime.Format("02/01/2006")
	//rows, err := db.Query("select to_char(sysdate, 'HH24:MI:SS') from dual")
	//rows, err := db.Query("insert into usuario(username,nombre,apellido,fecha_nacimiento,fecha_registro,correo,foto_perfil,contrasena) values ()")
	currentTime := time.Now()
	rows, err := db.Query("INSERT INTO usuario (username,nombre,apellido,fecha_nacimiento,fecha_registro,correo,foto_perfil,contrasena) VALUES ('" + newUser.Username + "','" + newUser.Nombre + "','" + newUser.Apellido + "',TO_DATE('" + newUser.Fecha_nacimiento + "','dd/mm/yyyy'),TO_DATE('" + currentTime.Format("02/01/2006") + "','dd/mm/yyyy'),'" + newUser.Correo + "','" + "usuarios/" + newUser.Username + ".jpeg" + "','" + newUser.Password + "')")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	/*	for rows.Next() {

		var nombre string

		rows.Scan(&nombre)
		fmt.Println(nombre)
	}*/

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func resetP(w http.ResponseWriter, r *http.Request) {
	var newUser pass
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	fmt.Println(newUser)
	//currentTime.Format("02/01/2006")
	//rows, err := db.Query("select to_char(sysdate, 'HH24:MI:SS') from dual")
	//rows, err := db.Query("insert into usuario(username,nombre,apellido,fecha_nacimiento,fecha_registro,correo,foto_perfil,contrasena) values ()")

	rows, err := db.Query("update usuario set contrasena='" + newUser.Password + "' where id_usuario= " + strconv.Itoa(newUser.Id_usuario) + "")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	/*	for rows.Next() {

		var nombre string

		rows.Scan(&nombre)
		fmt.Println(nombre)
	}*/

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var newUser actualizar
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	fmt.Println(newUser.Bandera)
	fmt.Println(newUser.Anterior)

	if newUser.Bandera {
		if archivoExiste("usuarios/" + newUser.Username + ".jpeg") {
			//elimino
			err := os.Remove("usuarios/" + newUser.Username + ".jpeg")
			if err != nil {
				fmt.Printf("Error eliminando archivo: %v\n", err)
			}
		}
		dec, ero := base64.StdEncoding.DecodeString(newUser.Foto)
		if ero != nil {
			panic(ero)
		}
		f, err := os.Create("usuarios/" + newUser.Username + ".jpeg")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if _, err := f.Write(dec); err != nil {
			panic(err)
		}
		if err := f.Sync(); err != nil {
			panic(err)
		}
	} else {
		//Solo cambio de nombre a la imagen existente

		oldName := "usuarios/" + newUser.Anterior + ".jpeg"
		newName := "usuarios/" + newUser.Username + ".jpeg"
		fmt.Println(oldName)
		fmt.Println(newName)
		err := os.Rename(oldName, newName)
		if err != nil {
			log.Fatal(err)
		}
	}

	//currentTime.Format("02/01/2006")
	//rows, err := db.Query("select to_char(sysdate, 'HH24:MI:SS') from dual")
	//rows, err := db.Query("insert into usuario(username,nombre,apellido,fecha_nacimiento,fecha_registro,correo,foto_perfil,contrasena) values ()")
	fmt.Println(newUser.Username)
	rows, err := db.Query("update usuario set username='" + newUser.Username + "', nombre='" + newUser.Nombre + "' , apellido='" + newUser.Apellido + "', correo='" + newUser.Correo + "', foto_perfil =  'usuarios/" + newUser.Username + ".jpeg" + "', contrasena='" + newUser.Password + "' where id_usuario= " + strconv.Itoa(newUser.ID_user) + "")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	/*	for rows.Next() {

		var nombre string

		rows.Scan(&nombre)
		fmt.Println(nombre)
	}*/

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}

func consultUser(w http.ResponseWriter, r *http.Request) {
	var newUser simpleID
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)

	//fmt.Println(newUser)

	//currentTime.Format("02/01/2006")
	rows, err := db.Query("select * from usuario where id_usuario = '" + strconv.Itoa(newUser.ID) + "'")

	if err != nil {
		fmt.Println("Error running query")
		//fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var resultado allUser
	//fmt.Println(reflect.TypeOf(rows))
	//var resultado usuario
	for rows.Next() {

		rows.Scan(&resultado.ID_user, &resultado.Username, &resultado.Nombre, &resultado.Apellido, &resultado.Fecha_nacimiento, &resultado.Fecha_registro, &resultado.Correo, &resultado.Foto, &resultado.Password)

	}

	/**/ //convierto imagen a base64

	if archivoExiste(resultado.Foto) {

		bytes, err := ioutil.ReadFile(resultado.Foto)
		if err != nil {
			log.Fatal(err)
		}

		var base64Encoding string

		// Determine the content type of the image file
		mimeType := http.DetectContentType(bytes)

		// Prepend the appropriate URI scheme header depending
		// on the MIME type
		switch mimeType {
		case "image/jpeg":
			base64Encoding += "data:image/jpeg;base64,"
		case "image/png":
			base64Encoding += "data:image/jpeg;base64,"
		}

		// Append the base64 encoded output
		base64Encoding += toBase64(bytes)

		// Print the full base64 representation of the image
		//fmt.Println(base64Encoding)
		resultado.Foto = base64Encoding

	} else {

		resultado.Foto = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAgAAAAIACAYAAAD0eNT6AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAE3BJREFUeNrs3e11FEcWBuAebQArInATgaUIPERgHIGHCCwiACIARcAQAXIEjCNAGwGzESAn0LNdqAcwC/rsj+q6z3POnPHZH3vWV71z375VXV1VAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAd7RQAijPwcHiqP06vO9/T9PsNqoJAgAwbVNPDX3f2I+6//iX7rvuPkO6aD/n3T+n77/bz3b/acPC1l8JBADg/o1+2X5+6hr7cib/88+7QPCf7p/PBQMQAIDvN/x9s/+5+64L+1dMk4NNFwo2lhVAAICoDX9/R/9r930YsAwpBPzZBYJzVwUIAFBy03/cfn6vvqzfc2nbfs7azxthAAQA0PRjh4FTewdAAIC5Nf5l+/VH1/y5u/MuCKyVAgQAyLnxr9qvZ1V5m/imljYRnrafV20YuFAOEAAgh6afNvCddHf8hyoyuDQNeGF5AAQAmLL5n3R3/Br/+F6YCIAAAGM3/rS2/7Iy6p/ap6WBNgQ8VwoQAGDIxp8a/utqPqfyRbFtP08cMAS3/E1TArhR80/j/veaf5ZSMHvX/o1ednsyABMAcNdvGgCYAMDtmv9jd/2znQY8VwowAYC7NP+0ye9EJWYtTQF+86QACABwk8af1pDfuusvxrYLAd4xAN/+3ikBfG7+6bz+d5p/UerqcknA3xRMAODK5m8XebmeeK8ACACg+QsBIACA5q/5CwEgAIDmjxAAAgAU2fxT0/+g+Yf2yIFBhP4dVAKCNn93/rztpkBgAgBBAkB6zv+xSlBdnhNw7LAgTACg/Ob/XPPnK3V1efAThPMvJSBQ819Wly/2gX+EgEVrt/t0dDCEYQmAKM3fpj+uY1MgsX4XlYAgXmv+XHeNdEERBAAo5O4/rflb9+c6dft5pgxEYQmA0pu/0T+3ZSkAEwAowEvNnztcMyAAwIzv/tMhLyuV4JaO2mvnRBkonSUASg4A6bS/pUpwB+lgoIcOCMIEAObX/JeaP/eQlo1MATABgBkGgLTxr1YJTAHABIA4zX+l+WMKACYAxAsA1v4xBQATAII1/6XmjykAmAAQLwB41S+9TwGaZvdAGTABgHybf635M8QUoNtXAgIAZOoPJWAgvysBpbEEQEkTgI+VY38ZTtoMuFUGTAAgr+b/WPNnYCZMCACQoV+VgIHZX0JRLAFQwt1/uvP/qBKM4LhpdufKgAkAuDMjFpsBEQAgI78oAcIm3I4lAOafYu3+Z1yWATABgAya/1LzxxQABADisfufsVlyogiWAJj7BOB9+3WkEozsgTcEYgIA0zX/Q82fiSyVAAEA/AgTj+CJAAB+hAnIPgBmzx4A5pteDxbvTAGYStPs/H5iAgAmAAQMoK4/BACY4Me3rjz/jwAKAgB+fGFkPysBAgAIALgGQQCAEdiFjQAAAgABWf9n8muwO4wKBABw94XrEAQAGOaivXwCAHLgWkQAAD+6uBZBAIAhGbuSC48CIgDAiGy8wrUIAgAB/aQEZKJWAgQA8KOLaxEEAAAgX15nyfxS68Fipwpk5LhpdufKgAkAQCw2AiIAAAACAPR/wToFEBMAEAAISQAgNw6mQgAAAAQAAEAAAAAEAABAAAAABAD4kaUSkBmvBEYAAAjIOQAIAACAAAAACAAAgAAAAAgAcEMXSkBmtkqAAADDO1cCMvNfJUAAAAAEAABAAAAABAAAQAAAAAQA+IGtEpAZT6YgAMDQmmYnAJAbZ1MgAAAAAgAAIABAb6y54noEAYCArLmSjabZuR4RAAAAAQCGslUCMmH8jwAAI/L2NXJh/I8AAAAIADCkjRKQib+UAAEAABAAYEA2XuFaBAGAaDx3TUZciwgAMLKtEuA6BAEAP7wwOm+nRAAAAQDXIAgAMAKHASEAgABAQHZf4xoEAYCA7L5man8rAXO2UAJmm14PFjtVYEKPmma3UQZMAMAUANcfCAAwAmuwTKa9+3f9IQCAAIBrDwQAGItHAZnKVgkQAMBdGPH8RwkQAEAAwLUHAgCMpXsroJ3YTGGrBAgA4E6MeOHTdYcAABP7SwkY2UYJEABgelslwDUHAgDxGMUyNk8AIADA1KzFInTC3XgZEPNPsQeLd+3XUiUYKXT63cQEANyR4VoDAQCmYk0WAQAEAALaKAEj8dgpAgDkoml228qJgJgAgACAKQAM4MJTJwgAkB/7AHD3DwIAJgDQO+v/FMXzrJSTZg8WO1VgQI+aZidoIgBAhgHAgUAMxgFAFPebqQQUxIiWobjzRwAAP9IIlzB/RlqUlWjtA2AY1v8RACDzAGAfAL2z/k+Rv5dKQGGMaumbO38EAJiBMyWgZ38qASUy1qK8VHuw+Nh+HaoEPTl2BDAmADAPGyWgJ87/RwCAGTGypS+WlBAAwASAgGwqpVj2AFBmsj1YvG+/jlSCe3rQNLsLZcAEAEwBiONc80cAgPl5owTck70kFM0SAOWmW48Dcj8e/8MEAGbKDm7uaqv5IwDAfBnhIjzCD1gCoOyE6+2A3I23/2ECAO7kCOZC80cAgPmzDIDQCAIAfsxBaITEHgDKT7kHi7ft12OV4AbS+P+BMmACAO7oiMXECAEA/KgjLIIAALPVnee+VgmukQ7/ERYRAMCdHcFo/oRiEyBx0q53A3A1Z/9jAgDu8AjG2f8IAFCwUyXAtQGXLAEQK/EeLD60X7VK8I2HTbPbKgMmAOBOjzjONH8EACjfWgn4hidECMkSAPFSr6OB+cLRv5gAQCBvlIDOWgkwAYBYUwCbAUls/sMEAEwBCGaj+SMAQDyvlEAIVAIEAAjGC4LCSyf/+fsjAIA7QPztIRabAImdgA8W79uvI5UI50E3BQITAAjKyYDxrDV/MAEAjwTG47W/YAIAn1gPjmOj+YMJAOwnAIftV5oCHKpG8R61AWCjDGACAB4JjGOr+YMAAN+yGbB8L5QAvrAEAPs0fLB43X6tVKLYu/+HygACAHwvANTV5V4AyvPEyX8gAIApgLt/EACUAEwB3P2DAABCgCmAu38QAMAUAHf/IACAKQDu/kEAAFMA3P2DAACmALj7BwEAZh8AvCNgvn5rA8CZMsAVv3FKAN/XvSPAEcHzs9H8wQQA+pgCvG8/tWrMhjf+gQkA9DIF8BKZ+Vhr/mACAH1OAtIU4EglspbC2nEbALZKASYA0JenSpC9U80fTABgiCmAxwLz5bE/EABgsADgscB82fgHt/1NUwK4GRsCs3Wm+YMJAIwxCXjXfi1VIgs2/oEJAIzGhsB8vND8QQCAUbQN57yyFJCDdOLfK2WAu7EEAHdNz84GmJLRP5gAwGSeKMFkjP5BAIBpWAqYjNE/9MASANw3RVsKGJPRP5gAQDZ+6xoTw3uq+YMJAOQ0BVi1X69VYlDnbfM/VgYwAYBstI1p3X6tVWJQpiwgAECW0gFB58oACAAQawqQ7lA9GjicWgmgP/YAQN+p+mCxU4XBQpbfLDABgCybv7tUQACAgASAYQPWUhVAAIAcHSoBIABAPE4EHJYJAAgAkKWflWBQPykBCACQo1oJ1BfmwCM10Gei9gjg4DwKCCYAkFvzX6qCOoMAAPFoTOOw0RIEAMjKL0qgzjAX1tKgjyR9sEjP/39UiXHYBwAmAJCLx0owauBSbxAAIAu/KoF6w5wYo8H970br9uuDSowqvXr5YfcKZsAEACaxUoLRpT0XlgHABAAmnQB8rLwEaArbptk9VAYwAYApmv+J5j+Zuq3/ShnABADGbv6p8X8QAEwBwAQAYnmt+WcxBXiuDGACAGPd/a+6AEAejptmd64MIADAkM0/nUX/zt1/VrZdCPBYINz0t0wJQPMvQJ3+Lt2+DEAAgF6b/0rzz9pRFwK8LRAEAOil8aeNZm8rm/7mFAJOlAKuZg8A/Ljxp2afGskfGv8sbdrPi6bZbZQCBADQ+OM5az9P2yCwVQoQAEDjj2fdTQQEARAA4PPb/FLTX2n8YYLAG0sDCAAQt/Gnt8n9XnmrXFQpAJy2QeBMKRAAIMbd/qpr/LWKUF0eIvSm/bxykBACAJTV9Pfvjk9Nf6kiXGHdfv40FUAAgHk3fiN+7jsVWNs0iAAA82n6v3ZN34Y++pBeMnTafs4sESAAgKZPTGlp4E9hAAEANH2EAWEAAQA0fYQBYQABADR9hAEQAEDTRxgAAQA0fYQBEADQ9FUEYQAEADR9iGJdOX0QAQBNH8K62E8GhAEEADR9EAaEAQQAsm/6R9Xl2fsrTR96DQPr9vOmDQPnyoEAQG5NP93p1yoCg9p2k4FTLylCAGCKpl9XX16ve6QiMIk0DUhvLDwTBhAAGLLpH37V9JcqAlnZ7xdYKwUCAH01/v1mvpVqQPb2mwfTfoGNciAAcNumX7dff1TW9WHOtu3ntLJEgADANU1/P+JPjd+6PpRlPxXwSCECAJ8b/9FXd/se3YOy7R8p9BQBAoC7fXf7ENSmmwqslUIAoPzGX7dfz9ztA19Jk4D0OOHaVEAAoLzGv6o8vgdcb3/I0EYpBADm2/TTHf5J1/hrFQFuORV4YXlAAGBejT81e2N+oA9p02B6lNDygABAxo1/WX3ZzQ/Qt3U3FRAEBAAyafz73fxL1QBGsOmCwEYpBACmafyr6nLUX6sGIAggAGj8AGNKbyU8tWFQAEDjB2LaVp4cEADQ+IHQQeCp9w4IANy98S/br5eVo3qBedpU9ggIANy68ac7/qVqAIIAAkD5jb/uGv9KNYACrSvnCAgA/KPx74/sTc/yO7kPKN2rLghcKIUAELn5ryob/IB4LroQ8EopBIBojT9t7Esb/JaqAQSWzhB4an+AABCh8e/H/c9UA+CzdRcELAsIAEU2/3S3/7oy7gf4HssCAkCRd/2p8XtLH8D10rLAkzYInCvFwP1JCQZt/qnpf9D8AW4s7ZF63/5+PlcKEwB3/QAxbbtpwEYpBIA5NP+UXt9W1voB+pL2BpgICABZN/9Vdfl4nwN9APqV9gT85iTBHnuWEvTW/FPjf635AwxivzfA0qoJQFbNPzX+lUoAjOKpxwUFgKkbv81+ANNYtyHgiTIIAFM1/3fV5VgKgAlCQOUEQQFgggDwXvMHmFzaHPhICLhDH1OCOzX/15o/QBbSb/G7biqLCcDgzX+lEgB5TQKaZnesDCYAQzX/E80fIM9JQHeDhglA780/7fR/qxIAWfN0gADQa/P/tMZUOeQHYA6cEyAA9NL8Pe4HMD/p2OAzZRAA7hMAbPoDmJ/0WOCxdwdc0d+U4Mrmv9L8AWYpTW/fejzQBOAuzb9uv95X1v0B5symQAHg1gEgrfsvVQJg9uwHEABu3PzT8/4vVQKgCGk/wEPHBQsA1zX/ujL6ByjNpg0Aj5Thq36nBP/npeYPUJxlt7EbE4Dv3v077Q+gXJYCTACuvPsHoEyHfucFgO/d/T9vv2qVACjaqv29XyqDJYB980+p8ENl7R8gAq8ONgH47JnmDxDGkQ2BJgD7x/4++P8DQCjb6vJdAWE3BJoAXN79AxBLuvk7MQFw9w9APKEfC4w+AXD3DxDXYeQpQNgJgLt/ACJPASJPAFauewBTgKhTgJATAM/9A/D1FKBpdg9MAOLc/Wv+AHyaAkQ8FyDqBCDd/deueQA626bZPTQBKLv5P9b8AfhGHe0dARGXAH53nQMQvT+EWgLw6B8A13gQ5ZHAaBOAx65tAK6wMgEocwJg8x8AVwmzGTDMBKBt/keaPwDXqLt+IQAUxOY/APSLTpglAON/AG4oxDJAiAmA8T8AtxBiGSDKEoDxPwD6xldCLAEY/wNwS8UvAxQfABz+A8AdPWxDwLbY/hjgD+jwHwD0j4AB4BfXMAD6xz9FWALYuYYBuIum2RXbJ4ueAER7tSMA+ogAcEkAAEAfCRgArP8DoI98R9F7AKz/A3Bfpe4DKHYCYP0fAP0kYACorP8DoJ+EDAA/u2YB0E/iBYAj1ywA+sn3lbmx4WBx2H59dM0C0JMHTbO7MAGQ1gAwBRAAMrR0rQKgr8QLADYAAqCvBAwAtWsVAH3lx0rdBOgEQAB6VdqJgMVNAJwACMBA/aWojYAlLgEcukwBGEAtAOTNI4AA6C8BA4AnAADQXwIGAEsAAOgv1yjuKQBPAAAwlJKeBBAAACBgAChqCcAjgADoMwEDAAAQMwCYAACgz5gAAAARAoAzAADQZwIGAGcAAKDPBAwAAMANlPVqQ2cAADCwUs4CEAAAQAAQAAAgQgAoZg+AUwAB0G8CBgAAQAAAAAQAAKDUALD05wRAvzEBAAAEAABAAAAAAQAAEAAAAAEAABAAAAABIBv/9ucEgHgB4MifEwDiBYBzf04AiBcA/vbnBIB4AQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABjZ/wQYAGVy88shRG24AAAAAElFTkSuQmCC"
	}

	// Print the full base64 representation of the image

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)

}

func archivoExiste(ruta string) bool {
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		return false
	}
	return true
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var newUser login
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	fmt.Println(newUser.Username)
	fmt.Println(newUser.Username)
	var resultado simpleID
	//currentTime.Format("02/01/2006")
	err = db.QueryRow("select id_usuario from usuario where username = '" + newUser.Username + "' and contrasena = '" + newUser.Password + "'").Scan(&resultado.ID)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resultado)

}

func forgetP(w http.ResponseWriter, r *http.Request) {
	var newUser reset
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)

	//currentTime.Format("02/01/2006")
	//rows, err := db.Query("select to_char(sysdate, 'HH24:MI:SS') from dual")
	//rows, err := db.Query("insert into usuario(username,nombre,apellido,fecha_nacimiento,fecha_registro,correo,foto_perfil,contrasena) values ()")

	var resultado id_correo
	//currentTime.Format("02/01/2006")
	err = db.QueryRow("select id_usuario,correo from usuario where username = '"+newUser.Usermail+"' or correo = '"+newUser.Usermail+"'").Scan(&resultado.ID, &resultado.Correo)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	send(resultado.Correo, resultado.ID, newUser.Usermail)
	json.NewEncoder(w).Encode(resultado)
	//send("hola")

}
func send(enviar string, identificador int, cuenta string) {
	from := "armin99.cr@gmail.com"
	pass := "armincr99"
	to := enviar
	fmt.Println(enviar)
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there " + cuenta + "\n\n" +
		"	Para recuperar contraseña darl al siguiente link \n\n	http://localhost:3000/reset/" + strconv.Itoa(identificador)

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		//log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, ")
}

func email(enviar string, identificador int, cuenta string) {
	// sender data
	from := "armin99.cr@gmail.com" //ex: "John.Doe@gmail.com"
	password := "armincr99"        // ex: "ieiemcjdkejspqz"
	// receiver address
	toEmail := os.Getenv(enviar) // ex: "Jane.Smith@yahoo.com"
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Subject: Our Golang Email\n"
	body := "Cuenta " //+ cuenta + ":\n Para recuperar cuenta entrar al siguiente enlace \n http://localhost:3000/reset/" + strconv.Itoa(identificador)
	message := []byte(subject + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// send mail
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}

const (
	/**
		Gmail SMTP Server
	**/
	SMTPServer = "smtp.gmail.com"
)

type Sender struct {
	User     string
	Password string
}

func NewSender(Username, Password string) Sender {

	return Sender{Username, Password}
}

func (sender Sender) SendMail(Dest []string, Subject, bodyMessage string) {

	msg := "From: " + sender.User + "\n" +
		"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTPServer+":587",
		smtp.PlainAuth("", sender.User, sender.Password, SMTPServer),
		sender.User, Dest, []byte(msg))

	if err != nil {

		fmt.Printf("smtp error: %s", err)
		return
	}

	fmt.Println("Mail sent successfully!")
}

func (sender Sender) WriteEmail(dest []string, contentType, subject, bodyMessage string) string {

	header := make(map[string]string)
	header["From"] = sender.User

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func (sender *Sender) WriteHTMLEmail(dest []string, subject, bodyMessage string) string {

	return sender.WriteEmail(dest, "text/html", subject, bodyMessage)
}

func (sender *Sender) WritePlainEmail(dest []string, subject, bodyMessage string) string {

	return sender.WriteEmail(dest, "text/plain", subject, bodyMessage)
}

type Dest struct {
	Name string
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(tasks)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

//funciones carga
func ingresoUser(w http.ResponseWriter, r *http.Request) {

	var newUser insertUser
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {

		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)

	//currentTime.Format("02/01/2006")
	//rows, err := db.Query("select to_char(sysdate, 'HH24:MI:SS') from dual")
	//rows, err := db.Query("insert into usuario(username,nombre,apellido,fecha_nacimiento,fecha_registro,correo,foto_perfil,contrasena) values ()")
	currentTime := time.Now()
	rows, err := db.Query("INSERT INTO usuario (username,nombre,apellido,fecha_registro,contrasena) VALUES ('" + newUser.Username + "','" + newUser.Nombre + "','" + newUser.Apellido + "',TO_DATE('" + currentTime.Format("02/01/2006") + "','dd/mm/yyyy'),'" + newUser.Password + "')")
	if err != nil {
		//fmt.Println("Error running query")
		//fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func ingresoTMP(w http.ResponseWriter, r *http.Request) {
	var newUser insertTMP
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	//fmt.Println(newUser)

	rows, err := db.Query("INSERT INTO temporada (nombre) VALUES ('" + newUser.Name + "')")
	if err != nil {
		//fmt.Println("Error running query")
		//fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func ingresoRMEM(w http.ResponseWriter, r *http.Request) {
	var newUser insertRM
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	//fmt.Println(newUser)

	rows, err := db.Query("insert into registro_membresia  (id_usuario,id_temporada,id_membresia) VALUES (" + strconv.Itoa(newUser.Id_user) + ",(select id_temporada from temporada where nombre= '" + newUser.Temporada + "'),(select id_membresia from membresia where nombre= '" + newUser.Membresia + "'))")
	if err != nil {
		//fmt.Println("Error running query")
		//fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func ingresoJOR(w http.ResponseWriter, r *http.Request) {
	var newUser ingresoJ
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	//fmt.Println(newUser)

	rows, err := db.Query("BEGIN insert_jornada(" + strconv.Itoa(newUser.Njornada) + ",'" + newUser.Temporada + "'); END;")
	if err != nil {
		//fmt.Println("Error running query")
		//fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func ingresoDEP(w http.ResponseWriter, r *http.Request) {
	var newUser ingresoD
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	fmt.Println(newUser)

	rows, err := db.Query("BEGIN insert_deporte('" + newUser.Nombre + "'); END;")
	if err != nil {
		//fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func ingresoEQ(w http.ResponseWriter, r *http.Request) {
	var newUser ingresoD
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	//fmt.Println(newUser)

	rows, err := db.Query("insert into equipo (nombre) VALUES('" + newUser.Nombre + "')")
	if err != nil {
		//fmt.Println("Error running query")
		//fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func ingresoPAR(w http.ResponseWriter, r *http.Request) {
	var newUser jsonPartido
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	//fmt.Println(newUser)
	//strconv.Itoa(newUser.Njornada)
	rows, err := db.Query("BEGIN insert_partido('" + newUser.Equipo_local + "','" + newUser.Equipo_visitante + "'," + strconv.Itoa(newUser.Puntos_local) + "," + strconv.Itoa(newUser.Puntos_visitante) + ",'" + newUser.Fecha_partido + "','" + newUser.Deporte + "'," + strconv.Itoa(newUser.Jornada) + ",'" + newUser.Temporada + "'); END;")
	if err != nil {
		//fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func ingresoPRE(w http.ResponseWriter, r *http.Request) {
	var newUser jsonPrediccion
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	fmt.Println(newUser)
	//strconv.Itoa(newUser.Njornada)
	rows, err := db.Query("BEGIN insert_prediccion('" + newUser.Username + "','" + newUser.Equipo_local + "','" + newUser.Equipo_visitante + "','" + newUser.Deporte + "','" + newUser.Fecha_partido + "'," + strconv.Itoa(newUser.Jornada) + ",'" + newUser.Temporada + "'," + strconv.Itoa(newUser.Puntos_local) + "," + strconv.Itoa(newUser.Puntos_visitante) + "," + strconv.Itoa(newUser.Puntos) + "); END;")
	if err != nil {
		//fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func cargar_puntos(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)
	fmt.Println(newUser)
	//strconv.Itoa(newUser.Njornada)
	rows, err := db.Query("BEGIN proceso_usuario; END;")
	if err != nil {
		//fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func cargar_recompensas(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)

	//strconv.Itoa(newUser.Njornada)
	rows, err := db.Query("BEGIN calculo_temporada; END;")
	if err != nil {
		//fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

//modulo admin
type jsonTemp struct {
	Name string `json:"name"`
}
type jsonDinero struct {
	Cantidad int `json:"cantidad"`
}

type jsonTempo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type jsonColor struct {
	ID    int    `json:"id"`
	Color string `json:"color"`
}

type jsonDeporte struct {
	Nombre string `json:"nombre"`
	Color  string `json:"color"`
	Foto   string `json:"foto"`
}
type jsonRecoAdmin struct {
	Username   string  `json:"username"`
	Nombre     string  `json:"nombre"`
	Apellido   string  `json:"apellido"`
	Membresia  string  `json:"membresia"`
	Total      float32 `json:"total"`
	Ultimo     float32 `json:"ultimo"`
	Incremento float32 `json:"incremento"`
	Puesto     string  `json:"puesto"`
}

type jsonPuntAdmin struct {
	Puesto   int    `json:"puesto"`
	Id       int    `json:"id"`
	Username string `json:"username"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	P0       int    `json:"p0"`
	P3       int    `json:"p3"`
	P5       int    `json:"p5"`
	P10      int    `json:"p10"`
	Puntos   int    `json:"puntos"`
}

type jsonUPred struct {
	Deporte    string `json:"deporte"`
	Local      string `json:"local"`
	Visitante  string `json:"visitante"`
	Plocal     int    `json:"plocal"`
	Pvisitante int    `json:"pvisitante"`
	Rlocal     int    `json:"rlocal"`
	Rvisitante int    `json:"rvisitante"`
	Puntos     int    `json:"puntos"`
	Fecha      string `json:"fecha"`
}

type json2ID struct {
	Id_user int `json:"id_user"`
	Id_temp int `json:"id_temp"`
}
type jsonCalendario struct {
	Id              int     `json:"id"`
	Title           string  `json:"title"`
	Start           string  `json:"start"`
	BackgroundColor string  `json:"backgroundColor"`
	Extended        jsonAdd `json:"extendedProps"`
}

type jsonAdd struct {
	ResultadoL int    `json:"resultadoL"`
	ResultadoV int    `json:"resultadoV"`
	Prediccion int    `json:"prediccion"`
	Deporte    string `json:"deporte"`
}

type jsonResultados struct {
	ID         int `json:"id"`
	ResultadoL int `json:"resultadol"`
	ResultadoV int `json:"Resultadov"`
}

/*id: "", //id_partido
title: 'J1vsJ2', //local vs visitante
start: '2021-05-11T11:25:00', //fecha y hora

backgroundColor:"red",// color deporte
extendedProps: {

	resultadoL:0,
	resultadoV:0,
	prediccion: 0 //si o no
},
description: 'Lecture'*/

type idSport struct {
	ID int `json:"id"`
}

type allColors []jsonColor

func getDinero(w http.ResponseWriter, r *http.Request) {
	var temp jsonTemp
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &temp)

	var resultado jsonDinero
	//currentTime.Format("02/01/2006")
	err = db.QueryRow("select sum(temporal) from( select (select precio from membresia where membresia.id_membresia = registro_membresia.id_membresia) as temporal from registro_membresia where id_temporada = (select id_temporada from temporada where nombre = '" + temp.Name + "'))").Scan(&resultado.Cantidad)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}

func getGold(w http.ResponseWriter, r *http.Request) {
	var temp jsonTemp
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &temp)

	var resultado jsonDinero
	//currentTime.Format("02/01/2006")
	err = db.QueryRow("select count(*) from registro_membresia where id_temporada =  (select id_temporada from temporada where nombre = '" + temp.Name + "') and id_membresia = (select id_membresia from membresia where nombre = 'gold')").Scan(&resultado.Cantidad)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}

func getSilver(w http.ResponseWriter, r *http.Request) {
	var temp jsonTemp
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &temp)

	var resultado jsonDinero
	//currentTime.Format("02/01/2006")
	err = db.QueryRow("select count(*) from registro_membresia where id_temporada =  (select id_temporada from temporada where nombre = '" + temp.Name + "') and id_membresia = (select id_membresia from membresia where nombre = 'silver')").Scan(&resultado.Cantidad)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}

func getBronze(w http.ResponseWriter, r *http.Request) {
	var temp jsonTemp
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &temp)

	var resultado jsonDinero
	//currentTime.Format("02/01/2006")
	err = db.QueryRow("select count(*) from registro_membresia where id_temporada =  (select id_temporada from temporada where nombre = '" + temp.Name + "') and id_membresia = (select id_membresia from membresia where nombre = 'bronze')").Scan(&resultado.Cantidad)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}

func getTA(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)
	fmt.Println(newUser)
	var resultado jsonTempo
	//currentTime.Format("02/01/2006")
	err = db.QueryRow("select nombre from temporada where actual = 1").Scan(&resultado.Name)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(resultado)
	json.NewEncoder(w).Encode(resultado)

}

func getColores(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)
	fmt.Println(newUser)

	//currentTime.Format("02/01/2006")

	rows, err := db.Query("select id_color,color from colores")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var resultado []jsonColor
	for rows.Next() {

		var color jsonColor
		rows.Scan(&color.ID, &color.Color)
		resultado = append(resultado, color)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resultado)

}

func createSport(w http.ResponseWriter, r *http.Request) {
	var temp jsonDeporte
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {

		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &temp)

	dec, ero := base64.StdEncoding.DecodeString(temp.Foto)
	if ero != nil {
		panic(ero)
	}
	f, err := os.Create("deportes/" + temp.Nombre + ".jpeg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	//currentTime.Format("02/01/2006")
	//rows, err := db.Query("select to_char(sysdate, 'HH24:MI:SS') from dual")
	//rows, err := db.Query("insert into usuario(username,nombre,apellido,fecha_nacimiento,fecha_registro,correo,foto_perfil,contrasena) values ()")

	rows, err := db.Query("BEGIN ingreso_deporte('" + temp.Nombre + "','" + temp.Color + "','deportes/" + temp.Nombre + ".jpeg'); END;")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	/*	for rows.Next() {

		var nombre string

		rows.Scan(&nombre)
		fmt.Println(nombre)
	}*/

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(temp)
}

func getDeportes(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)
	fmt.Println(newUser)

	//currentTime.Format("02/01/2006")

	rows, err := db.Query("select id_deporte,nombre from deporte")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var resultado []jsonTempo
	for rows.Next() {

		var color jsonTempo
		rows.Scan(&color.ID, &color.Name)
		resultado = append(resultado, color)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resultado)
}

func deleteDeporte(w http.ResponseWriter, r *http.Request) {
	var newUser idSport
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)
	fmt.Println(newUser)

	//currentTime.Format("02/01/2006")
	fmt.Println(newUser)

	rows, err := db.Query("delete from deporte where id_deporte =" + strconv.Itoa(newUser.ID))
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var resultado []jsonTempo
	for rows.Next() {

		var color jsonTempo
		rows.Scan(&color.ID, &color.Name)
		resultado = append(resultado, color)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resultado)
}

func AgetReco(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)

	//currentTime.Format("02/01/2006")
	var consulta = "select " +
		"(select username from usuario where id_usuario = recompensa.id_usuario) as Username," +
		"(select nombre from usuario where id_usuario = recompensa.id_usuario) as Nombre," +
		"(select apellido from usuario where id_usuario = recompensa.id_usuario) as Apellido, " +
		"(select tiers from  (select (select nombre from membresia where id_membresia = registro_membresia.id_membresia) as tiers,(select nombre from temporada where id_temporada = registro_membresia.id_temporada)  as nt from registro_membresia where id_usuario=recompensa.id_usuario order by nt desc) where rownum = 1) as Membresia," +
		"getTotal(recompensa.id_usuario) as Total, " +
		"getUltimo(recompensa.id_usuario) as ultimo," +
		"((getUltimo(recompensa.id_usuario)*100 )/getTotal(recompensa.id_usuario)) as incremento," +
		"getPuesto(getTotal(recompensa.id_usuario)) as Puesto " +
		"from recompensa group by id_usuario order by Total DESC"

	rows, err := db.Query(consulta)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var resultado []jsonRecoAdmin
	for rows.Next() {

		var j jsonRecoAdmin
		rows.Scan(&j.Username, &j.Nombre, &j.Apellido, &j.Membresia, &j.Total, &j.Ultimo, &j.Incremento, &j.Puesto)

		resultado = append(resultado, j)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resultado)
}

func Agetpuntaje(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)

	//currentTime.Format("02/01/2006")
	var consulta = "select " +
		"puesto_puntaje(puntos) as puesto," +
		"id_usuario," +
		"(select username from usuario where id_usuario = puntaje.id_usuario) as Username," +
		"(select nombre from usuario where id_usuario = puntaje.id_usuario) as Nombre," +
		"(select apellido from usuario where id_usuario = puntaje.id_usuario) as Apellido," +
		"p0,p3,p5,p10 ,puntos from puntaje where puesto_puntaje(puntos)!=0 and id_temporada = (select id_temporada  from temporada where actual = 2) order by puntos desc "

	rows, err := db.Query(consulta)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var resultado []jsonPuntAdmin
	for rows.Next() {

		var j jsonPuntAdmin
		rows.Scan(&j.Puesto, &j.Id, &j.Username, &j.Nombre, &j.Apellido, &j.P0, &j.P3, &j.P5, &j.P10, &j.Puntos)

		resultado = append(resultado, j)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resultado)
}

func AUPred(w http.ResponseWriter, r *http.Request) {
	var newUser json2ID
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)
	fmt.Println(newUser)
	//currentTime.Format("02/01/2006")
	var consulta = "select" +
		"(select (select nombre from deporte where id_deporte = partido.id_deporte) from partido where id_partido = prediccion.id_partido) as deporte," +
		"(select (select nombre from equipo where id_equipo = partido.equipo_local) from partido where id_partido = prediccion.id_partido) as E_local," +
		"(select (select nombre from equipo where id_equipo = partido.equipo_visitante) from partido where id_partido = prediccion.id_partido) as Visitante," +
		"prediccion.puntos_local as pre_local," +
		"prediccion.puntos_visitante as pre_visitante," +
		"(select puntos_local from partido where id_partido = prediccion.id_partido) as res_local," +
		"(select puntos_visitante from partido where id_partido = prediccion.id_partido) as res_visitante," +
		"puntos," +
		"(select fecha_partido from partido where id_partido = prediccion.id_partido) as Fecha " +
		" from prediccion where id_usuario = " + strconv.Itoa(newUser.Id_user) + " and id_partido in (select id_partido  from partido where id_jornada in (select id_jornada from jornada where id_temporada = " + strconv.Itoa(newUser.Id_temp) + "))"

	rows, err := db.Query(consulta)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var resultado []jsonUPred
	for rows.Next() {

		var j jsonUPred
		rows.Scan(&j.Deporte, &j.Local, &j.Visitante, &j.Plocal, &j.Pvisitante, &j.Rlocal, &j.Rvisitante, &j.Puntos, &j.Fecha)

		resultado = append(resultado, j)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(resultado)
	json.NewEncoder(w).Encode(resultado)

}

func getName(w http.ResponseWriter, r *http.Request) {
	var newUser simpleID
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)

	//fmt.Println(newUser)

	//currentTime.Format("02/01/2006")
	rows, err := db.Query("select username from usuario where id_usuario = '" + strconv.Itoa(newUser.ID) + "'")

	if err != nil {
		fmt.Println("Error running query")
		//fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var resultado ingresoD
	//fmt.Println(reflect.TypeOf(rows))
	//var resultado usuario
	for rows.Next() {

		rows.Scan(&resultado.Nombre)

	}

	/**/ //convierto imagen a base64

	// Print the full base64 representation of the image

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}

func getTemps(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)

	//currentTime.Format("02/01/2006")
	var consulta = "select id_temporada,nombre from temporada"

	rows, err := db.Query(consulta)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var resultado []id_correo
	for rows.Next() {

		var j id_correo
		rows.Scan(&j.ID, &j.Correo)

		resultado = append(resultado, j)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resultado)

}
func endTemp(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)

	//currentTime.Format("02/01/2006")
	var consulta = "BEGIN terminar_temporada; END;"

	rows, err := db.Query(consulta)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

func get_mensaje(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)

	//currentTime.Format("02/01/2006")

	var resultado jsonVacio
	err = db.QueryRow("select * from mensaje").Scan(&resultado.Mensaje)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resultado)
}

func get_deportes(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)

	//currentTime.Format("02/01/2006")
	var consulta = "select " +
		"id_partido," +
		"(select nombre from equipo where id_equipo=equipo_local) as E_local," +
		"(select nombre from equipo where id_equipo=equipo_visitante)as E_visitante," +
		"partido.fecha_partido as Fecha," +

		"(select color from deporte where id_deporte=partido.id_deporte)as color," +

		"(select nombre from deporte where id_deporte=partido.id_deporte)as deporte," +
		"partido.puntos_local as resultadoL," +
		"partido.puntos_visitante as resultadoV," +
		"tiene_predic(id_partido) as prediccion " +

		" from partido " +
		" where id_jornada in (select id_jornada from jornada where id_temporada =(select id_temporada from temporada where actual =1) order by jornada.numero_jornada desc FETCH NEXT 1 ROWS ONLY   )"

	rows, err := db.Query(consulta)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var resultado []jsonCalendario
	for rows.Next() {
		var local string
		var visitante string
		var j jsonCalendario
		rows.Scan(&j.Id, &local, &visitante, &j.Start, &j.BackgroundColor, &j.Extended.Deporte, &j.Extended.ResultadoL, &j.Extended.ResultadoV, &j.Extended.Prediccion)
		j.Title = local + " vs " + visitante

		resultado = append(resultado, j)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resultado)
}
func get_estadoUJ(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)

	//currentTime.Format("02/01/2006")
	var resultado simpleID
	err = db.QueryRow("select id_fase from jornada where id_temporada = (select id_temporada from temporada where actual =1) order by numero_jornada desc FETCH NEXT 1 ROWS ONLY  ").Scan(&resultado.ID)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resultado)
}

func set_fase2(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)
	var consulta = "update jornada set id_fase = 2 where id_jornada =" +
		"(select id_jornada from jornada where id_temporada = (select id_temporada from temporada where actual = 1) order by numero_jornada desc  FETCH NEXT 1 ROWS ONLY)"
	//currentTime.Format("02/01/2006")
	rows, err := db.Query(consulta)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)

}
func set_fase3(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)
	var consulta = "update jornada set id_fase = 3 where id_jornada =" +
		"(select id_jornada from jornada where id_temporada = (select id_temporada from temporada where actual = 1) order by numero_jornada desc  FETCH NEXT 1 ROWS ONLY)"
	//currentTime.Format("02/01/2006")
	rows, err := db.Query(consulta)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)

}

func end_jornada(w http.ResponseWriter, r *http.Request) {
	var newUser jsonVacio
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)
	var consulta = "BEGIN terminar_jornada; END;"
	//currentTime.Format("02/01/2006")
	rows, err := db.Query(consulta)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)

}

func set_res(w http.ResponseWriter, r *http.Request) {
	var newUser jsonResultados
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization") // You can add more headers here if needed
	} else {
		// Your code goes here
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid  Data")
	}

	json.Unmarshal(reqBody, &newUser)
	fmt.Println(newUser)
	var consulta = "Update partido set puntos_local =" + strconv.Itoa(newUser.ResultadoL) + " , puntos_visitante = " + strconv.Itoa(newUser.ResultadoV) + " where id_partido = " + strconv.Itoa(newUser.ID)
	//currentTime.Format("02/01/2006")
	fmt.Println(consulta)
	rows, err := db.Query(consulta)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)

}

func main() {
	var err error
	db, err = sql.Open("godror", "ADMIN/1234@172.17.0.2:1521/ORCL18")
	if err != nil {

		fmt.Println(err)
		return
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	//MANEJO DE USUARIO
	router.HandleFunc("/creaUser", createUser).Methods("POST")
	router.HandleFunc("/login", loginUser).Methods("POST")

	router.HandleFunc("/consultarUser", consultUser).Methods("POST")
	router.HandleFunc("/actualizarUser", updateUser).Methods("POST")
	router.HandleFunc("/consultID", forgetP).Methods("POST")
	router.HandleFunc("/resetP", resetP).Methods("POST")
	//CARGA DE ARCHIVO
	router.HandleFunc("/cargaUser", ingresoUser).Methods("POST")
	router.HandleFunc("/cargaTMP", ingresoTMP).Methods("POST")
	router.HandleFunc("/cargaRMEM", ingresoRMEM).Methods("POST")
	router.HandleFunc("/cargaJOR", ingresoJOR).Methods("POST")
	router.HandleFunc("/cargaDEP", ingresoDEP).Methods("POST")
	router.HandleFunc("/cargaEQ", ingresoEQ).Methods("POST")
	router.HandleFunc("/cargaPAR", ingresoPAR).Methods("POST")
	router.HandleFunc("/cargaPRE", ingresoPRE).Methods("POST")
	router.HandleFunc("/carga_puntos", cargar_puntos).Methods("POST")
	router.HandleFunc("/carga_recompensa", cargar_recompensas).Methods("POST")
	//modulo de administrador
	router.HandleFunc("/dinero_juego", getDinero).Methods("POST")
	router.HandleFunc("/numero_gold", getGold).Methods("POST")
	router.HandleFunc("/numero_silver", getSilver).Methods("POST")
	router.HandleFunc("/numero_bronze", getBronze).Methods("POST")
	router.HandleFunc("/temporada_actual", getTA).Methods("POST")
	router.HandleFunc("/colores", getColores).Methods("POST")
	router.HandleFunc("/crearDeporte", createSport).Methods("POST")

	router.HandleFunc("/deportes", getDeportes).Methods("POST")
	router.HandleFunc("/borraDeporte", deleteDeporte).Methods("POST")
	router.HandleFunc("/Arecompensas", AgetReco).Methods("POST")
	router.HandleFunc("/Apuntaje", Agetpuntaje).Methods("POST")

	router.HandleFunc("/AUPred", AUPred).Methods("POST")

	router.HandleFunc("/name_user", getName).Methods("POST")
	router.HandleFunc("/obtener_temporadas", getTemps).Methods("POST")
	router.HandleFunc("/end_season", endTemp).Methods("POST")
	router.HandleFunc("/mensaje", get_mensaje).Methods("POST")
	router.HandleFunc("/partido_ja", get_deportes).Methods("POST")
	router.HandleFunc("/estado_fase", get_estadoUJ).Methods("POST")
	router.HandleFunc("/fase2", set_fase2).Methods("POST")
	router.HandleFunc("/fase3", set_fase3).Methods("POST")
	router.HandleFunc("/end_jornada", end_jornada).Methods("POST")
	router.HandleFunc("/set_res", set_res).Methods("POST")

	fmt.Println("En puerto 3080")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":3080", handler))

}

func conection() {

}

/*
puntaje
recompensa
registo_membresia
membresia.
prediccion
usuario





*/
