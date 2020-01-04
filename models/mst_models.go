package models

import (
	"../structs"
	"strconv"
)

func GetOrderMstStatus() structs.JsonResponse {

	var (
		ordermststatus []structs.GetOrderMstStatus
		t              structs.Component
	)
	response := structs.JsonResponse{}

	err := idb.DB.Table("order_mst_status")

	err = err.Find(&ordermststatus)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = ordermststatus
	}

	return response
}

func GetOrderMstReason() structs.JsonResponse {

	var (
		ordermststatus []structs.GetOrderMstReason
		t              structs.Component
	)
	response := structs.JsonResponse{}

	err := idb.DB.Table("order_mst_reason")

	err = err.Find(&ordermststatus)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = ordermststatus
	}

	return response
}

func GetCmsPrivilegs() structs.JsonResponse {

	var (
		ordermststatus []structs.GetPrivileges
		t              structs.Component
	)
	response := structs.JsonResponse{}

	err := idb.DB.Table("cms_privileges")

	err = err.Find(&ordermststatus)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = ordermststatus
	}

	return response
}

func GetMstOutlet(id_mst_branch string, limit string, offset string) structs.JsonResponse {

	var (
		ordermststatus []structs.GetMstOutlet
		t              structs.Component
	)
	response := structs.JsonResponse{}

	err := idb.DB.Table("mst_outlet").Select("*")
	err = err.Joins("left join mst_branch on mst_branch.id = mst_outlet.id_mst_branch")
	err = err.Joins("left join mst_biz_type on mst_biz_type.id = mst_outlet.id_mst_biz_type")
	err = err.Order("outlet_name asc")

	if id_mst_branch != "" {
		err = err.Where(`mst_outlet.id_mst_branch = ?`, id_mst_branch)
	}

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}

	err = err.Find(&ordermststatus)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = ordermststatus
	}

	return response
}

func GetMstOutletDetail(id string) structs.JsonResponse {

	var (
		outlet structs.GetMstOutletDetail
		t      structs.Component
	)
	response := structs.JsonResponse{}

	err := idb.DB.Table("mst_outlet").Select("*")
	err = err.Joins("join mst_branch on mst_branch.id = mst_outlet.id_mst_branch")
	err = err.Joins("join mst_biz_type on mst_biz_type.id = mst_outlet.id_mst_biz_type")
	err = err.Where(`mst_outlet.id = ?`, id)

	err = err.Find(&outlet)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = outlet
	}

	return response
}

func GetMstBizType() structs.JsonResponse {

	var (
		mst_biz_type []structs.GetMstBizType
		t            structs.Component
	)
	response := structs.JsonResponse{}

	err := idb.DB.Table("mst_biz_type")

	err = err.Find(&mst_biz_type)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = mst_biz_type
	}

	return response
}

func GetMstBranch() structs.JsonResponse {

	var (
		ordermststatus []structs.GetMstBranch
		t              structs.Component
	)
	response := structs.JsonResponse{}

	err := idb.DB.Table("mst_branch")
	err = err.Order("mst_branch.branch_name asc")

	err = err.Find(&ordermststatus)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = ordermststatus
	}

	return response
}

func GetMstLogDesc() structs.JsonResponse {

	var (
		ordermststatus []structs.GetMstLogDesc
		t              structs.Component
	)
	response := structs.JsonResponse{}

	err := idb.DB.Table("mst_log_desc")

	err = err.Find(&ordermststatus)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = ordermststatus
	}

	return response
}

func GetTargetMstStatus() structs.JsonResponse {

	var (
		ordermststatus []structs.GetTargetMstStatus
		t              structs.Component
	)
	response := structs.JsonResponse{}

	err := idb.DB.Table("target_mst_status")
	err = err.Where("id != 6")

	err = err.Find(&ordermststatus)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = ordermststatus
	}

	return response
}

func GetVueMenusPrivileges(id_cms_privileges string) structs.JsonResponse {

	var (
		ordermststatus []structs.VueMenusPrivileges
		t              structs.Component
	)
	response := structs.JsonResponse{}

	err := idb.DB.Table("vue_menus_privileges").Select("vue_menus_privileges.id," +
		"vue_menus_privileges.id_vue_menus,vue_menus_privileges.id_cms_privileges," +
		"vue_menus.title,vue_menus.name,vue_menus.group,vue_menus.icon,vue_menus.component")

	err = err.Joins("join vue_menus on vue_menus.id = vue_menus_privileges.id_vue_menus")
	err = err.Where("vue_menus_privileges.id_cms_privileges = ?", id_cms_privileges)
	err = err.Order("vue_menus.id asc")

	err = err.Find(&ordermststatus)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = ordermststatus
	}

	return response
}

func GetMstUnits(id_mst_branch string, merk string, types string, model string, kode_unit string,
	otr string, limit string, offset string) structs.JsonResponse {

	var (
		mst_units []structs.MstUnits
		t         structs.Component
	)
	response := structs.JsonResponse{}

	query := `mst_unit.id, mst_unit.id_mst_branch, mst_unit.year, mst_unit.kode_unit, mst_unit.merk, 
mst_unit.type, mst_unit.model, mst_unit.otr, mst_branch.branch_name `

	err := idb.DB.Table("mst_unit").Select(query)
	err = err.Joins(`join mst_branch on mst_branch.id = mst_unit.id_mst_branch`)

	if id_mst_branch != "" {
		err = err.Where(`mst_unit.id_mst_branch = ?`, id_mst_branch)
	}
	if merk != "" {
		err = err.Where(`mst_unit.merk LIKE ? `, "%"+merk+"%")
	}
	if types != "" {
		err = err.Where(`mst_unit.type LIKE ? `, "%"+types+"%")
	}
	if model != "" {
		err = err.Where(`mst_unit.model LIKE ? `, "%"+model+"%")
	}
	if kode_unit != "" {
		err = err.Where(`mst_unit.kode_unit LIKE ? `, "%"+kode_unit+"%")
	}
	if otr != "" {
		err = err.Where(`mst_unit.otr = ? `, otr)
	}

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}

	err = err.Find(&mst_units)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = mst_units
	}

	return response
}

func GetMstUnit(id string) structs.JsonResponse {

	var (
		mst_units structs.MstUnit
		t         structs.Component
	)
	response := structs.JsonResponse{}

	query := `mst_unit.id, mst_unit.id_mst_branch, mst_unit.year, mst_unit.kode_unit, mst_unit.merk, 
mst_unit.type, mst_unit.model, mst_unit.otr, mst_branch.branch_name `

	err := idb.DB.Table("mst_unit").Select(query)
	err = err.Joins(`join mst_branch on mst_branch.id = mst_unit.id_mst_branch`)

	err = err.Where(`mst_unit.id = ?`, id)

	err = err.Find(&mst_units)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = mst_units
	}

	return response
}

func UpdateMstUnits(id string, merk string, types string, model string, kode_unit string,
	otr string) structs.JsonResponse {

	var (
		mst_units structs.UpdateMstUnits
		t         structs.Component
	)
	response := structs.JsonResponse{}

	otr_conv, _ := strconv.Atoi(otr)

	mst_units.Merk = merk
	mst_units.Type = types
	mst_units.Model = model
	mst_units.KodeUnit = kode_unit
	mst_units.Otr = otr_conv

	err := idb.DB.Table("mst_unit").Where(`id = ?`, id).Updates(&mst_units)

	err = err.Find(&mst_units)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = mst_units
	}

	return response
}

func CreateMstUnits(id_mst_branch string, year string, merk string, types string, model string, kode_unit string,
	otr string, id_cms_users string) structs.JsonResponse {

	var (
		mst_units structs.CreateMstUnits
		t         structs.Component
	)
	response := structs.JsonResponse{}

	otr_conv, _ := strconv.Atoi(otr)
	id_mst_branch_conv, _ := strconv.Atoi(id_mst_branch)
	id_cms_users_conv, _ := strconv.Atoi(id_cms_users)
	year_conv, _ := strconv.Atoi(year)

	mst_units.Merk = merk
	mst_units.Year = year_conv
	mst_units.IdMstBranch = id_mst_branch_conv
	mst_units.Type = types
	mst_units.Model = model
	mst_units.KodeUnit = kode_unit
	mst_units.Otr = otr_conv
	mst_units.IdCmsUsers = id_cms_users_conv

	err := idb.DB.Table("mst_unit").Create(&mst_units)

	//err = err.Find(&mst_units)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = mst_units
	}

	return response
}

func UpdateMstOutlet(id string, outlet_id string, outlet_sys_id string, outlet_name string, id_mst_branch string,
	id_mst_biz_type string, outlet_location string, outlet_Address string, id_mst_address string,
	outlet_fif_code string, outled_desc string, outlet_status string,
	outlet_lat string, outlet_lng string, id_cms_users string, updated_by string) structs.JsonResponse {

	var (
		mst_outlet structs.UpdateMstOutlet
		t          structs.Component
	)
	response := structs.JsonResponse{}

	id_mst_branch_conv, _ := strconv.Atoi(id_mst_branch)
	id_mst_biz_type_conv, _ := strconv.Atoi(id_mst_biz_type)
	id_cms_users_conv, _ := strconv.Atoi(id_cms_users)
	id_mst_address_conv, _ := strconv.Atoi(id_mst_address)
	updated_by_conv, _ := strconv.Atoi(updated_by)

	mst_outlet.OutletId = outlet_id
	mst_outlet.OutletSysId = outlet_sys_id
	mst_outlet.OutletName = outlet_name
	mst_outlet.IdMstBranch = id_mst_branch_conv
	mst_outlet.IdMstBizType = id_mst_biz_type_conv
	mst_outlet.OutletLocation = outlet_location
	mst_outlet.OutletAddress = outlet_Address
	mst_outlet.IdMstAddress = id_mst_address_conv
	mst_outlet.OutletFifCode = outlet_fif_code
	mst_outlet.OutletDesc = outled_desc
	mst_outlet.OutletStatus = outlet_status
	mst_outlet.OutletLat = outlet_lat
	mst_outlet.OutletLng = outlet_lng
	mst_outlet.IdCmsUsers = id_cms_users_conv
	mst_outlet.UpdatedAt = t.GetTimeNow()
	mst_outlet.UpdatedBy = updated_by_conv

	err := idb.DB.Table("mst_outlet").Where(`id = ?`, id).Updates(&mst_outlet)

	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = mst_outlet
	}

	return response
}

func CreateMstOutlet(outlet_id string, outlet_sys_id string, outlet_name string, id_mst_branch string,
	id_mst_biz_type string, outlet_location string, outlet_Address string, id_mst_address string,
	outlet_fif_code string, outled_desc string, outlet_status string,
	outlet_lat string, outlet_lng string, id_cms_users string) structs.JsonResponse {

	var (
		mst_outlet structs.CreateMstOutlet
		t          structs.Component
	)
	response := structs.JsonResponse{}

	id_mst_branch_conv, _ := strconv.Atoi(id_mst_branch)
	id_mst_biz_type_conv, _ := strconv.Atoi(id_mst_biz_type)
	id_cms_users_conv, _ := strconv.Atoi(id_cms_users)
	id_mst_address_conv, _ := strconv.Atoi(id_mst_address)

	mst_outlet.OutletId = outlet_id
	mst_outlet.OutletSysId = outlet_sys_id
	mst_outlet.OutletName = outlet_name
	mst_outlet.IdMstBranch = id_mst_branch_conv
	mst_outlet.IdMstBizType = id_mst_biz_type_conv
	mst_outlet.OutletLocation = outlet_location
	mst_outlet.OutletAddress = outlet_Address
	mst_outlet.IdMstAddress = id_mst_address_conv
	mst_outlet.OutletFifCode = outlet_fif_code
	mst_outlet.OutletDesc = outled_desc
	mst_outlet.OutletStatus = outlet_status
	mst_outlet.OutletLat = outlet_lat
	mst_outlet.OutletLng = outlet_lng
	mst_outlet.IdCmsUsers = id_cms_users_conv

	err := idb.DB.Table("mst_outlet").Create(&mst_outlet)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = mst_outlet
	}

	return response
}
