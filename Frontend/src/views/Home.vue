<template>
  <v-container>
    <v-row justify="space-around">
      <v-sheet :width="width" :height="height" :elevation="elevation" :color="color"></v-sheet>
    </v-row>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    width: 500,
    height: 500,
    elevation: 4,
    color: "white",
  }),
};
</script>

