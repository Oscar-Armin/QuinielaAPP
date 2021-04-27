package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/godror/godror"
	"github.com/gorilla/mux"
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
	Fecha_registro   string `json:"registro"`
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
	var newLog login
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newLog)

	fmt.Println(newLog)

	//currentTime.Format("02/01/2006")
	rows, err := db.Query("select to_char(sysdate, 'HH24:MI:SS') from dual")

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
	json.NewEncoder(w).Encode(newLog)

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
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid User Data")
	}

	json.Unmarshal(reqBody, &newUser)

	//fmt.Println(newUser)

	//currentTime.Format("02/01/2006")
	rows, err := db.Query("select id_usuario from usuario where username = '" + newUser.Username + "' and contrasena = '" + newUser.Password + "'")

	if err != nil {
		fmt.Println("Error running query")
		//fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var resultado simpleID
	//fmt.Println(reflect.TypeOf(rows))
	//var resultado usuario
	for rows.Next() {
		fmt.Println("info")

		rows.Scan(&resultado.ID)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)

}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(tasks)
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

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/creaUser", createUser).Methods("POST")
	router.HandleFunc("/login", loginUser).Methods("POST")
	fmt.Println("En puerto 3080")
	log.Fatal(http.ListenAndServe(":3080", router))

}

func conection() {

}

/*
puntaje
recompensa
registo_membresia
membresia





*/
