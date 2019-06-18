package main

import (
	"fmt"
	"go5/util"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func createFBS(w http.ResponseWriter, r *http.Request) {
	var fbs []int
	for i := 0; i < 10000000; i++ {
		fbs = util.GetFibonacciSerie(50)
	}
	w.Write([]byte(fmt.Sprintf("%v", fbs)))

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/fb", createFBS)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
