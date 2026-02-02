import request from '@/utils/request'

export function getAlerts() {
  return request({
    url: '/alerts',
    method: 'get'
  })
}

// Alert Rules
export function getAlertRules() {
  return request({
    url: '/alert-rules',
    method: 'get'
  })
}

export function createAlertRule(data) {
  return request({
    url: '/alert-rules',
    method: 'post',
    data
  })
}

export function deleteAlertRule(id) {
  return request({
    url: `/alert-rules/${id}`,
    method: 'delete'
  })
}

export function updateAlertRule(id, data) {
  return request({
    url: `/alert-rules/${id}`,
    method: 'put',
    data
  })
}

// Alert Channels
export function getAlertChannels() {
  return request({
    url: '/alert-channels',
    method: 'get'
  })
}

export function createAlertChannel(data) {
  return request({
    url: '/alert-channels',
    method: 'post',
    data
  })
}

export function deleteAlertChannel(id) {
  return request({
    url: `/alert-channels/${id}`,
    method: 'delete'
  })
}

export function updateAlertChannel(id, data) {
  return request({
    url: `/alert-channels/${id}`,
    method: 'put',
    data
  })
}
