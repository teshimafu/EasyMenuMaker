<template>
  <div>
    <label for="name">Name:{{ currentUser.name }}</label>
  </div>
</template>

<script setup lang="ts">
import { useAuth } from '@/composables/useAuth'

const { currentUser } = useAuth()
</script>
