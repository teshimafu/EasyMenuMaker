import { useAuth } from '@/composables/useAuth'

export default defineNuxtRouteMiddleware((to, from) => {
  const { currentUser } = useAuth()
  if (!currentUser.value && to.path !== '/signin') {
    const path = '/signin'
    return { path }
  }
})
