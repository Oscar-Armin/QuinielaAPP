package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/godror/godror"
)

// Types
type task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type allTasks []task

// Persistence
var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome the my GO API!")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	conection()
	/*router := mux.NewRouter().StrictSlash(true)

	/*router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tas", getTasks)
	log.Fatal(http.ListenAndServe(":3000", router))*/
}

func conection() {
	db, err := sql.Open("godror", "ADMIN/1234@172.17.0.2:1521/ORCL18")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select first_name,last_name from animal")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	fmt.Println(rows)
	for rows.Next() {
		fmt.Println("pase")
		var nombre string
		var apellido string
		rows.Scan(&nombre, &apellido)
		fmt.Println("Departamento es " + nombre + " con localidad " + apellido)
	}

}
