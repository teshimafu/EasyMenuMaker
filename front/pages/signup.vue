<template>
  <div>
    <h1>Signup</h1>
    <form @submit.prevent="signup">
      <div>
        <label for="name">Name:</label>
        <input type="text" id="name" v-model="name" />
      </div>
      <div>
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="email" />
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" />
      </div>
      <button type="submit">Signup</button>
    </form>
    <div>
      <p>{{ errorMessage }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import useCustomFetch from '@/api/apiClient'

const name = ref('')
const email = ref('')
const password = ref('')
const errorMessage = ref('')

const router = useRouter()

const signup = async () => {
  const response = await useCustomFetch('signup', {
    method: 'POST',
    body: {
      name: name.value,
      email: email.value,
      password: password.value
    }
  })
  if (response.data.value) {
    router.push('/home')
    return
  }
  if (response.error.value) {
    errorMessage.value = response.error.value.data.message
    return
  }
}
</script>
