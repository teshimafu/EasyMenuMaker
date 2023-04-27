import useCustomFetch from '@/api/apiClient'

export const useAuth = () => {
  const currentUser = useState<{
    name: string | undefined
    email: string | undefined
  }>('currentUser', () => ({
    name: undefined,
    email: undefined
  }))

  const signup = async (args: {
    name: string
    email: string
    password: string
  }) => {
    return useCustomFetch('signup', {
      method: 'POST',
      body: {
        name: args.name,
        email: args.email,
        password: args.password
      }
    })
  }

  const signin = async (email: string, password: string) => {
    const { data, error } = await useCustomFetch('signin', {
      method: 'POST',
      body: { email, password },
      default: () => ({ token: '' })
    })
    if (data.value?.token) {
      localStorage.setItem('accessToken', data.value.token)
    }
    return { currentUser, error }
  }

  const me = async () => {
    const { data, error } = await useCustomFetch('me', {
      method: 'GET',
      default: () => ({ id: '', name: '', email: '' })
    })
    if (data.value?.id) {
      currentUser.value.email = data.value.email
      currentUser.value.name = data.value.name
    }
    return { data, error }
  }

  const signout = () => async () => {
    localStorage.setItem('accessToken', '')
    currentUser.value = { name: undefined, email: undefined }
  }
  return {
    currentUser,
    signup,
    signin,
    me,
    signout
  }
}
