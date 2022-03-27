package main
import (
	"html/template"
	"net/http"
)
var andrey = usuario{Nome:"andrey costa de queiroz",Email:"queirozandrey5@gmail.com" ,Posts: []string{}}
var tmpls map[string]*template.Template
func init(){
	
	if tmpls == nil{
		tmpls = make(map[string]*template.Template)
	}
	tmpls["index"] = template.Must(template.ParseFiles("template/index.html", "template/base.html"))
	tmpls["add"] = template.Must(template.ParseFiles("template/add.html", "template/base.html"))
	tmpls["edit"] = template.Must(template.ParseFiles("template/edit.html", "template/base.html"))
	
}
func renderizar(w http.ResponseWriter, nome string, template string, dados interface{}){
	t, ok := tmpls[nome]
	if !ok {http.Error(w, "fudeu mané", http.StatusInternalServerError)}
	err := t.ExecuteTemplate(w, template, dados)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}