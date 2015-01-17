package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/bslatkin/template-perf/static"
)

type Data struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Sex   string `json:"sex"`
	Legs  int    `json:"legs"`
}

var (
	clientTemplate = template.Must(template.New("static").Parse(static.Files["client_render.tpl"]))
	serverTemplate = template.Must(template.New("dynamic").Parse(static.Files["server_render.tpl"]))
	data           = []Data{
		Data{"Pillar", "Ticked Tabby", "Female (neutered)", 3},
		Data{"Hedtral", "Tuxedo", "Male (neutered)", 4},
	}
)

func clientRenderHandler(w http.ResponseWriter, r *http.Request) {
	encoded, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buffer bytes.Buffer
	json.HTMLEscape(&buffer, encoded)

	if err := clientTemplate.Execute(w, template.JS(buffer.Bytes())); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serverRenderHandle(w http.ResponseWriter, r *http.Request) {
	if err := serverTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func init() {
	http.HandleFunc("/client_render", clientRenderHandler)
	http.HandleFunc("/server_render", serverRenderHandle)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
