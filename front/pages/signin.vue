<script setup lang="ts">
definePageMeta({
  layout: 'signin'
})
const router = useRouter()
const { signin } = useAuth()
const isLoading = ref(false)
const email = ref('')
const password = ref('')

const userLogin = async () => {
  const { pending, error } = signin(email.value, password.value)
  watchEffect(() => {
    isLoading.value = pending.value
    if (!pending.value) {
      if (error.value) {
        isLoading.value = false
        alert(error.value.message)
      } else {
        router.push('/home')
      }
    }
  })
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
