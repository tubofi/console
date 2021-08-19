package request

import "gin-vue-admin/model"

type PotentialSearch struct{
	model.Potential
	PageInfo
}