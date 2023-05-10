import vuetify from 'vite-plugin-vuetify'

export default defineNuxtConfig({
  runtimeConfig: {
    public: {
      baseURL: process.env.API_SERVER_URL ?? 'http://localhost:8080/'
    }
  },
  build: {
    transpile: ['vuetify']
  },
  hooks: {
    'vite:extendConfig': config => {
      config.plugins!.push(vuetify())
    }
  },
  vite: {
    ssr: {
      noExternal: ['vuetify']
    },
    define: {
      'process.env.DEBUG': false
    }
  },
  css: ['@/assets/main.scss']
})
