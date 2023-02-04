package main

import (
	"embed"
	"net/http"
	"rest-api-sim/connection"
	"rest-api-sim/handler"
	"rest-api-sim/helper"
	"rest-api-sim/model"
	"text/template"
)

func main() {
	// handling all endpoint
	mux := http.NewServeMux()

	// endpoint of "index.html" will be "/"
	// result "index.html" is:
	// - database connection status
	// - rows.Next(), means query all records from database
	// - form name with post method action to database
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		helper.Output(err)

		name := r.Form.Get("form-name")
		button := r.FormValue("submit")

		ping := connection.Status()

		var sUser model.User
		var info string

		if ping {
			if name == "" {
				sUser = handler.QueryAllData()
			} else if name != "" && button == "pres" {
				info = handler.QueryExec(name)
				sUser = handler.QueryAllData()
			} else if name != "" && button == "del" {
				info = handler.QueryExecDel(name)
				sUser = handler.QueryAllData()
			}
		}

		ExecutorFiles().ExecuteTemplate(w, "index.html", map[string]interface{}{
			"Name":     sUser.Name,
			"Presence": sUser.Presence,
			"Absence":  sUser.Absence,
			"Status":   ping,
			"Info":     info,
		})
	})

	mux.HandleFunc("/information.html", func(w http.ResponseWriter, r *http.Request) {
		ExecutorFiles().ExecuteTemplate(w, "information.html", nil)
	})

	http.ListenAndServe("localhost:3000", mux)
}

//go:embed views/*
var views embed.FS

func ExecutorFiles() *template.Template {
	temp, err := template.ParseFS(views, "views/*")
	helper.Output(err)
	return temp
}
