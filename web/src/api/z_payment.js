import service from '@/utils/request'

// @Tags Payment
// @Summary 创建Payment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Payment true "创建Payment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payment/createPayment [post]
export const createPayment = (data) => {
    return service({
        url: '/bill/createPayment',
        method: 'post',
        data
    })
}

// @Tags Payment
// @Summary 删除Payment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Payment true "删除Payment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payment/deletePayment [delete]
export const deletePayment = (data) => {
    return service({
        url: '/bill/deletePayment',
        method: 'delete',
        data
    })
}

// @Tags Payment
// @Summary 删除Payment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Payment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payment/deletePayment [delete]
export const deletePaymentByIds = (data) => {
    return service({
        url: '/bill/deletePaymentByIds',
        method: 'delete',
        data
    })
}

// @Tags Payment
// @Summary 更新Payment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Payment true "更新Payment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payment/updatePayment [put]
export const updatePayment = (data) => {
    return service({
        url: '/bill/updatePayment',
        method: 'put',
        data
    })
}

// @Tags Payment
// @Summary 用id查询Payment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Payment true "用id查询Payment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payment/findPayment [get]
export const findPayment = (params) => {
    return service({
        url: '/bill/findPayment',
        method: 'get',
        params
    })
}

// @Tags Payment
// @Summary 分页获取Payment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Payment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payment/getPaymentList [get]
export const getPaymentList = (params) => {
    return service({
        url: '/bill/getPaymentList',
        method: 'get',
        params
    })
}

export const getAllStudents = (params) => {
    return service({
        url: '/bill/getAllStudents',
        method: 'get',
        params
    })
}

