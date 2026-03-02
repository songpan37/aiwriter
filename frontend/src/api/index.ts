import axios from 'axios'

const MOCK_MODE = true

const mockResponses: Record<string, {code: number, message: string, data?: any}> = {
  '/auth/login': { code: 0, message: 'success', data: { token: 'mock-token-123', user: { id: 1, username: 'testuser', email: 'test@example.com', avatar: '' } } },
  '/auth/register': { code: 0, message: 'success' },
  '/works': { code: 0, data: { list: [], pagination: { page: 1, pageSize: 10, total: 0, totalPages: 0 } } },
}

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api/v1',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
  },
})

api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

api.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

const originalGet = api.get.bind(api)
const originalPost = api.post.bind(api)

api.get = async (url: string, config?: any) => {
  if (MOCK_MODE && mockResponses[url]) {
    return mockResponses[url]
  }
  return originalGet(url, config)
}

api.post = async (url: string, data?: any, config?: any) => {
  if (MOCK_MODE && mockResponses[url]) {
    return mockResponses[url]
  }
  return originalPost(url, data, config)
}

export default api
