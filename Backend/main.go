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
	Date int32
}

func dgraphQuery(url string, query string, Date dateStruct) []byte {

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

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

//Go application entrypoint
func main() {

	// -- Fetching data from Dgraph -- //

	url := "http://localhost:8080/query"

	// -- Some buyer
	query4 := `{"query": "{ query(func: has(IDBuyer)){ IDBuyer NameBuyer AgeBuyer Transactions{ IDTransaction IPBuyer DevicesBuyer Products{ IDProduct NameProduct PriceProduct }}}}"}`
	var jsonStr4 = []byte(query4)
	req4, err4 := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr4))
	req4.Header.Set("Content-Type", "application/json")

	client4 := &http.Client{}
	resp4, err4 := client4.Do(req4)
	if err4 != nil {
		panic(err4)
	}
	defer resp4.Body.Close()
	buyer, _ := ioutil.ReadAll(resp4.Body)
	// fmt.Println("response Body:", string(body))

	// -- Api to frontend -- //

	r := chi.NewRouter()

	// -- Buyers --//
	r.Post("/buyers", func(w http.ResponseWriter, r *http.Request) {

		setupCORS(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}

		decoder := json.NewDecoder(r.Body)
		var Date dateStruct
		err := decoder.Decode(&Date)
		if err != nil {
			panic(err)
		}

		query := `{"query": "{ buyers(func: has(IDBuyer)){ IDBuyer NameBuyer AgeBuyer }}"}`
		buyers := dgraphQuery(url, query, Date)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buyers))
	})

	// -- Transactions --//
	r.Post("/transactions", func(w http.ResponseWriter, r *http.Request) {

		setupCORS(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}

		decoder := json.NewDecoder(r.Body)
		var Date dateStruct
		err := decoder.Decode(&Date)
		if err != nil {
			panic(err)
		}
		//1596931200, 1597017600
		query := `{"query": "{ transactions(func: eq(Date, 1597017600)){ IDTransaction IPBuyer DevicesBuyer Products{ IDProduct }}}"}`
		transactions := dgraphQuery(url, query, Date)

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
		var Date dateStruct
		err := decoder.Decode(&Date)
		if err != nil {
			panic(err)
		}

		query := `{"query": "{ products(func: has(IDProduct )){ IDProduct NameProduct PriceProduct }}"}`
		products := dgraphQuery(url, query, Date)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(products))
	})

	r.Get("/buyer", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buyer))
	})

	fmt.Println("Serving on port 3000 - " + time.Now().String())
	http.ListenAndServe(":3000", r)
}
