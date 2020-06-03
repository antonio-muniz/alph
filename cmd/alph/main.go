package main

import (
	"context"
	"fmt"
	nethttp "net/http"

	"github.com/antonio-muniz/alph/cmd/alph/internal"
	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http"
)

func main() {
	ctx := context.Background()
	sys, err := internal.System(ctx)
	if err != nil {
		panic(err)
	}
	router := http.Router(sys)
	fmt.Println("Starting server at port 8080...")
	nethttp.ListenAndServe(":8080", router)
}
