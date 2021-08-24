package service

import (
	"fmt"
	"gin-vue-admin/global"
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
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", res.Credentials)

	return err, res
}
