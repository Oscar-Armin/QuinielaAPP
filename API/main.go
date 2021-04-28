package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	fmt.Println(newUser)

	//currentTime.Format("02/01/2006")
	//rows, err := db.Query("select to_char(sysdate, 'HH24:MI:SS') from dual")
	//rows, err := db.Query("insert into usuario(username,nombre,apellido,fecha_nacimiento,fecha_registro,correo,foto_perfil,contrasena) values ()")
	currentTime := time.Now()
	rows, err := db.Query("INSERT INTO usuario (username,nombre,apellido,fecha_nacimiento,fecha_registro,correo,foto_perfil,contrasena) VALUES ('" + newUser.Username + "','" + newUser.Nombre + "','" + newUser.Apellido + "',TO_DATE('" + newUser.Fecha_nacimiento + "','dd/mm/yyyy'),TO_DATE('" + currentTime.Format("02/01/2006") + "','dd/mm/yyyy'),'" + newUser.Correo + "','" + newUser.Foto + "','" + newUser.Password + "')")
	if err != nil {
		fmt.Println("Error running query")
		//fmt.Println(err)
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
		fmt.Println("info")

		rows.Scan(&resultado.ID_user, &resultado.Nombre, &resultado.Username, &resultado.Apellido, &resultado.Fecha_nacimiento, &resultado.Fecha_registro, &resultado.Correo, &resultado.Foto, &resultado.Password)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)

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
	fmt.Println(newUser.Password)
	var resultado simpleID
	//currentTime.Format("02/01/2006")
	err = db.QueryRow("select id_usuario from usuario where username = '" + newUser.Username + "' and contrasena = '" + newUser.Password + "'").Scan(&resultado.ID)

	if err != nil {
		fmt.Println("Error running query")
		//fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(resultado.ID)
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
		fmt.Println("adas")
		fmt.Println(err)
		return
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/creaUser", createUser).Methods("POST")
	router.HandleFunc("/login", loginUser).Methods("POST")
	http.HandleFunc("/upload", uploadFile)
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
membresia





*/
