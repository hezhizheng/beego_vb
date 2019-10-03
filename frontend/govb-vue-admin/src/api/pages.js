import request from '@/utils/request'
import api from './API_LIST.js'

export function index(params) {
  return request({
    url: api.PAGES,
    method: 'get',
    params
  })
}

export function store(params) {
  return request({
    url: api.PAGES,
    method: 'post',
    data : params
  })
}

export function destroy(params) {
  return request({
    url: api.PAGES + '/' + params.id,
    method: 'delete',
  })
}

export function update(params) {
  return request({
    url: api.PAGES + '/' + params.id,
    method: 'put',
    data : params
  })
}
