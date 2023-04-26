export const useCustomFetch: typeof useFetch = (path, options) => {
  const config = useRuntimeConfig()
  // トークンをローカルストレージから取得する（または他の方法で取得）
  const token = localStorage.getItem('accessToken')
  if (options === undefined) {
    options = {}
  }
  if (options?.headers === undefined) {
    options.headers = new Headers()
  }
  options.headers = Object.assign(options.headers, {
    'Content-Type': 'application/json'
  })
  // トークンが存在する場合、Authorizationヘッダーを設定する
  if (token) {
    options.headers = Object.assign(options.headers, {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`
    })
  }
  return useFetch(path, {
    baseURL: config.public.baseURL,
    ...options,
    async onResponseError({ response }) {
      console.error(response.status, response.statusText)
      if (response.status === 401) {
        localStorage.setItem('accessToken', '')
      }
    }
  })
}

export default useCustomFetch
