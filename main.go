package main

import (
	"net/http"

	"github.com/gorilla/mux"
)


func main(){
	r:= mux.NewRouter().StrictSlash(false)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/", fs)
	r.HandleFunc("/", index)
	r.HandleFunc("/add", andrey.Post)
	r.HandleFunc("/salvar", andrey.salva)
	r.HandleFunc("/edit/{id}", andrey.edit)
	r.HandleFunc("/atualiza/{id}", andrey.atualizar)
	r.HandleFunc("/delete/{id}", andrey.delete)
	servidor := http.Server{}
	servidor.Addr=":8080"
	servidor.Handler = r
	
	servidor.ListenAndServe()
}

