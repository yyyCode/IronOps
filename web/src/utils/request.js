import axios from 'axios'
import { ElMessage } from 'element-plus'

const service = axios.create({
  baseURL: '/api/v1',
  timeout: 5000
})

// Request interceptor
service.interceptors.request.use(
  config => {
    // Add logic here to add auth token if available
    // const token = localStorage.getItem('token')
    // if (token) {
    //   config.headers['X-User'] = token
    // }
    // For MVP, we simulate a logged-in user
    config.headers['X-User'] = 'admin' 
    return config
  },
  error => {
    console.log(error)
    return Promise.reject(error)
  }
)

// Response interceptor
service.interceptors.response.use(
  response => {
    const res = response.data
    // You can check custom error codes here if needed
    return res
  },
  error => {
    console.log('err' + error)
    ElMessage({
      message: error.message || 'Request Error',
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
