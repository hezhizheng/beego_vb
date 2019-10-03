import request from '@/utils/request'
import api from './API_LIST.js'

export function index(params) {
  return request({
    url: api.DASHBOARDS,
    method: 'get',
    params
  })
}
