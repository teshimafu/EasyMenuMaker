import axios from 'axios'

const config = useRuntimeConfig()

console.log(config)
const apiClient = axios.create({
  baseURL: config.public.baseURL,
  headers: {
    'Content-Type': 'application/json'
  }
})

apiClient.interceptors.request.use(config => {
  // トークンをローカルストレージから取得する（または他の方法で取得）
  const token = localStorage.getItem('accessToken')

  // トークンが存在する場合、Authorizationヘッダーを設定する
  if (token) {
    config.headers['Authorization'] = `Bearer ${token}`
  }

  return config
})

export default apiClient
