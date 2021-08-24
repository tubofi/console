package service

import (
	"fmt"
	"gin-vue-admin/global"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
)

var policyString = "{\n  \"version\": \"2.0\",\n  \"statement\": [\n    {\n      \"effect\": \"allow\",\n      \"action\": [\n        \"cos:PutObject\",\n        \"cos:GetObject\",\n        \"cos:HeadObject\",\n        \"cos:OptionsObject\",\n        \"cos:ListParts\",\n        \"cos:GetObjectTagging\"\n      ],\n      \"resource\": [\n        \"qcs::cos:ap-chengdu:uid/1304003768:saisicode-1304003768/*\"\n      ]\n    }\n  ]\n}"

func GetCredential() (err error, res *sts.GetFederationTokenResponse) {
	credential := common.NewCredential(
		global.GVA_CONFIG.TencentCOS.SecretID,
		global.GVA_CONFIG.TencentCOS.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = global.GVA_CONFIG.TencentCOS.BaseURL
	client, _ := sts.NewClient(credential, global.GVA_CONFIG.TencentCOS.Region, cpf)

	request := sts.NewGetFederationTokenRequest()

	request.Name = common.StringPtr("name")
	request.Policy = common.StringPtr(policyString)
	request.DurationSeconds = common.Uint64Ptr(7200)

	response, err := client.GetFederationToken(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return err, nil
	}
	return nil, response
}
