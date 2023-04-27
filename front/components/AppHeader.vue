<template>
  <v-app-bar app color="primary" dark>
    <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
    <v-toolbar-title>My App</v-toolbar-title>
    <label for="name">Name:{{ currentUser.name }}</label>
    <v-spacer></v-spacer>
  </v-app-bar>
</template>

<script setup lang="ts">
import { useAuth } from '@/composables/useAuth'
const drawer = ref(false)
const { currentUser } = useAuth()
</script>
