import useCustomFetch from '@/api/apiClient'

export const useAuth = () => {
  const currentUser = useState<{
    name: string | undefined
    email: string | undefined
  }>('currentUser', () => ({
    name: undefined,
    email: undefined
  }))

  const signin = (email: string, password: string) => {
    const { data, error, pending } = useCustomFetch('signin', {
      method: 'POST',
      body: { email, password },
      default: () => ({ token: '' })
    })
    watchEffect(() => {
      currentUser.value.email = data.value?.token
      currentUser.value.name = data.value?.token
      console.log(currentUser.value)
    })
    console.log(data)
    return { currentUser, error, pending }
  }

  const signout = () => async () => {
    currentUser.value = { name: undefined, email: undefined }
  }
  return {
    currentUser,
    signin,
    signout
  }
}
