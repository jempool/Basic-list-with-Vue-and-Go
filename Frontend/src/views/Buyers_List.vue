<template>
  <div id="buyers-list">
    <v-card>
      <!-- Buyer List Table -->
      <v-card-title>
        <v-text-field
          v-model="searchBuyers"
          append-icon="search"
          label="Search"
          single-line
          hide-details
        ></v-text-field>
      </v-card-title>
      <v-data-table
        id="buyer-list-table"
        :headers="headersBuyers"
        :items="buyers"
        :search="searchBuyers"
        :items-per-page="itemsBuyers"
        @click:row="handleClick"
      ></v-data-table>
      <!-- /Buyer List Table -->

      <!-- Dialog Details -->
      <div class="text-center">
        <v-dialog v:bind="dialog" v-model="open_dialog" width="auto">
          <v-card>
            <v-card-text>
              <!-- Card content -->
              <br />
              <!-- Details Table -->
              <v-card elevation="2" outlined>
                <v-card-title
                  >Purchase history
                  <v-divider class="mx-4 white" inset vertical></v-divider>
                  <v-text-field
                    v-model="searchDetails"
                    append-icon="search"
                    label="Search"
                    single-line
                    hide-details
                  ></v-text-field>
                </v-card-title>
                <v-data-table
                  id="table-details"
                  :headers="headersDetails"
                  :items="details"
                  :search="searchDetails"
                  :items-per-page="itemsDetails"
                ></v-data-table>
              </v-card>
              <!-- /Details Table -->
              <br />
              <!-- Recomendations Table -->
              <v-card elevation="2" outlined>
                <v-card-title
                  >Buyers with same ip
                  <v-divider class="mx-4 white" inset vertical></v-divider>
                  <v-text-field
                    v-model="searchRecommendations"
                    append-icon="search"
                    label="Search"
                    single-line
                    hide-details
                  ></v-text-field>
                </v-card-title>
                <v-data-table
                  id="table-details"
                  :headers="headersRecommendations"
                  :items="recommendations"
                  :search="searchRecommendations"
                  :items-per-page="itemsRecommendations"
                ></v-data-table>
              </v-card>
              <!-- /Recomendations Table -->

              <!-- /Card content -->
            </v-card-text>
            <!-- <v-divider></v-divider> -->
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="primary" text @click="open_dialog = false">
                OK
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </div>
      <!-- /Dialog Details -->
    </v-card>
  </div>
</template>

<script>
import axios from "axios";
var dateFormat = require("dateformat");

export default {
  data: () => ({
    //Selected buyer
    BuyerAge: "",
    BuyerId: "",
    BuyerName: "",
    open_dialog: false,

    //Buyers
    buyers: [],
    itemsBuyers: 7,
    searchBuyers: "",
    headersBuyers: [
      { text: "Name", align: "start", value: "BuyerName" },
      { text: "Buyer Id", value: "BuyerId" },
      { text: "Age", value: "BuyerAge" },
    ],

    //Details buyer
    details: [],
    itemsDetails: 5,
    searchDetails: "",
    headersDetails: [
      { text: "Transaction Id", align: "start", value: "TransactionId" },
      { text: "Transaction Ip", value: "TransactionIp" },
      { text: "Device", value: "TransactionDevice" },
      { text: "Date", value: "TransactionDate" },
      { text: "Products", value: "Products" },
    ],

    //Details buyer
    recommendations: [],
    itemsRecommendations: 5,
    searchRecommendations: "",
    headersRecommendations: [
      { text: "Name", align: "start", value: "BuyerName" },
      { text: "Recommended product", value: "ProductName" },
    ],
  }),
  methods: {
    handleClick(event) {
      this.open_dialog = true;
      this.BuyerAge = JSON.parse(JSON.stringify(event)).BuyerAge;
      this.BuyerId = JSON.parse(JSON.stringify(event)).BuyerId;
      this.BuyerName = JSON.parse(JSON.stringify(event)).BuyerName;

      axios
        .post(
          "http://127.0.0.1:3000/buyerdetails",
          JSON.stringify({ Data: this.BuyerId })
        )
        .then((response) => {
          let res = JSON.parse(JSON.stringify(response.data.data));
          this.recommendations = JSON.parse(
            JSON.stringify(res.sameIp_someProducts)
          ).map((x) => {
            let json = {
              BuyerName: x.buyer.BuyerName,
              ProductName: x.product1.ProductName,
            };
            return json;
          });

          let food = [];
          res.purchase_history.map((x, i) => x.product1 ? (food[i] = x.product1.ProductName) : null );
          res.purchase_history.map((x,i) => x.product2 ? (food[i] += ` | ${x.product2.ProductName}`) : null );
          res.purchase_history.map((x,i) => x.product3 ? (food[i] += ` | ${x.product3.ProductName}`) : null );
          res.purchase_history.map((x,i) => x.product4 ? (food[i] += ` | ${x.product4.ProductName}`) : null );
          res.purchase_history.map((x,i) => x.product5 ? (food[i] += ` | ${x.product5.ProductName}`) : null );
          //TransactionDate: 1597017600, product1
          this.details = res.purchase_history.map((item, i) => {
            let json = {
              TransactionId: item.TransactionId,
              TransactionIp: item.TransactionIp,
              TransactionDevice: item.TransactionDevice,
              TransactionDate: dateFormat( new Date(item.TransactionDate * 1000), "isoDate" ),
              Products: food[i],
            };
            return json;
          });
        });
    },
  },
  created() {
    /* -- Fetch all Buyers -- */
    axios.get("http://127.0.0.1:3000/allbuyers").then((response) => {
      this.buyers = JSON.parse(JSON.stringify(response.data.data.buyers));
    });
  },
};
</script>

<style>
#buyers-list {
  padding: 12px 20px 10px 20px;
}
#buyer-list-table tr {
  cursor:pointer;
}
#table-details tr:hover {
  background-color: transparent !important;
}
</style>
