<script setup lang="ts">
import { useAuth } from '@/composables/useAuth'
definePageMeta({
  layout: 'signin'
})
const router = useRouter()
const { signin, me } = useAuth()
const isLoading = ref(false)
const email = ref('')
const password = ref('')

const userLogin = async () => {
  isLoading.value = true
  const { error } = await signin(email.value, password.value)
  isLoading.value = false
  if (error.value) {
    alert(error.value.message)
  } else {
    await me()
    router.push('/home')
  }
}
</script>

<template>
  <div>
    <div>
      <label for="email">Email:</label>
      <input type="email" id="email" v-model="email" />
    </div>
    <div>
      <label for="password">Password:</label>
      <input type="password" id="password" v-model="password" />
    </div>
    <button @click="userLogin">ログイン</button>
  </div>
  <Loading v-if="isLoading" />
</template>
