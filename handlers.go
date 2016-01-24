package main

import (
	"html/template"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Something went wrong :-("))
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	rules, err := DiskStore.GetAllRules()
	if err != nil {
		ErrorPage(w, r)
		return
	}

	tmpl := template.New("base")
	template.Must(tmpl.Parse(BaseTemplateStr))
	template.Must(tmpl.Parse(AdminTemplateStr))
	tmpl.Execute(w, rules)
}

func NewRuleHandler(w http.ResponseWriter, r *http.Request) {

}

func CreateRuleHandler(w http.ResponseWriter, r *http.Request) {


}

func ViewRuleHandler(w http.ResponseWriter, r *http.Request) {

}

func EditRuleHandler(w http.ResponseWriter, r *http.Request) {

}

func UpdateRuleHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteRuleHandler(w http.ResponseWriter, r *http.Request) {

}
