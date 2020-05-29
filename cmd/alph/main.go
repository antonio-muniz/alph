package main

import (
	"fmt"
	nethttp "net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http"
)

func main() {
	router := http.Router()
	fmt.Println("Starting server at port 8080...")
	nethttp.ListenAndServe(":8080", router)
}
