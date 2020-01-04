package models

import (
	"../structs"
)

func GetTargetUpload(id_mst_branch string, limit string, offset string) structs.JsonResponse {

	var (
		targetvisumdetail []structs.GetTargetUpload
		t                 structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("target_upload").Select(`target_upload.id, target_upload.id_cms_users,
target_upload.url,target_upload.description, target_upload.total , target_upload.success,
target_upload.status, target_upload.result,to_char(target_upload.created_at, 'YYYY-MM-DD HH24:MI') as created_at , cms_users.name`)

	err = err.Joins("left join cms_users on target_upload.id_cms_users = cms_users.id")
	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	if id_mst_branch != "" {
		err = err.Where("target_upload.id_mst_branch = ?", id_mst_branch)
	}
	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	err = err.Order("target_upload.id desc")

	err = err.Find(&targetvisumdetail)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = targetvisumdetail
	}

	return response
}
