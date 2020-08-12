<template>
  <v-container>
    <v-row justify="space-around">
      <v-sheet :width="width" :height="height" :color="color">
        <v-card>
          <v-card-title>
            Buyers
            <v-spacer></v-spacer>
            <v-text-field
              v-model="search"
              append-icon="mdi-magnify"
              label="Search"
              single-line
              hide-details
            ></v-text-field>
          </v-card-title>
          <v-data-table :headers="headers" :items="buyers" :search="search"></v-data-table>
        </v-card>
      </v-sheet>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      width: 500,
      height: 500,
      color: "white",

      search: "",
      headers: [
        { text: "Name", align: "start", value: "NameBuyer" },
        { text: "Identification", value: "IDBuyer" },
        { text: "Age", value: "AgeBuyer" },
      ],
      buyers: [],
    };
  },
  created() {
    axios.get("http://127.0.0.1:3000/buyers").then((response) => {
      this.buyers = JSON.parse(JSON.stringify(response.data.data.buyers));
    });
  },
};
</script>