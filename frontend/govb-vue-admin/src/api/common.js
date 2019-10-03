import request from '@/utils/request'

export function mavonUpload(params) {
  return request({
    url: '/sm.ms/api/upload',
    method: 'post',
    data: params,
  })
}
