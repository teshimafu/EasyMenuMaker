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
import axios, { AxiosError } from 'axios'
import apiClient from '@/api/apiClient'

const name = ref('')
const email = ref('')
const password = ref('')
const errorMessage = ref('')

const router = useRouter()

async function signup() {
  try {
    const response = await apiClient.post('/signup', {
      name: name.value,
      email: email.value,
      password: password.value
    })

    switch (response.status) {
      case 201:
        router.push('/home')
        return
      default:
        errorMessage.value = 'Unknown error'
        return
    }
  } catch (error) {
    if (axios.isAxiosError(error)) {
      const axiosError = error as AxiosError<{ message: string }>
      console.error('Error:', axiosError.response?.data.message)
      errorMessage.value = axiosError.response?.data.message ?? 'Unknown error'
    } else {
      console.error('Unknown error:', error)
    }
  }
}
</script>
