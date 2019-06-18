package main

import (
	"go8/rest_test/handler"
	"log"
	"net/http"
	"strconv"
)

const HttpHost = "127.0.0.1"
const HttpPort = 8082

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "POST":
			handler.Create(writer, request)
		case "PUT":
			handler.Update(writer, request)
		case "GET":
			handler.Read(writer, request)
		case "DELETE":
			handler.Delete(writer, request)
		case "OPTIONS":
			writer.Header().Set("Access-Control-Allow-Origin", "*")
			writer.Header().Add("Access-Control-Allow-Headers", "*")
			writer.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
			writer.WriteHeader(204)
		default:
			log.Printf("Method is not allowed: %s", request.Method)
			return
		}

	})

	log.Fatal(http.ListenAndServe(HttpHost+":"+strconv.Itoa(HttpPort), nil))
}
