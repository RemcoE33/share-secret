package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/RemcoE33/share-secret/handlers"
	"github.com/RemcoE33/share-secret/store"
	chi "github.com/go-chi/chi/v5"
)

//go:embed all:uibuild
var svelteStatic embed.FS

// config settings
var (
	project    string = "ghco-invest" //firestore projects
	collection string = "secrets"     //firestore collection
	port       string = ":8080"
)

func main() {
	fb, err := store.NewFirestore(project, collection)
	if err != nil {
		panic("could not initialize store")
	}
	server := handlers.NewServer(fb)
	r := chi.NewRouter()

	// building embedded directory for the svelte frontend
	s, err := fs.Sub(svelteStatic, "uibuild")
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(http.FS(s))

	r.Handle("/", staticServer)
	r.Handle("/_app/*", staticServer)                                 // Need to serve any app components from the embedded files
	r.Handle("/favicon.png", staticServer)                            // Also serve favicon
	r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) { // Everything else returns the index
		r.URL.Path = "/" // Replace the request path
		staticServer.ServeHTTP(w, r)
	})

	// API endpooints
	r.Post("/api/secret", server.HandleCreateSecret)
	r.Get("/api/secret/{id}", server.HandleGetSecret)

	p := os.Getenv("PORT")

	if p != "" {
		port = p
	}

	fmt.Println("Running on port: ", port)
	log.Fatal(http.ListenAndServe(port, r))

}
