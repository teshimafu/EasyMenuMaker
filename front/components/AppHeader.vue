<template>
  <v-app-bar app color="primary" dark>
    <v-navigation-drawer app>Navigation Lists</v-navigation-drawer>
    <v-app-bar-nav-icon> </v-app-bar-nav-icon>
    <v-toolbar-title>My App</v-toolbar-title>
    <v-spacer></v-spacer>
  </v-app-bar>
</template>

<script setup lang="ts">
import { useAuth } from '@/composables/useAuth'
const drawer = ref(false)
const { currentUser } = useAuth()

const items = reactive([
  { link: `a`, icon: `b`, title: `c` },
  { link: `a`, icon: `b`, title: `c` },
  { link: `a`, icon: `b`, title: `c` },
  { link: `a`, icon: `b`, title: `c` },
  { link: `a`, icon: `b`, title: `c` }
])
</script>
