package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

const defaultJs = `
rule(function(data){
	return true;
})
`

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
	rule := NewRule("", defaultJs, []string{""})
	tmpl := template.New("base")
	template.Must(tmpl.Parse(BaseTemplateStr))
	template.Must(tmpl.Parse(NewTemplateStr))
	tmpl.Execute(w, rule)

}

func CreateRuleHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	script := r.FormValue("script")
	var endpoints []string
	for i := 0; ; i++ {
		endpoint := r.FormValue("endpoint_" + strconv.Itoa(i))
		if endpoint == "" {
			break
		} else {
			endpoints = append(endpoints, endpoint)
		}
	}

	rule := NewRule(title, script, endpoints)

	err := DiskStore.SaveRule(rule)
	if err != nil {
		ErrorPage(w, r)
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func ViewRuleHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.Replace(r.URL.Path, "/view/", "", 1)
	rule, err := DiskStore.GetRule(id)
	if err != nil {
		ErrorPage(w, r)
		return
	}

	tmpl := template.New("base")
	template.Must(tmpl.Parse(BaseTemplateStr))
	template.Must(tmpl.Parse(ViewTemplateStr))
	tmpl.Execute(w, rule)
}

func EditRuleHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.Replace(r.URL.Path, "/edit/", "", 1)
	rule, err := DiskStore.GetRule(id)
	if err != nil {
		ErrorPage(w, r)
		return
	}
	tmpl := template.New("base")
	template.Must(tmpl.Parse(BaseTemplateStr))
	template.Must(tmpl.Parse(EditTemplateStr))
	tmpl.Execute(w, rule)
}

func UpdateRuleHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	rule, err := DiskStore.GetRule(id)
	if err != nil {
		log.Println("ERROR:", err)
		ErrorPage(w, r)
		return
	}

	(*rule).Title = r.FormValue("title")
	(*rule).Script = r.FormValue("script")
	var endpoints []string
	for i := 0; ; i++ {
		endpoint := r.FormValue("endpoint_" + strconv.Itoa(i))
		if endpoint == "" {
			break
		} else {
			endpoints = append(endpoints, endpoint)
		}
	}
	(*rule).Endpoints = endpoints

	err = DiskStore.SaveRule(rule)
	if err != nil {
		ErrorPage(w, r)
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func DeleteRuleHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	err := DiskStore.DeleteRule(id)
	if err != nil {
		log.Println("ERROR:", err)
		ErrorPage(w, r)
		return
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
