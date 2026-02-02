import request from '@/utils/request'

export function getServices() {
  return request({
    url: '/services',
    method: 'get'
  })
}

export function createService(data) {
  return request({
    url: '/services',
    method: 'post',
    data
  })
}

export function addInstance(data) {
  return request({
    url: '/instances',
    method: 'post',
    data
  })
}

export function controlInstance(id, action) {
  return request({
    url: `/instances/${id}/control`,
    method: 'post',
    data: { action }
  })
}
