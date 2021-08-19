import service from '@/utils/request'

export const createStudent = (data) => {
    return service({
        url: '/student/createStudent',
        method: 'post',
        data
    })
}

export const deleteStudent = (data) => {
    return service({
        url: '/student/deleteStudent',
        method: 'delete',
        data
    })
}

export const updateStudent = (data) => {
    return service({
        url: '/student/updateStudent',
        method: 'put',
        data
    })
}

export const findStudent = (params) => {
    return service({
        url: '/student/findStudent',
        method: 'get',
        params
    })
}

export const getStudentList = (params) => {
    return service({
        url: '/student/getStudentList',
        method: 'get',
        params
    })
}

export const getTrialStudentList = (params) => {
    return service({
        url: '/student/getTrialStudentList',
        method: 'get',
        params
    })
}

export const turnStudent = (data) => {
    return service({
        url: '/student/turnStudent',
        method: 'put',
        data
    })
}

