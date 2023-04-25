<script setup lang="ts">
const { currentUser, signin } = useAuth()
let isLoading = ref(false)
const email = ref('')
const password = ref('')

const userLogin = async () => {
  isLoading.value = true
  const { pending } = signin(email.value, password.value)
  watchEffect(() => {
    if (!pending.value) {
      isLoading.value = false
    }
  })
  console.log(pending.value)
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
