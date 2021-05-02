package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/godror/godror"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Types
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

var db *sql.DB

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome the my GO API!")
}

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

		resultado.Foto = "data:image/png;base64,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"
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

	json.NewEncoder(w).Encode(resultado)

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

func main() {
	var err error
	db, err = sql.Open("godror", "ADMIN/1234@172.17.0.2:1521/ORCL18")
	if err != nil {

		fmt.Println(err)
		return
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/creaUser", createUser).Methods("POST")
	router.HandleFunc("/login", loginUser).Methods("POST")
	//router.HandleFunc("/upload", uploadFile).Methods("POST")
	router.HandleFunc("/consultarUser", consultUser).Methods("POST")
	router.HandleFunc("/actualizarUser", updateUser).Methods("POST")
	router.HandleFunc("/consultID", forgetP).Methods("POST")
	router.HandleFunc("/resetP", resetP).Methods("POST")

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
