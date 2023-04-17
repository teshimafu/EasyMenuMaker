import axios, { AxiosError } from 'axios'

export const signup = async (name: string, email: string, password: string) => {
  try {
    const response = await axios.post('http://localhost:8080/signup', {
      name: name,
      email: email,
      password: password
    })

    switch (response.status) {
      case 201:
        return response.data
      default:
        return
    }
  } catch (error) {
    if (axios.isAxiosError(error)) {
      const axiosError = error as AxiosError<{ message: string }>
      console.error('Error:', axiosError.response?.data.message)
    } else {
      console.error('Unknown error:', error)
    }
  }
}
