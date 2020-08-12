<template >
  <!-- <v-card class="justify-center" max-width="500"> -->
    <v-list :three-line="threeLine" :rounded="rounded" id="buyers">
      <!-- <v-subheader align="center" class="justify-center">REPORTS</v-subheader> -->
      <v-list-item-group color="primary" align="center" class="justify-center">
        <v-list-item v-for="(item, i) in buyers" :key="i">
          <v-list-item-content>
                         <v-divider class="mx-4 gray" inset></v-divider> 
            <v-list-item-title v-html="item.title"></v-list-item-title>
            <v-list-item-subtitle v-if="threeLine" v-html="item.subtitle"></v-list-item-subtitle>
             <v-divider class="mx-4 gray" inset></v-divider> 
          </v-list-item-content>
        </v-list-item>
      </v-list-item-group>
    </v-list>
  <!-- </v-card> -->
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      buyers: [],
      threeLine: true,
      rounded: false,
    };
  },
  created() {
    axios.get("http://127.0.0.1:3000/buyers").then((response) => {  
      console.log(response.data.data)   
      let items = JSON.parse(JSON.stringify(response.data.data.buyers));
      this.buyers = items.map((item) => {
        let response = {
          title: `<span class='text--primary'><b>Name:</b> ${item.NameBuyer}</span>`,
          subtitle: `<span class='text--primary'><b>Id:</b> ${item.IDBuyer}</span> &mdash; <b>Age:</b> ${item.AgeBuyer}`,
        };
        return response;
      });
    });
  },
};
</script>














