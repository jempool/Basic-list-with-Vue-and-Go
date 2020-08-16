<template>
  <v-app>
    <template>
      <div>
        <v-app-bar color="primary" dense dark>
          <v-app-bar-nav-icon></v-app-bar-nav-icon>
          <router-link to="/">
            <v-toolbar-title class="white--text">Home</v-toolbar-title>
          </router-link>
          <v-divider class="mx-4 white" inset vertical></v-divider>
          <router-link to="/all_data">
            <v-toolbar-title class="white--text">All Data</v-toolbar-title>
          </router-link>
          <v-divider class="mx-4 white" inset vertical></v-divider>
          <router-link to="/buyers_list">
            <v-toolbar-title class="white--text">Buyers List</v-toolbar-title>
          </router-link>                  
        </v-app-bar>
      </div>
    </template>
    <router-view />
  </v-app>
</template>

<script>
export default {
  name: "App",
  data: () => ({
    //
  }),
};
</script>

<style>
a {
  text-decoration: none !important;
}
#all {
  padding: 5px 20px 5px 20px;
}
</style>
