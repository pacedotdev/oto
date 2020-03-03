package main

import (
	"example/api"
	"example/generated"
	"fmt"
	"log"
	"net/http"

	"github.com/pacedotdev/oto/otohttp"
)

func main() {
	g := api.GreeterService{}
	server := otohttp.NewServer()
	generated.RegisterGreeterService(server, g)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.Handle("/oto/", server)
	fmt.Println("listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
