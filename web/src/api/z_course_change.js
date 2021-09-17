
import service from '@/utils/request'

// @Tags CourseChange
// @Summary 创建CourseChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseChange true "创建CourseChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /courseChange/createCourseChange [post]
export const createCourseChange = (data) => {
    return service({
        url: '/courseChange/createCourseChange',
        method: 'post',
        data
    })
}

// @Tags CourseChange
// @Summary 删除CourseChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseChange true "删除CourseChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /courseChange/deleteCourseChange [delete]
export const deleteCourseChange = (data) => {
    return service({
        url: '/courseChange/deleteCourseChange',
        method: 'delete',
        data
    })
}

// @Tags CourseChange
// @Summary 删除CourseChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CourseChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /courseChange/deleteCourseChange [delete]
export const deleteCourseChangeByIds = (data) => {
    return service({
        url: '/courseChange/deleteCourseChangeByIds',
        method: 'delete',
        data
    })
}

// @Tags CourseChange
// @Summary 更新CourseChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseChange true "更新CourseChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /courseChange/updateCourseChange [put]
export const updateCourseChange = (data) => {
    return service({
        url: '/courseChange/updateCourseChange',
        method: 'put',
        data
    })
}

// @Tags CourseChange
// @Summary 用id查询CourseChange
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CourseChange true "用id查询CourseChange"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /courseChange/findCourseChange [get]
export const findCourseChange = (params) => {
    return service({
        url: '/courseChange/findCourseChange',
        method: 'get',
        params
    })
}

// @Tags CourseChange
// @Summary 分页获取CourseChange列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取CourseChange列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /courseChange/getCourseChangeList [get]
export const getCourseChangeList = (params) => {
    return service({
        url: '/courseChange/getCourseChangeList',
        method: 'get',
        params
    })
}