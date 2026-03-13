package main

import (
	"example/public"
	"html/template"
	"log"
	"net/http"

	"github.com/siuyin/dflt"
)

var t *template.Template

func main() {
	port := dflt.EnvString("PORT", "8080")
	log.Printf("PORT=%s", port)
	http.Handle("/", http.FileServer(http.FS(public.Content)))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
