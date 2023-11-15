package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", holaMundo)

	http.HandleFunc("/suma", handleSuma)

	http.HandleFunc("/chau", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Chau Mundo!!!")
	})

	fmt.Println("Servidor escuchando en localhost:8080")
	http.ListenAndServe(":8080",nil)
}

type SumaRequest struct {
	Numero1 int `json:"numero1"`
	Numero2 int `json:"numero2"`
}

type SumaResponse struct {
	Suma int `json:"suma"`
}

func handleSuma(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req SumaRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err!= nil {
		http.Error(w, "Error en el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	resultado := req.Numero1 + req.Numero2
	respuesta := SumaResponse{Suma: resultado}

	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(respuesta) 
	if err!=nil{
		http.Error(w, "Error al crear respuesta", http.StatusInternalServerError)
		return
	}

}

func holaMundo(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet{
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w,"Hola Mundo!!!")
}