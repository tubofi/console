import service from '@/utils/request'

// @Tags Wage
// @Summary 创建Wage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wage true "创建Wage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wage/createWage [post]
export const createWage = (data) => {
    return service({
        url: '/bill/createWage',
        method: 'post',
        data
    })
}

// @Tags Wage
// @Summary 删除Wage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wage true "删除Wage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wage/deleteWage [delete]
export const deleteWage = (data) => {
    return service({
        url: '/bill/deleteWage',
        method: 'delete',
        data
    })
}

// @Tags Wage
// @Summary 删除Wage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Wage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wage/deleteWage [delete]
export const deleteWageByIds = (data) => {
    return service({
        url: '/bill/deleteWageByIds',
        method: 'delete',
        data
    })
}

// @Tags Wage
// @Summary 更新Wage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wage true "更新Wage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wage/updateWage [put]
export const updateWage = (data) => {
    return service({
        url: '/bill/updateWage',
        method: 'put',
        data
    })
}

// @Tags Wage
// @Summary 用id查询Wage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Wage true "用id查询Wage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wage/findWage [get]
export const findWage = (params) => {
    return service({
        url: '/bill/findWage',
        method: 'get',
        params
    })
}

// @Tags Wage
// @Summary 分页获取Wage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Wage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wage/getWageList [get]
export const getWageList = (params) => {
    return service({
        url: '/bill/getWageList',
        method: 'get',
        params
    })
}