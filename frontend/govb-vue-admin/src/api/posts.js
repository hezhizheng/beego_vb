import request from '@/utils/request'
import api from './API_LIST.js'

export function getPosts(params) {
  return request({
    url: api.POSTS,
    method: 'get',
    params
  })
}

export function delPosts(params) {
  return request({
    url: api.POSTS + '/' + params.id,
    method: 'delete',
  })
}

export function editPosts(params) {
  return request({
    url: api.POSTS + '/' + params.id + '/edit',
    method: 'get',
  })
}

export function storePosts(params) {
  return request({
    url: api.POSTS,
    method: 'post',
    data : params
  })
}

export function updatePosts(params) {
  return request({
    url: api.POSTS + '/' + params.id,
    method: 'put',
    data : params
  })
}
