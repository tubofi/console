import service from '@/utils/request'

export const getTmpSecret = (params) => {
    return service({
        url: '/upload/getTmpSecret',
        method: 'get',
        params
    })
};

export const getImageTmpSecret = (params) => {
    return service({
        url: '/upload/getImageTmpSecret',
        method: 'get',
        params
    })
};

// @Tags Upload
// @Summary 创建Upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Upload true "创建Upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /upload/createUpload [post]
export const createUpload = (data) => {
    return service({
        url: '/upload/createUpload',
        method: 'post',
        data
    })
};

// @Tags Upload
// @Summary 删除Upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Upload true "删除Upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /upload/deleteUpload [delete]
export const deleteUpload = (data) => {
    return service({
        url: '/upload/deleteUpload',
        method: 'delete',
        data
    })
};

// @Tags Upload
// @Summary 删除Upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /upload/deleteUpload [delete]
export const deleteUploadByIds = (data) => {
    return service({
        url: '/upload/deleteUploadByIds',
        method: 'delete',
        data
    })
};

// @Tags Upload
// @Summary 更新Upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Upload true "更新Upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /upload/updateUpload [put]
export const updateUpload = (data) => {
    return service({
        url: '/upload/updateUpload',
        method: 'put',
        data
    })
};

// @Tags Upload
// @Summary 用id查询Upload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Upload true "用id查询Upload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /upload/findUpload [get]
export const findUpload = (params) => {
    return service({
        url: '/upload/findUpload',
        method: 'get',
        params
    })
};

// @Tags Upload
// @Summary 分页获取Upload列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Upload列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /upload/getUploadList [get]
export const getUploadList = (params) => {
    return service({
        url: '/upload/getUploadList',
        method: 'get',
        params
    })
};