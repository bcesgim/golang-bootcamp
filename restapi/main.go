package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Product struct {
	ID 	  string `json:"id, omitempty"`// bson:"_id,omitempty"`
	Title string `json:"title, omitempty"`// bson:"username,omitempty"`
	Price float64 `json:"price, omitempty"` // bson:"firstname,omitempty"`
}

type Cart struct {
	Product  Product `json:"item, omitempty"`
	Quantity int `json:"quantity, omitempty"`
}

var products []Product

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//Write Status Code
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	fmt.Fprintln(w, "Example-1; Request was succesful")
}

func GetItemsEndpoint (response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(products)
}

func GetItemEndpoint (response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // Gets params
	// Loop through players and find one with the id from the params
	for _, item := range products {
		if item.ID == params["id"]{
			json.NewEncoder(response).Encode(item)
			return
		}
	}
	json.NewEncoder(response).Encode(&Product{})
}

func CreateItemEndpoint (response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	var product Product
	_ = json.NewDecoder(request.Body).Decode(&product)
	product.ID = params["id"]
	products = append(products, product)
	json.NewEncoder(response).Encode(products)
}

func UpdateItemEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			var product Product
			_ = json.NewDecoder(request.Body).Decode(&product)
			product.ID = params["id"]
			products = append(products, item)
			json.NewEncoder(response).Encode(product)
			return
		}
	}
}

func DeleteItemEndpoint (response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[: index], products[index + 1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(products)
}


func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/example", myHandlerFunc)

	// Hardcoded data - @todo: add database
	products = append(products, Product { ID: "1", Title: "Item 1",  Price: 42 })
	products = append(products, Product { ID: "2", Title: "Item 2", Price: 199.99 })

	// Route handlers & endpoints
	mux.HandleFunc("/cart", GetItemsEndpoint)//.Methods("GET")
	mux.HandleFunc("/cart/{id}", GetItemEndpoint)//.Methods("GET")
	mux.HandleFunc("/cart/{id}", CreateItemEndpoint)//.Methods("POST")
	mux.HandleFunc("/cart/{id}", UpdateItemEndpoint)//.Methods("PUT")
	mux.HandleFunc("/cart/{id}", DeleteItemEndpoint)//.Methods("DELETE")

	server := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           mux,
		ReadTimeout:       30 * time.Second, //Reading all of the request including body
		ReadHeaderTimeout: 15 * time.Second, //For Request headers
		WriteTimeout:      30 * time.Second, //Timeout of writing the response
		IdleTimeout:       10 * time.Second, //Timeout for Idle connections, when keepalive is true
		//MaxHeaderBytes: Size of Request Headers accepted.
		ErrorLog: log.New(os.Stdout, "Server: ", log.Lshortfile), //Logger for Server errors
	}

	//Function that can be invoked when state changes on the connections
	server.ConnState = func(c net.Conn, cs http.ConnState) {
		log.Println("Connection from Address: " + c.RemoteAddr().String() + " - State: " + cs.String())
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		log.Println("Server Starting")
		log.Fatal(server.ListenAndServe())
	}()
	<-c
	// Shut down, waiting no longer than 30 seconds before halting
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println("Closing down")
}
