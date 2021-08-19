import service from '@/utils/request'

export const createCourse = (data) => {
    return service({
        url: '/course/createCourse',
        method: 'post',
        data
    })
}

export const deleteCourse = (data) => {
    return service({
        url: '/course/deleteCourse',
        method: 'delete',
        data
    })
}

export const deleteCourseByIds = (data) => {
    return service({
        url: '/course/deleteCourseByIds',
        method: 'delete',
        data
    })
}

export const updateCourse = (data) => {
    return service({
        url: '/course/updateCourse',
        method: 'put',
        data
    })
}

export const findCourse = (params) => {
    return service({
        url: '/course/findCourse',
        method: 'get',
        params
    })
}

export const getCourseList = (params) => {
    return service({
        url: '/course/getCourseList',
        method: 'get',
        params
    })
}

export const getClassListFromCourse = (params) => {
    return service({
        url: '/course/getClassListFromCourse',
        method: 'get',
        params
    })
}

