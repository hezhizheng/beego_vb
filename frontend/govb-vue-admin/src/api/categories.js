import request from '@/utils/request'
import api from './API_LIST.js'

export function getCategories(params) {
  return request({
    url: api.CATEGORIES,
    method: 'get',
    params
  })
}

export function index(params) {
  return request({
    url: api.CATEGORIES_LIST,
    method: 'get',
    params
  })
}

export function store(params) {
  return request({
    url: api.CATEGORIES_LIST,
    method: 'post',
    data : params
  })
}

export function update(params) {
  return request({
    url: api.CATEGORIES_LIST + '/' + params.id,
    method: 'put',
    data : params
  })
}

export function destroy(params) {
  return request({
    url: api.CATEGORIES_LIST + '/' + params.id,
    method: 'delete',
  })
}
