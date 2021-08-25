package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"time"
)

func GetTmpSecret(path string) (err error, res *sts.CredentialResult) {
	if path == "" { path = "*"}
	appid := global.GVA_CONFIG.TencentCOS.AppID
	bucket := global.GVA_CONFIG.TencentCOS.Bucket
	region := global.GVA_CONFIG.TencentCOS.Region
	secretID := global.GVA_CONFIG.TencentCOS.SecretID
	SecretKey := global.GVA_CONFIG.TencentCOS.SecretKey
	c := sts.NewClient(
		secretID,
		SecretKey,
		nil,
	)
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          region,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"cos:PostObject",
						"cos:PutObject",
						"cos:GetObject",
					},
					Effect: "allow",
					Resource: []string{
						//这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						"qcs::cos:" + region + ":uid/" + appid + ":" + bucket + "/" + path,
					},
				},
			},
		},
	}
	res, err = c.GetCredential(opt)
	if err != nil {
		return err, nil
	}
	return err, res
}

// CreateUpload 创建Upload记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateUpload(upload model.Upload) (err error) {
	err = global.GVA_DB.Create(&upload).Error
	return err
}

// DeleteUpload 删除Upload记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteUpload(upload model.Upload) (err error) {
	err = global.GVA_DB.Delete(&upload).Error
	return err
}

// DeleteUploadByIds 批量删除Upload记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteUploadByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Upload{},"id in ?",ids.Ids).Error
	return err
}

// UpdateUpload 更新Upload记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateUpload(upload model.Upload) (err error) {
	err = global.GVA_DB.Save(&upload).Error
	return err
}

// GetUpload 根据id获取Upload记录
// Author [piexlmax](https://github.com/piexlmax)
func GetUpload(id uint) (err error, upload model.Upload) {
	err = global.GVA_DB.Where("id = ?", id).First(&upload).Error
	return
}

// GetUploadInfoList 分页获取Upload记录
// Author [piexlmax](https://github.com/piexlmax)
func GetUploadInfoList(info request.UploadSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Upload{})
	var uploads []model.Upload
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&uploads).Error
	return err, uploads, total
}

