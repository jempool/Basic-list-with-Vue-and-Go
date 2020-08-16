package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

/* -- Struct for the json requests -- */
type requestStruct struct {
	Data string
}

/* -- Function to fetch all Buyers -- */
func dgraphQueryAllBuyers(url string) []byte {

	query := `{"query": "{ buyers(func: has(BuyerName)) { BuyerId BuyerName BuyerAge }}"}`

	var jsonStr = []byte(query)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)

	return []byte(response)
}

/* -- Function to fetch all Buyers, Products and Transactions
   given a date in unix timestamp format -- */
func dgraphQueryAll(url string, date requestStruct) []byte {

	query := `{"query": "{var(func: has(TransactionId)) @filter(eq(TransactionDate, %v)) {product1{uidp1 as uid ProductId ProductName ProductPrice} product2{uidp2 as uid ProductId ProductName ProductPrice} product3{uidp3 as uid ProductId ProductName ProductPrice} product4{uidp4 as uid ProductId ProductName ProductPrice} product5{uidp5 as uid ProductId ProductName ProductPrice}} products(func: has(ProductId)) @filter(uid(uidp1) and uid(uidp2) and uid(uidp3) and uid(uidp4) and uid(uidp5)) {ProductId ProductName ProductPrice} var(func: eq(TransactionDate, %v)) {buyer {uidb as uid}} buyers(func: has(BuyerName)) @filter(uid(uidb)) {BuyerId BuyerName BuyerAge} transactions(func: eq(TransactionDate, %v)) {TransactionId TransactionIp TransactionDevice TransactionDate}}"}`
	query = fmt.Sprintf(query, date.Data, date.Data, date.Data)

	var jsonStr = []byte(query)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)

	return []byte(response)
}

/* -- Function to fetch all Buyers, Products and Transactions
   given a date in unix timestamp format -- */
func dgraphQueryBuyerDetails(url string, BuyerID string) []byte {

	query := `{"query": "{purchase_history(func: eq(BuyerId,%v),orderdesc: TransactionDate){TransactionId ip as TransactionIp TransactionDevice TransactionDate product1{ProductId ProductName ProductPrice}product2{ProductId ProductName ProductPrice}product3{ProductId ProductName ProductPrice}product4{ProductId ProductName ProductPrice}product5{ProductId ProductName ProductPrice}}sameIp_someProducts(func: has(TransactionIp))@filter(eq(TransactionIp, val(ip))and NOT eq(BuyerId,%v)){buyer{BuyerName}product1{ProductName}}}"}`
	query = fmt.Sprintf(query, BuyerID, BuyerID)

	var jsonStr = []byte(query)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)

	return []byte(response)
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

	r := chi.NewRouter()
	const url = "http://localhost:8080/query" // URL to dgraph alpha

	/* -- Handler for requests  all Buyers, Products and Transactions given
	   a date in unix timestamp format -- */
	r.Post("/allbydate", func(w http.ResponseWriter, r *http.Request) {

		setupCORS(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}

		decoder := json.NewDecoder(r.Body)
		var date requestStruct
		err := decoder.Decode(&date)
		if err != nil {
			panic(err)
		}
		// fmt.Println(date.Date)
		response := dgraphQueryAll(url, date)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})

	/* -- Handler for requests all Buyers -- */
	r.Get("/allbuyers", func(w http.ResponseWriter, r *http.Request) {

		setupCORS(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}

		response := dgraphQueryAllBuyers(url)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})

	/* -- Handler for requests details of a Buyer given an
	   id in string format -- */
	r.Post("/buyerdetails", func(w http.ResponseWriter, r *http.Request) {

		setupCORS(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}

		decoder := json.NewDecoder(r.Body)
		var data requestStruct
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}

		reponse := dgraphQueryBuyerDetails(url, data.Data)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(reponse))
	})

	fmt.Println("Serving on port 3000 - " + time.Now().String())
	http.ListenAndServe(":3000", r)
}
