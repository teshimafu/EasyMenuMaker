<template>
  <Loading v-if="pending" />
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import useCustomFetch from '@/api/apiClient'

const { pending, error } = useCustomFetch('users', {
  method: 'GET'
})
// if (error.value) {
//   const router = useRouter()
//   router.push('/signup')
// }
</script>
