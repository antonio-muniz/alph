package main

import (
	"fmt"
	nethttp "net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http"
)

func main() {
	sys, err := internal.System()
	if err != nil {
		panic(err)
	}
	router := http.Router(sys)
	fmt.Println("Starting server at port 8080...")
	nethttp.ListenAndServe(":8080", router)
}
