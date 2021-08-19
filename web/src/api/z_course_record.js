import service from '@/utils/request'

// @Tags CourseRecord
// @Summary 创建CourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseRecord true "创建CourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /courseRecord/createCourseRecord [post]
export const createCourseRecord = (data) => {
    return service({
        url: '/class/createCourseRecord',
        method: 'post',
        data
    })
}

// @Tags CourseRecord
// @Summary 删除CourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseRecord true "删除CourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /courseRecord/deleteCourseRecord [delete]
export const deleteCourseRecord = (data) => {
    return service({
        url: '/class/deleteCourseRecord',
        method: 'delete',
        data
    })
}

// @Tags CourseRecord
// @Summary 删除CourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /courseRecord/deleteCourseRecord [delete]
export const deleteCourseRecordByIds = (data) => {
    return service({
        url: '/class/deleteCourseRecordByIds',
        method: 'delete',
        data
    })
}

// @Tags CourseRecord
// @Summary 更新CourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseRecord true "更新CourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /courseRecord/updateCourseRecord [put]
export const updateCourseRecord = (data) => {
    return service({
        url: '/class/updateCourseRecord',
        method: 'put',
        data
    })
}

// @Tags CourseRecord
// @Summary 用id查询CourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseRecord true "用id查询CourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /courseRecord/findCourseRecord [get]
export const findCourseRecord = (params) => {
    return service({
        url: '/class/findCourseRecord',
        method: 'get',
        params
    })
}

// @Tags CourseRecord
// @Summary 分页获取CourseRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取CourseRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /courseRecord/getCourseRecordList [get]
export const getCourseRecordList = (params) => {
    return service({
        url: '/class/getCourseRecordList',
        method: 'get',
        params
    })
}

export const getCramCourseRecordList = (params) => {
    return service({
        url: '/class/getCramCourseRecordList',
        method: 'get',
        params
    })
}

export const cramCourseRecord = (data) => {
    return service({
        url: '/class/cramCourseRecord',
        method: 'put',
        data
    })
}

export const ignoreCramCourseRecord = (data) => {
    return service({
        url: '/class/ignoreCramCourseRecord',
        method: 'put',
        data
    })
}

export const getFeedbackCourseRecordList = (params) => {
    return service({
        url: '/class/getFeedbackCourseRecordList',
        method: 'get',
        params
    })
}

export const feedbackCourseRecord = (data) => {
    return service({
        url: '/class/feedbackCourseRecord',
        method: 'put',
        data
    })
}