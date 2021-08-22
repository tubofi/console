import service from '@/utils/request'

// @Tags Outcome
// @Summary 创建Outcome
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Outcome true "创建Outcome"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /outcome/createOutcome [post]
export const createOutcome = (data) => {
    return service({
        url: '/bill/createOutcome',
        method: 'post',
        data
    })
}

// @Tags Outcome
// @Summary 删除Outcome
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Outcome true "删除Outcome"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /outcome/deleteOutcome [delete]
export const deleteOutcome = (data) => {
    return service({
        url: '/bill/deleteOutcome',
        method: 'delete',
        data
    })
}

// @Tags Outcome
// @Summary 删除Outcome
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Outcome"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /outcome/deleteOutcome [delete]
export const deleteOutcomeByIds = (data) => {
    return service({
        url: '/bill/deleteOutcomeByIds',
        method: 'delete',
        data
    })
}

// @Tags Outcome
// @Summary 更新Outcome
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Outcome true "更新Outcome"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /outcome/updateOutcome [put]
export const updateOutcome = (data) => {
    return service({
        url: '/bill/updateOutcome',
        method: 'put',
        data
    })
}

// @Tags Outcome
// @Summary 用id查询Outcome
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Outcome true "用id查询Outcome"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /outcome/findOutcome [get]
export const findOutcome = (params) => {
    return service({
        url: '/bill/findOutcome',
        method: 'get',
        params
    })
}

// @Tags Outcome
// @Summary 分页获取Outcome列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Outcome列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /outcome/getOutcomeList [get]
export const getOutcomeList = (params) => {
    return service({
        url: '/bill/getOutcomeList',
        method: 'get',
        params
    })
}