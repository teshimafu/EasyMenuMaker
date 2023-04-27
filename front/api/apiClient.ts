export const useCustomFetch: typeof useFetch = (path, options) => {
  const config = useRuntimeConfig()
  return useFetch(path, {
    baseURL: config.public.baseURL,
    ...options,
    onRequest({ request, options }) {
      options.headers = options.headers || {}
      const token = localStorage.getItem('accessToken')
      options.headers = Object.assign(options.headers, {
        'Content-Type': 'application/json'
      })
      // トークンが存在する場合、Authorizationヘッダーを設定する
      if (token) {
        options.headers = Object.assign(options.headers, {
          Authorization: `Bearer ${token}`
        })
      }
    },
    async onResponseError({ response }) {
      console.error(response.status, response.statusText)
      if (response.status === 401) {
        localStorage.setItem('accessToken', '')
      }
    }
  })
}

export default useCustomFetch
