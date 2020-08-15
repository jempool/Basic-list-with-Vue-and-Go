package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	// "reflect"
)

type dateStruct struct {
	Date string
}

/* -- Function to fetch all Buyers, Products and Transactions 
      given a date in unix timestamp format -- */
func dgraphQueryAll(url string, date dateStruct) []byte {	

	query := `{"query": "{var(func: has(TransactionId)) @filter(eq(TransactionDate, %v)) {product1{uidp1 as uid ProductId ProductName ProductPrice} product2{uidp2 as uid ProductId ProductName ProductPrice} product3{uidp3 as uid ProductId ProductName ProductPrice} product4{uidp4 as uid ProductId ProductName ProductPrice} product5{uidp5 as uid ProductId ProductName ProductPrice}} products(func: has(ProductId)) @filter(uid(uidp1) and uid(uidp2) and uid(uidp3) and uid(uidp4) and uid(uidp5)) {ProductId ProductName ProductPrice} var(func: eq(TransactionDate, %v)) {buyer {uidb as uid}} buyers(func: has(BuyerName)) @filter(uid(uidb)) {BuyerId BuyerName BuyerAge} transactions(func: eq(TransactionDate, %v)) {TransactionId TransactionIp TransactionDevice TransactionDate}}"}`
	query = fmt.Sprintf(query, date.Date, date.Date, date.Date)

	var jsonStr = []byte(query)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	transactions, _ := ioutil.ReadAll(resp.Body)

	return []byte(transactions)
}

/* -- Function to allow CORS: Access-Control-Allow-Origin -- */
func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

/* -- Go application entrypoint -- */
func main() {

	// -- Api to frontend -- //

	const url = "http://localhost:8080/query" // URL to dgraph alpha

	r := chi.NewRouter()

	/* -- Handler for requests  all Buyers, Products and Transactions given
	   a date in unix timestamp format -- */
	r.Post("/allbydate", func(w http.ResponseWriter, r *http.Request) {

		setupCORS(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}

		decoder := json.NewDecoder(r.Body)
		var date dateStruct
		err := decoder.Decode(&date)
		if err != nil {
			panic(err)
		}

		fmt.Println(date.Date)

		response := dgraphQueryAll(url, date)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})

	// -- All Buyers --//
	r.Post("/transactions", func(w http.ResponseWriter, r *http.Request) {

		setupCORS(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}

		decoder := json.NewDecoder(r.Body)
		var date dateStruct
		err := decoder.Decode(&date)
		if err != nil {
			panic(err)
		}
		//1596931200, " + Date + "
		// query := `{"query": "{ transactions(func: eq(Date, " + Date + ")){ IDTransaction IPBuyer DevicesBuyer Products{ IDProduct }}}"}`
		transactions := dgraphQueryAll(url, date)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(transactions))
	})

	// -- Products -- //
	r.Post("/products", func(w http.ResponseWriter, r *http.Request) {

		setupCORS(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}

		decoder := json.NewDecoder(r.Body)
		var date dateStruct
		err := decoder.Decode(&date)
		if err != nil {
			panic(err)
		}

		products := dgraphQueryAll(url, date)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(products))
	})

	fmt.Println("Serving on port 3000 - " + time.Now().String())
	http.ListenAndServe(":3000", r)
}
