package util

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/govwa/user/session"
)

func increment(num int) int {
	return num + 1
}

func SafeRender(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {

	s := session.New()
	sid := s.GetSession(r, "id") //make uid available to all page
	data["uid"] = sid

	funcs := template.FuncMap(make(map[string]interface{}))
	funcs["inc"] = increment

	template := template.New(name).Funcs(funcs)
	template, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Println(err.Error())
	}

	err = template.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Println(err.Error())
	}
}

func RenderAsJson(w http.ResponseWriter, data ...interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func UnSafeRender(w http.ResponseWriter, name string, data ...interface{}) {

	template := template.Must(template.ParseGlob("templates/*"))
	template.ExecuteTemplate(w, name, data)
}

func ToHTML(text string) template.HTML {
	return template.HTML(text)
}
