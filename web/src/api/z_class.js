import service from '@/utils/request'

export const createClass = (data) => {
    return service({
        url: '/class/createClass',
        method: 'post',
        data
    })
}

export const deleteClass = (data) => {
    return service({
        url: '/class/deleteClass',
        method: 'delete',
        data
    })
}

export const updateClass = (data) => {
    return service({
        url: '/class/updateClass',
        method: 'put',
        data
    })
}

export const findClass = (params) => {
    return service({
        url: '/class/findClass',
        method: 'get',
        params
    })
}

export const getClassList = (params) => {
    return service({
        url: '/class/getClassList',
        method: 'get',
        params
    })
}

export const getIdleStudents = (params) => {
    return service({
        url: '/class/getIdleStudents',
        method: 'get',
        params
    })
}