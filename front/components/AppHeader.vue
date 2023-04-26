<template>
  <div>
    <label for="name">Name:{{ name }}</label>
  </div>
</template>

<script setup lang="ts">
interface Props {
  name: string
}

defineProps<Props>()
</script>
