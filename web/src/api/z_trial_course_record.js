import service from '@/utils/request'

// @Tags TrialCourseRecord
// @Summary 创建TrialCourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TrialCourseRecord true "创建TrialCourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /trialCourseRecord/createTrialCourseRecord [post]
export const createTrialCourseRecord = (data) => {
    return service({
        url: '/trial/createTrialCourseRecord',
        method: 'post',
        data
    })
};

// @Tags TrialCourseRecord
// @Summary 删除TrialCourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TrialCourseRecord true "删除TrialCourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /trialCourseRecord/deleteTrialCourseRecord [delete]
export const deleteTrialCourseRecord = (data) => {
    return service({
        url: '/trial/deleteTrialCourseRecord',
        method: 'delete',
        data
    })
};

// @Tags TrialCourseRecord
// @Summary 删除TrialCourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TrialCourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /trialCourseRecord/deleteTrialCourseRecord [delete]
export const deleteTrialCourseRecordByIds = (data) => {
    return service({
        url: '/trial/deleteTrialCourseRecordByIds',
        method: 'delete',
        data
    })
};

// @Tags TrialCourseRecord
// @Summary 更新TrialCourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TrialCourseRecord true "更新TrialCourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /trialCourseRecord/updateTrialCourseRecord [put]
export const updateTrialCourseRecord = (data) => {
    return service({
        url: '/trial/updateTrialCourseRecord',
        method: 'put',
        data
    })
};

// @Tags TrialCourseRecord
// @Summary 用id查询TrialCourseRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TrialCourseRecord true "用id查询TrialCourseRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /trialCourseRecord/findTrialCourseRecord [get]
export const findTrialCourseRecord = (params) => {
    return service({
        url: '/trial/findTrialCourseRecord',
        method: 'get',
        params
    })
};

// @Tags TrialCourseRecord
// @Summary 分页获取TrialCourseRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TrialCourseRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /trialCourseRecord/getTrialCourseRecordList [get]
export const getTrialCourseRecordList = (params) => {
    return service({
        url: '/trial/getTrialCourseRecordList',
        method: 'get',
        params
    })
};

export const getAllTrialStudents = (params) => {
    return service({
        url: '/trial/getAllTrialStudents',
        method: 'get',
        params
    })
};

