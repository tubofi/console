import service from '@/utils/request'

// @Tags Bill
// @Summary 创建Bill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Bill true "创建Bill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bill/createBill [post]
export const createBill = (data) => {
    return service({
        url: '/bill/createBill',
        method: 'post',
        data
    })
}

// @Tags Bill
// @Summary 删除Bill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Bill true "删除Bill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bill/deleteBill [delete]
export const deleteBill = (data) => {
    return service({
        url: '/bill/deleteBill',
        method: 'delete',
        data
    })
}

// @Tags Bill
// @Summary 删除Bill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Bill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bill/deleteBill [delete]
export const deleteBillByIds = (data) => {
    return service({
        url: '/bill/deleteBillByIds',
        method: 'delete',
        data
    })
}

// @Tags Bill
// @Summary 更新Bill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Bill true "更新Bill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bill/updateBill [put]
export const updateBill = (data) => {
    return service({
        url: '/bill/updateBill',
        method: 'put',
        data
    })
}

// @Tags Bill
// @Summary 用id查询Bill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Bill true "用id查询Bill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bill/findBill [get]
export const findBill = (params) => {
    return service({
        url: '/bill/findBill',
        method: 'get',
        params
    })
}

// @Tags Bill
// @Summary 分页获取Bill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Bill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bill/getBillList [get]
export const getBillList = (params) => {
    return service({
        url: '/bill/getBillList',
        method: 'get',
        params
    })
}