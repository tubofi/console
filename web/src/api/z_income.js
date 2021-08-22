import service from '@/utils/request'

// @Tags Income
// @Summary 创建Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Income true "创建Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /income/createIncome [post]
export const createIncome = (data) => {
    return service({
        url: '/bill/createIncome',
        method: 'post',
        data
    })
}

// @Tags Income
// @Summary 删除Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Income true "删除Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /income/deleteIncome [delete]
export const deleteIncome = (data) => {
    return service({
        url: '/bill/deleteIncome',
        method: 'delete',
        data
    })
}

// @Tags Income
// @Summary 删除Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /income/deleteIncome [delete]
export const deleteIncomeByIds = (data) => {
    return service({
        url: '/bill/deleteIncomeByIds',
        method: 'delete',
        data
    })
}

// @Tags Income
// @Summary 更新Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Income true "更新Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /income/updateIncome [put]
export const updateIncome = (data) => {
    return service({
        url: '/bill/updateIncome',
        method: 'put',
        data
    })
}

// @Tags Income
// @Summary 用id查询Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Income true "用id查询Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /income/findIncome [get]
export const findIncome = (params) => {
    return service({
        url: '/bill/findIncome',
        method: 'get',
        params
    })
}

// @Tags Income
// @Summary 分页获取Income列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Income列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /income/getIncomeList [get]
export const getIncomeList = (params) => {
    return service({
        url: '/bill/getIncomeList',
        method: 'get',
        params
    })
}