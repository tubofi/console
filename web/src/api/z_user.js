import service from '@/utils/request'

export const getTmpSecret = (params) => {
    return service({
        url: '/cos/getTmpSecret',
        method: 'get',
        params
    })
}