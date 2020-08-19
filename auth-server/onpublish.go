package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var STREAMING_KEY string

func onpublish(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	fmt.Println("Request", r)
	fmt.Println("URL", r.URL)
	fmt.Println("Form", r.Form)
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Body", string(b))
	key := r.Form.Get("name")
	if key == STREAMING_KEY {
		fmt.Printf("good key %s == %s\n", key, STREAMING_KEY)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Good to go"))
	} else {
		fmt.Printf("bad key %s != %s\n", key, STREAMING_KEY)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid straming key"))
	}
}
