import service from '@/utils/request'

export const getProfileClass = (params) => {
    return service({
        url: '/class/getProfileClass',
        method: 'get',
        params
    })
}