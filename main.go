package main

import (
	"fmt"
	"log"
	"net/http"

	"rsc.io/quote"
	quotev2 "rsc.io/quote/v2"
	"rsc.io/sampler"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(quotev2.HelloV2())
	fmt.Println(sampler.Hello())

	rtr := mux.NewRouter()

	rtr.HandleFunc("/{topic}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte("Topic: " + vars["topic"]))
	})
	http.Handle("/", rtr)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
