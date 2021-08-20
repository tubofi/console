import service from '@/utils/request'

// @Tags Teacher
// @Summary 创建Teacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Teacher true "创建Teacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teacher/createTeacher [post]
export const createTeacher = (data) => {
    return service({
        url: '/teacher/createTeacher',
        method: 'post',
        data
    })
}

// @Tags Teacher
// @Summary 删除Teacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Teacher true "删除Teacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teacher/deleteTeacher [delete]
export const deleteTeacher = (data) => {
    return service({
        url: '/teacher/deleteTeacher',
        method: 'delete',
        data
    })
}

// @Tags Teacher
// @Summary 删除Teacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Teacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teacher/deleteTeacher [delete]
export const deleteTeacherByIds = (data) => {
    return service({
        url: '/teacher/deleteTeacherByIds',
        method: 'delete',
        data
    })
}

// @Tags Teacher
// @Summary 更新Teacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Teacher true "更新Teacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teacher/updateTeacher [put]
export const updateTeacher = (data) => {
    return service({
        url: '/teacher/updateTeacher',
        method: 'put',
        data
    })
}

// @Tags Teacher
// @Summary 用id查询Teacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Teacher true "用id查询Teacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teacher/findTeacher [get]
export const findTeacher = (params) => {
    return service({
        url: '/teacher/findTeacher',
        method: 'get',
        params
    })
}

// @Tags Teacher
// @Summary 分页获取Teacher列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Teacher列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teacher/getTeacherList [get]
export const getTeacherList = (params) => {
    return service({
        url: '/teacher/getTeacherList',
        method: 'get',
        params
    })
}