import service from '@/utils/request'

// @Tags LivingBill
// @Summary 创建LivingBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingBill true "创建LivingBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /livingBill/createLivingBill [post]
export const createLivingBill = (data) => {
    return service({
        url: '/living/createLivingBill',
        method: 'post',
        data
    })
};

// @Tags LivingBill
// @Summary 删除LivingBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingBill true "删除LivingBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /livingBill/deleteLivingBill [delete]
export const deleteLivingBill = (data) => {
    return service({
        url: '/living/deleteLivingBill',
        method: 'delete',
        data
    })
};

// @Tags LivingBill
// @Summary 删除LivingBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LivingBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /livingBill/deleteLivingBill [delete]
export const deleteLivingBillByIds = (data) => {
    return service({
        url: '/living/deleteLivingBillByIds',
        method: 'delete',
        data
    })
};

// @Tags LivingBill
// @Summary 更新LivingBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingBill true "更新LivingBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /livingBill/updateLivingBill [put]
export const updateLivingBill = (data) => {
    return service({
        url: '/living/updateLivingBill',
        method: 'put',
        data
    })
};

// @Tags LivingBill
// @Summary 用id查询LivingBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingBill true "用id查询LivingBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /livingBill/findLivingBill [get]
export const findLivingBill = (params) => {
    return service({
        url: '/living/findLivingBill',
        method: 'get',
        params
    })
};

// @Tags LivingBill
// @Summary 分页获取LivingBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取LivingBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /livingBill/getLivingBillList [get]
export const getLivingBillList = (params) => {
    return service({
        url: '/living/getLivingBillList',
        method: 'get',
        params
    })
};

// @Tags LivingCost
// @Summary 创建LivingCost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingCost true "创建LivingCost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /livingCost/createLivingCost [post]
export const createLivingCost = (data) => {
    return service({
        url: '/living/createLivingCost',
        method: 'post',
        data
    })
};

// @Tags LivingCost
// @Summary 删除LivingCost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingCost true "删除LivingCost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /livingCost/deleteLivingCost [delete]
export const deleteLivingCost = (data) => {
    return service({
        url: '/living/deleteLivingCost',
        method: 'delete',
        data
    })
};

// @Tags LivingCost
// @Summary 删除LivingCost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LivingCost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /livingCost/deleteLivingCost [delete]
export const deleteLivingCostByIds = (data) => {
    return service({
        url: '/living/deleteLivingCostByIds',
        method: 'delete',
        data
    })
};

// @Tags LivingCost
// @Summary 更新LivingCost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingCost true "更新LivingCost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /livingCost/updateLivingCost [put]
export const updateLivingCost = (data) => {
    return service({
        url: '/living/updateLivingCost',
        method: 'put',
        data
    })
};

// @Tags LivingCost
// @Summary 用id查询LivingCost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LivingCost true "用id查询LivingCost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /livingCost/findLivingCost [get]
export const findLivingCost = (params) => {
    return service({
        url: '/living/findLivingCost',
        method: 'get',
        params
    })
};

// @Tags LivingCost
// @Summary 分页获取LivingCost列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取LivingCost列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /livingCost/getLivingCostList [get]
export const getLivingCostList = (params) => {
    return service({
        url: '/living/getLivingCostList',
        method: 'get',
        params
    })
};