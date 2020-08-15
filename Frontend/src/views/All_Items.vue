<template>
  <v-row id="all">
    <v-col cols="12" sm="6" md="4">
      <!-- product component-->
      <br />
      <br />
      <br />
      <br />
      <br />
      <v-container>
        <v-row justify="space-around">
          <v-card class="mx-auto" elevation="2" outlined>
            <v-card-title>
              Products
              <v-spacer></v-spacer>
              <v-text-field
                v-model="searchProduct"
                append-icon="mdi-magnify"
                label="Search"
                single-line
                hide-details
              ></v-text-field>
            </v-card-title>
            <v-data-table
              :headers="headersProducts"
              :items="products"
              :search="searchProduct"
            ></v-data-table>
          </v-card>
        </v-row>
      </v-container>
      <!-- /product component-->
    </v-col>

    <v-col cols="12" sm="6" md="4">
      <!-- datepicker -->
      <v-dialog
        ref="dialog"
        v-model="modal"
        :return-value.sync="date"
        persistent
        width="290px"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="date"
            label="Picker in dialog"
            prepend-icon="event"
            readonly
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker v-model="date" scrollable>
          <v-spacer></v-spacer>
          <v-btn text color="primary" @click="modal = false">Cancel</v-btn>
          <v-btn text color="primary" @click="$refs.dialog.save(date)"
            >OK</v-btn
          >
        </v-date-picker>
      </v-dialog>
      <!-- /datepicker -->

      <!-- button -->
      <div class="text-center">
        <v-btn class="ma-2" tile color="primary" dark @click="fetch"
          >Search</v-btn
        >
      </div>
      <!-- /button -->

      <!-- buyer component-->
      <v-container>
        <v-row justify="space-around">
          <v-card class="mx-auto" elevation="2" outlined>
            <v-card-title>
              Buyers
              <v-spacer></v-spacer>
              <v-text-field
                v-model="searchBuyer"
                append-icon="mdi-magnify"
                label="Search"
                single-line
                hide-details
              ></v-text-field>
            </v-card-title>
            <v-data-table
              :headers="headersBuyers"
              :items="buyers"
              :search="searchBuyer"
            ></v-data-table>
          </v-card>
        </v-row>
      </v-container>
      <!-- /buyer component-->
    </v-col>

    <v-col cols="12" sm="6" md="4">
      <!-- transaction component-->
      <br />
      <br />
      <br />
      <br />
      <br />
      <v-container id="transaction">
        <v-row justify="space-around">
          <v-card class="mx-auto" elevation="2" outlined>
            <v-card-title>
              Transactions
              <v-spacer></v-spacer>
              <v-text-field
                v-model="searchTransaction"
                append-icon="mdi-magnify"
                label="Search"
                single-line
                hide-details
              ></v-text-field>
            </v-card-title>
            <v-data-table
              :headers="headersTransactions"
              :items="transactions"
              :search="searchTransaction"
            ></v-data-table>
          </v-card>
        </v-row>
      </v-container>
      <!-- /transaction component-->
    </v-col>
  </v-row>
</template>

<script>
import axios from "axios";
import qs from "qs";

export default {
  data: () => ({
    date: new Date("2020-08-11").toISOString().substr(0, 10),
    menu: false,
    modal: false,

    searchProduct: "",
    searchBuyer: "",
    searchTransaction: "",

    headersBuyers: [
      { text: "Name", align: "start", value: "BuyerName" },
      { text: "Buyer Id", value: "BuyerId" },
      { text: "Age", value: "BuyerAge" },
    ],
    buyers: [],

    headersProducts: [
      { text: "Name", align: "start", value: "ProductName" },
      { text: "Product Id", value: "ProductId" },
      { text: "Price", value: "ProductPrice" },
    ],
    products: [],

    headersTransactions: [
      { text: "Transaction Id", align: "start", value: "TransactionId" },
      { text: "Transaction Ip ", value: "TransactionIp" },
      { text: "Device", value: "TransactionDevice" },
    ],

    transactions: [],
  }),
  methods: {
    fetch() {
      console.log(this.date);
      console.log(new Date(this.date).getTime() / 1000);
    },
  },
  created() {
    /* -- Fetch all Buyers, Products and Transactions given the 
          current date in unix timestamp format -- */
    axios
      .post(
        "http://127.0.0.1:3000/allbydate",
        JSON.stringify({
          Date: (new Date(this.date).getTime() / 1000).toString(),
        })
      )
      .then((response) => {
        console.log(response.data.data);
        this.buyers = JSON.parse(JSON.stringify(response.data.data.buyers));
        this.products = JSON.parse(JSON.stringify(response.data.data.products));
        this.transactions = JSON.parse(
          JSON.stringify(response.data.data.transactions)
        );
      });
  },
};
</script>

<style>
/* .container {
  padding: 2px 1px 1px 1px;
} */
</style>
