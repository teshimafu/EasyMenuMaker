export default defineNuxtConfig({
  runtimeConfig: {
    public: {
      baseURL: process.env.API_SERVER_URL ?? 'http://localhost:8080/'
    }
  }
})
