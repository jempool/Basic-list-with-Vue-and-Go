package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/go-chi/chi"
)

//buyer struct
type buyer struct {
	IDBuyer   string
	NameBuyer string
	AgeBuyer  int
}

var buyers = []buyer{
	buyer{
		IDBuyer:   "14627bb1",
		NameBuyer: "Kathryne",
		AgeBuyer:  19,
	},
	buyer{
		IDBuyer:   "e431bdf3",
		NameBuyer: "Piggy",
		AgeBuyer:  25,
	},
	buyer{
		IDBuyer:   "899efe6",
		NameBuyer: "Clyde",
		AgeBuyer:  46,
	},
}

//products struct
type product struct {
	IDProduct    string
	NameProduct  string
	PriceProduct int
}

var products = []product{
	product{
		IDProduct:    "9187f1d3",
		NameProduct:  "Progresso Traditional Chicken Noodle Soup",
		PriceProduct: 5978,
	},
	product{
		IDProduct:    "210dea49",
		NameProduct:  "Creamy tomato soup",
		PriceProduct: 495,
	},
	product{
		IDProduct:    "fc5de8c5",
		NameProduct:  "Vegetable broth",
		PriceProduct: 2939,
	},
}

type transaction struct {
	IDTransaction string
	IDBuyer       string
	IPBuyer       string
	DevicesBuyer  []string
	IDProducts    []string
}

var transactions = []transaction{
	transaction{
		IDTransaction: "#00005f29f680",
		IDBuyer:       "d1c4a356",
		IPBuyer:       "100.90.107.125",
		DevicesBuyer:  []string{"linux"},
		IDProducts:    []string{"7eeb79ef", "1f31e33b", "f42fe530"},
	},
	transaction{
		IDTransaction: "#00005f29f681",
		IDBuyer:       "e5f46faa",
		IPBuyer:       "133.231.98.30",
		DevicesBuyer:  []string{"windows"},
		IDProducts:    []string{"8324b3dc", "7eeb79ef", "b1a92834"},
	},
	transaction{
		IDTransaction: "#00005f29f682",
		IDBuyer:       "8db99729",
		IPBuyer:       "14.117.89.93",
		DevicesBuyer:  []string{"windows"},
		IDProducts:    []string{"f8228bd8", "ff850bc"},
	},
}

//Go application entrypoint
func main() {

	r := chi.NewRouter()

	r.Get("/buyers", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(buyers)
	})

	r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(products)
	})

	r.Get("/transactions", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(transactions)
	})

	fmt.Println("Serving on port 3000 - " + time.Now().String())
	http.ListenAndServe(":3000", r)
}
