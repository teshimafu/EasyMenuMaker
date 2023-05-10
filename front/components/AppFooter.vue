<template>
  <v-footer app absolute height="24" class="bg-grey-lighten-1">
    <div class="m-auto">
      <div>&copy; {{ copyRight }}</div>
    </div>
  </v-footer>
</template>

<script setup lang="ts">
const copyRight = `2023 - ${new Date().getFullYear()} EasyMenuMaker`
</script>

<style scoped>
.m-auto {
  margin: 0 auto;
}
</style>
