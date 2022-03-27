package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
type usuario struct{
	Nome string
	Email string
	Posts []string
}
func (m *usuario) salva(w http.ResponseWriter, r * http.Request){
	r.ParseForm()
	novo := r.PostFormValue("postagem")
	m.Posts = append(m.Posts, novo)
	http.Redirect(w,r, "/add", 302)
}
func (m *usuario)atualizar(w http.ResponseWriter, r * http.Request){
	v := mux.Vars(r)
	k ,_:= strconv.Atoi(v["id"])
	r.ParseForm()
	m.Posts[k] =r.PostFormValue("postagem")
	http.Redirect(w, r, "/add", 302)
}
func index(w http.ResponseWriter, r * http.Request){

	renderizar(w, "index", "base", andrey)		

}
func (m *usuario)Post(w http.ResponseWriter, r* http.Request){
	
	renderizar(w, "add", "base", m)
	
}
func (m *usuario)edit(w http.ResponseWriter, r* http.Request){
	variavei := mux.Vars(r)
	k := variavei["id"]
	l , _:= strconv.Atoi(k)
	type novo struct{
		Post string
		Id int
	}

	model := novo{m.Posts[l], l}
	renderizar(w, "edit", "base", model)
}
func remove(slice []string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}
func(m*usuario)delete(w http.ResponseWriter, r * http.Request){
	v := mux.Vars(r)
	k,_ := strconv.Atoi(v["id"])
	
	m.Posts = remove(m.Posts, k)
	http.Redirect(w,r, "/add", 302)
}