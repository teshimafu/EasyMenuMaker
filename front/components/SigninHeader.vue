<template>
  <div>
    <button v-if="showSignupButton" v-on:click="signup">サインアップ</button>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
const router = useRouter()
const route = useRoute()

const showSignupButton = computed(() => route.path !== '/signup')

const signup = () => {
  router.push('/signup')
}
</script>
