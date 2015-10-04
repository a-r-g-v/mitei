package main

import (
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/goji/param"
	"github.com/zenazn/goji/web"
	"net/http"
	"strconv"
)

func IndexController(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl, err := pongo2.DefaultSet.FromFile("./view/index.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tables := []Table{}
	DB.Find(&tables)
	tpl.ExecuteWriter(pongo2.Context{"naptlist": tables}, w)

}

func CreateController(c web.C, w http.ResponseWriter, r *http.Request) {
	var table Table
	r.ParseForm()
	err := param.Parse(r.Form, &table)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := Allocate(table.TargetIP, table.TargetPort, table.BoundPort)
	if result != true {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "Create")

}
func RemoveController(c web.C, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(c.URLParams["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	table := Table{}
	DB.First(&table, id)
	fmt.Print(table.Id, table.TargetIP)

	if table.Id == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	result := Release(table.TargetIP, table.TargetPort, table.BoundPort)
	if result != true {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "Remove")
}
