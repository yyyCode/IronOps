import request from '@/utils/request'

export function getAuditLogs() {
  return request({
    url: '/audits',
    method: 'get'
  })
}
