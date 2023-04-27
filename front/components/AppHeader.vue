<template>
  <v-app-bar app color="primary" dark>
    <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
    <v-toolbar-title>My App</v-toolbar-title>
    <label for="name">Name:{{ name }}</label>
    <v-spacer></v-spacer>
  </v-app-bar>
</template>

<script setup lang="ts">
interface Props {
  name: string
}

defineProps<Props>()
const drawer = ref(false)
</script>
