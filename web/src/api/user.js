import request from '@/utils/request'

export function getUsers() {
  return request({
    url: '/users',
    method: 'get'
  })
}

export function getRoles() {
  return request({
    url: '/roles',
    method: 'get'
  })
}

// Placeholder for future use
export function createUser(data) {
    return request({
        url: '/register',
        method: 'post',
        data
    })
}
