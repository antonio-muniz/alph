package main

import (
	"fmt"
	nethttp "net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http"
)

func main() {
	components, err := internal.Components()
	if err != nil {
		panic(err)
	}
	router := http.Router(components)
	fmt.Println("Starting server at port 8080...")
	nethttp.ListenAndServe(":8080", router)
}
