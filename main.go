package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

type MyHandler struct {

}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	path := dir + "/static/" + r.URL.Path[1:]
	log.Println(path)
	log.Println(r.URL.Path[1:])

	if r.URL.Path[1:] == "" {
		path = path + "index.html"
		data, err := ioutil.ReadFile(string(path))
		if err == nil {
			w.Write(data)
		}
	} else {
		http.ServeFile(w, r, string(path))
	}
}

func main() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":4000", nil)
}