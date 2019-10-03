import request from '@/utils/request'
import api from './API_LIST.js'

export function login(email, password) {
  return request({
    url: api.LOGIN,
    method: 'post',
    data: {
      email,
      password
    }
  })
}

export function getInfo(token) {
  return request({
    url: api.ME,
    method: 'get',
    // params: { token }
  })
}

export function logout() {
  return request({
    url: api.LOGOUT,
    method: 'post'
  })
}
