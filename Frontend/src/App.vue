<template>
  <v-app>
    <template>
      <div>
        <v-app-bar color="primary" dense dark>
          <v-app-bar-nav-icon></v-app-bar-nav-icon>

          <!-- <router-link to="/">
            <v-toolbar-title class="white--text">All Items</v-toolbar-title>
          </router-link>
          <v-divider class="mx-4 white" inset vertical></v-divider> -->
          <router-link to="/all_items">
            <v-toolbar-title class="white--text">All Items</v-toolbar-title>
          </router-link>
          <v-divider class="mx-4 white" inset vertical></v-divider>
          <router-link to="/buyerslist">
            <v-toolbar-title class="white--text">Compradores</v-toolbar-title>
          </router-link>
          <v-divider class="mx-4 white" inset vertical></v-divider>
          <router-link to="/buyer">
            <v-toolbar-title class="white--text"
              >Historial de compras</v-toolbar-title
            >
          </router-link>

          <v-menu left bottom>
            <template v-slot:activator="{ on, attrs }">
              <v-btn icon v-bind="attrs" v-on="on">
                <v-icon>mdi-dots-vertical</v-icon>
              </v-btn>
            </template>
            <v-list>
              <v-list-item v-for="n in 5" :key="n" @click="() => {}">
                <v-list-item-title>Option {{ n }}</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
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
  /* background:#f5f5f5 */
}
</style>
