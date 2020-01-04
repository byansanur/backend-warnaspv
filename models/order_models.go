package models

import (
	"../structs"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
	"time"
	//"strconv"
)

func GetOrderStatus(id_order_mst_status string, id_mst_outlet string, id_mst_branch []int, created_at string, offset string, limit string, first_name string, status string) structs.JsonResponse {

	var (
		orderstatus []structs.GetOrderStatus
		t           structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()
	//order := "order"

	err := idb.DB.Table(order + " orders").Select("distinct (orders.id),to_char(orders.created_at, 'YYYY-MM-DD HH24:MI') as created_at ,  orders.id_order_mst_reason," +
		"orders.id_order_mst_status,  contact.first_name, contact.last_name,order_product_ufi.id_mst_unit," +
		"mst_unit.merk, mst_unit.type, mst_unit.year, order_loan.plafond, order_loan.installment, order_loan.tenor," +
		"order_mst_status.status as status, order_mst_reason.reason as reason")

	err = err.Joins("left join contact on contact.id = orders.id_contact")
	err = err.Joins(" left join order_product_ufi on order_product_ufi.id_order = orders.id")
	err = err.Joins(" left join mst_unit on mst_unit.id = order_product_ufi.id_mst_unit")
	err = err.Joins(" left join order_loan on order_loan.id_order = orders.id")
	err = err.Joins("left join order_mst_reason on order_mst_reason.id = orders.id_order_mst_reason")
	err = err.Joins("join order_mst_status on order_mst_status.id = orders.id_order_mst_status")
	err = err.Joins("join mst_outlet on mst_outlet.id = orders.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	if first_name != "" {
		err = err.Where("contact.first_name::text LIKE ?", "%"+first_name+"%")
	}

	if status != "" {
		err = err.Where("order_mst_status.status::text LIKE ?", "%"+status+"%")
	}

	if id_order_mst_status != "" {
		err = err.Where("orders.id_order_mst_status = ?", id_order_mst_status)
	}

	fmt.Println("Len ", len(id_mst_branch), id_mst_branch)
	if len(id_mst_branch) >= 1 {
		fmt.Println("ayam")
		err = err.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	if id_mst_outlet != "" {
		err = err.Where("orders.id_mst_outlet = ?", id_mst_outlet)
	}
	if created_at != "" {
		err = err.Where("orders.created_at::text LIKE ?", "%"+created_at+"%")
	}
	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	err = err.Order("orders.id desc")

	err = err.Find(&orderstatus)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = orderstatus
	}

	return response
}

func OrderDetail(id_order string) structs.JsonResponse {

	var (
		status structs.GetOrderDetail
		t      structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table(order + " orders").Select("distinct (orders.id), to_char(orders.created_at, 'YYYY-MM-DD HH24:MI') as created_at , orders.id_order_mst_reason," +
		"orders.id_order_mst_status, contact.first_name, contact.last_name, cms_users.name, order_mst_status.status," +
		"mst_data_source.datasource, order_product_ufi.id_mst_unit, order_product_ufi.nopol, order_product_ufi.tax_status," +
		"order_product_ufi.owner, mst_unit.merk, mst_unit.model, mst_unit.type, mst_unit.year, mst_unit.otr," +
		"order_loan.plafond, order_loan.installment, order_loan.down_payment, order_loan.need, order_loan.tenor," +
		"order_loan.otr_custom, order_surety.name as surety_name, order_mst_reason.reason")

	err = err.Joins("left join contact on  contact.id =  orders.id_contact")
	err = err.Joins("left join order_product_ufi on  order_product_ufi.id_order = orders.id")

	err = err.Joins("left join mst_unit on  mst_unit.id =  order_product_ufi.id_mst_unit")
	err = err.Joins("left join order_loan on order_loan.id_order = orders.id")
	err = err.Joins("left join order_mst_reason on order_mst_reason.id = orders.id_order_mst_reason")
	err = err.Joins("left join order_surety on  order_surety.id_order = orders.id")
	err = err.Joins("left join order_mst_status on  order_mst_status.id = orders.id_order_mst_status")
	err = err.Joins("left join mst_data_source on  mst_data_source.id = orders.id_mst_data_source")
	err = err.Joins("left join cms_users on  cms_users.id = orders.id_cms_users")
	err = err.Joins("left join mst_outlet on mst_outlet.id = orders.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	err = err.Where("orders.id = ?", id_order)

	err = err.Find(&status)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = status
	}

	return response
}

func UpdateOrderStatus(id string, id_order_mst_status string,
	id_order_mst_reason string, updated_by string) structs.JsonResponse {

	var (
		updateorderstatus structs.UpdateOrderStatus
	)

	response := structs.JsonResponse{}

	id_conv, _ := strconv.Atoi(id)
	id_order_mst_status_conv, _ := strconv.Atoi(id_order_mst_status)
	id_order_mst_reason_conv, _ := strconv.Atoi(id_order_mst_reason)
	updated_by_conv, _ := strconv.Atoi(updated_by)

	fmt.Println("id", id_conv)
	fmt.Println("status ", id_order_mst_status_conv)
	fmt.Println("reason", id_order_mst_reason_conv)
	fmt.Println("up ", updated_by_conv)

	updateorderstatus.Id = id_conv
	updateorderstatus.IdOrderMstReason = id_order_mst_reason_conv
	updateorderstatus.IdOrderMstStatus = id_order_mst_status_conv
	updateorderstatus.UpdatedBy = updated_by_conv

	err := idb.DB.Table(order+" orders").Where("id = ?", id).Updates(updateorderstatus).Error

	fmt.Println("err", err)
	if err != nil {

		response.ApiStatus = 0
		response.ApiMessage = err.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = "success"
		response.Data = updateorderstatus
	}

	return response
}

func OrderUsers(id_mst_outlet string, id_mst_branch []int, created_at string, limit string, offset string) structs.JsonResponse {

	var (
		status []structs.OrderUsers
		t      structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	var orders = `(select count(id) from "order" a where a.id_cms_users = cms_users.id and a.created_at::text like '%` + created_at + `%') as orders,`
	var approve = `(select count(id) from "order" a where a.id_cms_users = cms_users.id and  a.id_order_mst_status = 1 and a.created_at::text like '%` + created_at + `%') as approve,`
	var process = `(select count(id) from "order" a where a.id_cms_users = cms_users.id and  a.id_order_mst_status = 2 and a.created_at::text like '%` + created_at + `%') as process,`
	var cancel = `(select count(id) from "order" a where a.id_cms_users = cms_users.id and  a.id_order_mst_status = 3 and a.created_at::text like '%` + created_at + `%') as cancel,`
	var reject = `(select count(id) from "order" a where a.id_cms_users = cms_users.id and  a.id_order_mst_status = 4 and a.created_at::text like '%` + created_at + `%') as reject,`
	var paid = `(select count(id) from "order" a where a.id_cms_users = cms_users.id and  a.id_order_mst_status = 5 and a.created_at::text like '%` + created_at + `%') as paid`

	fmt.Println(orders)

	err := idb.DB.Table("cms_users").Select("distinct (cms_users.id), cms_users.npm, cms_users.name , mst_outlet.outlet_name," +
		"cms_users.status,cms_privileges.name as privileges_name, mst_branch.branch_name," +
		orders + approve + process + cancel + reject + paid)

	err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")
	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	if id_mst_outlet != "" {
		err = err.Where("cms_users.id_mst_outlet = ?", id_mst_outlet)
	}

	err = err.Find(&status)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	}

	if len(status) <= 0 {

		response.ApiStatus = 1
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = status

	}

	return response
}

func GetDealsDetail(id_cms_users string, id_mst_branch string, created_at string, offset string, limit string) structs.JsonResponse {

	var (
		orderstatus []structs.GetOrderStatus
		t           structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()
	//order := "order"
 
	err := idb.DB.Table(order + " orders").Select("distinct (orders.id), orders.created_at,  orders.id_order_mst_reason," +
		"orders.id_order_mst_status,  contact.first_name, contact.last_name,order_product_ufi.id_mst_unit," +
		"mst_unit.merk, mst_unit.type, mst_unit.year, order_loan.plafond, order_loan.installment, order_loan.tenor," +
		"order_mst_status.status as status, order_mst_reason.reason as reason")

	err = err.Joins("left join contact on contact.id = orders.id_contact")
	err = err.Joins(" left join order_product_ufi on order_product_ufi.id_order = orders.id")
	err = err.Joins(" left join mst_unit on mst_unit.id = order_product_ufi.id_mst_unit")
	err = err.Joins(" left join order_loan on order_loan.id_order = orders.id")
	err = err.Joins("left join order_mst_reason on order_mst_reason.id = orders.id_order_mst_reason")
	err = err.Joins("join order_mst_status on order_mst_status.id = orders.id_order_mst_status")
	err = err.Joins("join mst_outlet on mst_outlet.id = orders.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	if id_cms_users != "" {
		err = err.Where("orders.id_cms_users = ?", id_cms_users)
	}

	if id_mst_branch != "" {
		err = err.Where("mst_outlet.id_mst_branch= ?", id_mst_branch)
	}

	if created_at != "" {
		err = err.Where("orders.created_at::text like  ?", `'%`+created_at+`%'`)
	}

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	err = err.Order("orders.created_at desc")

	err = err.Find(&orderstatus)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = orderstatus
	}

	return response
}

func GetOrderBranch(id_mst_branch string, created_at string) structs.JsonResponse {

	var (
		orderstatus []structs.GetOrderBranch
		t           structs.Component
	)

	response := structs.JsonResponse{}

	err := idb.DB.Table(order + " orders").Select("count(orders.id) as value,mst_branch.branch_name as name")

	err = err.Joins("join mst_outlet on mst_outlet.id = orders.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	err = err.Where("orders.id_order_mst_status = 5")
	err = err.Group("mst_branch.branch_name")

	if created_at != "" {
		err = err.Where("orders.created_at::text like  ?", "%"+created_at+"%")
	}
	if id_mst_branch != "" {
		err = err.Where("orders.id_mst_branch = ?", id_mst_branch)
	}

	err = err.Order("value desc")

	err = err.Find(&orderstatus)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = orderstatus
	}

	return response
}

func GetOrderChart(id_mst_branch []int, created_at string) structs.JsonResponse {

	var (
		orderstatus structs.GetOrderChart
		t           structs.Component
	)

	response := structs.JsonResponse{}

	err := idb.DB.Table(order + " orders").Select("count(orders.id) as booking")

	err = err.Joins("join mst_outlet on mst_outlet.id = orders.id_mst_outlet")

	err = err.Where("orders.id_order_mst_status = 5")

	if created_at != "" {
		err = err.Where("orders.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch  IN (?)", id_mst_branch)
	}

	fmt.Println("created_at input ", created_at)

	err = err.First(&orderstatus)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = orderstatus
	}

	return response
}

func OrderDownload(id_mst_outlet string, id_mst_branch []int, created_at string, limit string, offset string) structs.JsonResponse {

	var (
		orderdownload []structs.OrderDownload
		t             structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table(order + " orders").Select("orders.id," +
		"orders.created_at , orders.id_mst_data_source, orders.category, orders.id_mst_outlet," +
		" orders.id_mst_cabang_fif, orders.id_contact, orders.id_mst_product," +
		//"(select orders.status_address )"+
		"(CASE WHEN orders.status_address = 'S'  THEN 'Sesuai' ELSE 'Tidak Sesuai' END) as status_address," +

		"orders.no_order, orders.id_order_mst_status, orders.id_order_mst_reason, orders.survey,orders.id_cms_users," +
		"order_loan.down_payment, order_loan.installment,order_loan.need,order_loan.otr_custom,order_loan.plafond," +
		"order_mst_reason.reason, order_mst_status.status as order_status, order_loan.tenor," +
		"order_note.note as order_note, order_product_ufi.nomor_taksasi,order_product_ufi.nopol,order_product_ufi.otr_taksasi," +
		"order_product_ufi.owner, order_product_ufi.tax_status," +
		"mst_cabang_fif.branch_name as fif_branch, mst_cabang_fif.code_pos_fif, mst_cabang_fif.pos_name," +
		"mst_unit.kode_unit, mst_unit.merk, mst_unit.model, mst_unit.type, mst_unit.otr, mst_unit.year ," +
		"order_surety.name as surety_name, order_surety.birth_date as surety_birth_date," +
		"  (select mst_job.job from mst_job where order_surety.id_mst_job = mst_job.id) as surety_job," +
		"order_surety.company as surety_company, order_surety.position as surety_position," +
		"order_surety.working_time as surety_working_time ,order_surety.income as surety_income," +
		"order_surety.outlay as surety_outlay, order_surety.birth_place as surety_birth_place, " +
		"contact.first_name as contact_first_name, contact.last_name as contact_last_name," +
		"contact.nik as contact_nik, contact.birth_date as contact_birth_date, contact.birth_place as contact_birth_place," +
		"contact.gender as contact_gender , contact.id_mst_religion, mst_religion.agama as contact_religion," +
		"contact_mst_status_marital.status as contact_material ," +
		" (select mst_job.job from mst_job where contact.id_mst_job = mst_job.id) as contact_job," +
		"  (select mst_data_source.datasource from mst_data_source where contact.id_mst_data_source = mst_data_source.id) as contact_datasource," +
		"contact_additional.company as contact_company, contact_additional.email as contact_email," +
		"contact_additional.mother as contact_mother, contact_additional.family as contact_family," +
		"contact_mst_status_place.status as contact_place_status, contact_mst_status_employee.status as contact_employee_status," +
		"contact_additional.position as contact_position ,contact_additional.working_time as contact_working_time," +
		"contact_additional.income as contact_income, contact_additional.outlay as contact_outlay," +
		"contact_note.note as contact_note," +
		"contact_address.address as contact_address, contact_address.rt as contact_rt, contact_address.rw as contact_rw,  " +
		"mst_address.kelurahan as contact_kelurahan, mst_address.kecamatan as contact_kecamatan, mst_address.kabupaten as contact_kabupaten," +
		"mst_address.provinsi as contact_provinsi, mst_address.kodepos as contact_kodepos," +
		"mst_category_address.category as contact_address_category," +
		" (select contact_phone.number from contact_phone where contact_phone.id_contact = contact.id order by id desc limit 1) as contact_phone_satu," +
		" (select contact_phone.number from contact_phone where contact_phone.id_contact = contact.id order by id asc limit 1) as contact_phone_dua," +
		" cms_users.name as users_name ,cms_users.npm as users_npm, cms_privileges.name as privileges_name," +
		" (select mst_outlet.outlet_name from mst_outlet where cms_users.id_mst_outlet = mst_outlet.id) as users_outlet," +
		" (select mst_branch.branch_name from mst_branch where mst_outlet.id_mst_branch= mst_branch.id) as users_branch," +
		"(select cms_users_cabang.id_cms_users_oh from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_oh," +
		"(select name from cms_users where cms_users.id = id_cms_users_oh )  as oh_name," +
		"(select cms_users_cabang.id_cms_users_spv from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_spv," +
		"(select name from cms_users where cms_users.id = id_cms_users_spv )  as spv_name")

	err = err.Joins("left join order_loan on order_loan.id_order = orders.id and order_loan.id = (select max(a.id) from order_loan a where a.id_order = order_loan.id_order)") //
	err = err.Joins("left join mst_cabang_fif on mst_cabang_fif.id =  orders.id_mst_cabang_fif")
	err = err.Joins("left join mst_outlet on mst_outlet.id =  orders.id_mst_outlet")

	err = err.Joins("left join contact on contact.id = orders.id_contact ")
	err = err.Joins("left join mst_product on mst_product.id = orders.id_mst_product")
	err = err.Joins("left join order_mst_status on order_mst_status.id = orders.id_order_mst_status")
	err = err.Joins("left join order_mst_reason on order_mst_reason.id = orders.id_order_mst_reason")
	err = err.Joins("left join cms_users on cms_users.id = orders.id_cms_users")
	err = err.Joins("left join order_note on order_note.id_order = orders.id and order_note.id = (select max(a.id) from order_note a where a.id_order = order_note.id_order)")                                    //
	err = err.Joins("left join order_product_ufi on order_product_ufi.id_order = orders.id and order_product_ufi.id = (select max(a.id) from order_product_ufi a where a.id_order = order_product_ufi.id_order)") //
	err = err.Joins("left join mst_unit on mst_unit.id = order_product_ufi.id_mst_unit")
	err = err.Joins("left join order_surety on order_surety.id_order = orders.id and order_surety.id = (select max(a.id) from order_surety a where a.id_order = order_surety.id_order)") //
	err = err.Joins("left join mst_religion on mst_religion.id = contact.id_mst_religion")
	err = err.Joins("left join contact_mst_status_marital on contact_mst_status_marital.id = contact.id_contact_mst_status_marital")
	err = err.Joins("left join contact_additional on contact_additional.id_contact = contact.id and contact_additional.id = (select max(a.id) from contact_additional a where a.id_contact = contact_additional.id_contact)")
	err = err.Joins("left join contact_mst_status_place on contact_mst_status_place.id = contact_additional.id_contact_mst_status_place")
	err = err.Joins("left join contact_mst_status_employee on contact_mst_status_employee.id = contact_additional.id_contact_mst_status_employee")
	err = err.Joins("left join contact_note on contact_note.id_contact = contact.id and contact_note.id = (select max(a.id) from contact_note a where a.id_contact =  contact_note.id_contact )")
	err = err.Joins("left join contact_address on contact_address.id_contact = contact.id ")
	err = err.Joins("left join mst_address on mst_address.id = contact_address.id_mst_address")
	err = err.Joins("left join mst_category_address on mst_category_address.id = contact_address.id_mst_category_address")
	err = err.Joins("left join cms_privileges on cms_privileges.id = cms_users.id_cms_privileges")
	err = err.Joins("left join cms_users_cabang on cms_users_cabang.id_cms_users = cms_users.id")

	//err = err.Where("date(orders.created_at) BETWEEN ? AND ?", created_at1, created_at2)

	if limit != "" {
		err = err.Limit(limit)
	}

	if offset != "" {
		err = err.Offset(offset)
	}
	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	if id_mst_outlet != "" {
		err = err.Where("orders.id_mst_outlet = ?", id_mst_outlet)
	}
	if created_at != "" {
		err = err.Where("orders.created_at::text like ?", "%"+created_at+"%")
	}
	err = err.Order("orders.id desc")

	err = err.Find(&orderdownload)
	errx := err.Error

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	f.SetCellValue("Sheet1", "A1", "Created_at")
	f.SetCellValue("Sheet1", "B1", "Category")
	f.SetCellValue("Sheet1", "C1", "status_address")
	f.SetCellValue("Sheet1", "D1", "no_order")
	f.SetCellValue("Sheet1", "E1", "survey")
	f.SetCellValue("Sheet1", "F1", "down_payment")
	f.SetCellValue("Sheet1", "G1", "installment")
	f.SetCellValue("Sheet1", "H1", "need")
	f.SetCellValue("Sheet1", "I1", "Tenor")
	f.SetCellValue("Sheet1", "J1", "otr_custom")
	f.SetCellValue("Sheet1", "K1", "plafond")
	f.SetCellValue("Sheet1", "L1", "reason")
	f.SetCellValue("Sheet1", "M1", "order_status")
	f.SetCellValue("Sheet1", "N1", "order_note")
	f.SetCellValue("Sheet1", "O1", "nomor_transaksi")
	f.SetCellValue("Sheet1", "P1", "nopol")
	f.SetCellValue("Sheet1", "Q1", "otr_transaksi")
	f.SetCellValue("Sheet1", "R1", "owner")
	f.SetCellValue("Sheet1", "S1", "tax_status")
	f.SetCellValue("Sheet1", "T1", "fif_branch")
	f.SetCellValue("Sheet1", "U1", "code_pos_fif")
	f.SetCellValue("Sheet1", "V1", "pos_name")
	f.SetCellValue("Sheet1", "W1", "kode_unit")
	f.SetCellValue("Sheet1", "X1", "merk")
	f.SetCellValue("Sheet1", "Y1", "type")
	f.SetCellValue("Sheet1", "Z1", "model")
	f.SetCellValue("Sheet1", "AA1", "year")
	f.SetCellValue("Sheet1", "AB1", "surety_name")
	f.SetCellValue("Sheet1", "AC1", "surety_birth_date")
	f.SetCellValue("Sheet1", "AD1", "surety_birth_place")
	f.SetCellValue("Sheet1", "AE1", "surety_job")
	f.SetCellValue("Sheet1", "AF1", "surety_company")
	f.SetCellValue("Sheet1", "AG1", "surety_position")
	f.SetCellValue("Sheet1", "AH1", "surety_working_time")
	f.SetCellValue("Sheet1", "AI1", "surety_income")
	f.SetCellValue("Sheet1", "AJ1", "surety_outlay")
	f.SetCellValue("Sheet1", "AK1", "contact_first_name")
	f.SetCellValue("Sheet1", "AL1", "contact_last_name")
	f.SetCellValue("Sheet1", "AM1", "contact_nik")
	f.SetCellValue("Sheet1", "AN1", "contact_birth_date")
	f.SetCellValue("Sheet1", "AO1", "contact_birth_place")
	f.SetCellValue("Sheet1", "AP1", "contact_religion")
	f.SetCellValue("Sheet1", "AQ1", "contact_material")
	f.SetCellValue("Sheet1", "AR1", "contact_job")
	f.SetCellValue("Sheet1", "AS1", "contact_datasource")
	f.SetCellValue("Sheet1", "AT1", "contact_company")
	f.SetCellValue("Sheet1", "AU1", "contact_mother")
	f.SetCellValue("Sheet1", "AV1", "contact_family")
	f.SetCellValue("Sheet1", "AW1", "contact_place_status")
	f.SetCellValue("Sheet1", "AX1", "contact_employee_status")
	f.SetCellValue("Sheet1", "AY1", "contact_position")
	f.SetCellValue("Sheet1", "AZ1", "contact_working_time")
	f.SetCellValue("Sheet1", "BA1", "contact_income")
	f.SetCellValue("Sheet1", "BB1", "contact_outlay")
	f.SetCellValue("Sheet1", "BC1", "contact_note")
	f.SetCellValue("Sheet1", "BD1", "contact_phone_1")
	f.SetCellValue("Sheet1", "BE1", "contact_phone_2")
	f.SetCellValue("Sheet1", "BF1", "contact_address")
	f.SetCellValue("Sheet1", "BG1", "contact_rt")
	f.SetCellValue("Sheet1", "BH1", "contact_rw")
	f.SetCellValue("Sheet1", "BI1", "contact_kelurahan")
	f.SetCellValue("Sheet1", "BJ1", "contact_kecamatan")
	f.SetCellValue("Sheet1", "BK1", "contact_kebupaten")
	f.SetCellValue("Sheet1", "BL1", "contact_provinsi")
	f.SetCellValue("Sheet1", "BM1", "contact_kodepos")
	f.SetCellValue("Sheet1", "BN1", "contact_address_category")
	f.SetCellValue("Sheet1", "BO1", "users_name")
	f.SetCellValue("Sheet1", "BP1", "users_npm")
	f.SetCellValue("Sheet1", "BQ1", "privileges_name")
	f.SetCellValue("Sheet1", "BR1", "users_outlet")
	f.SetCellValue("Sheet1", "BS1", "users_branch")
	f.SetCellValue("Sheet1", "BT1", "Oh Name")
	f.SetCellValue("Sheet1", "BU1", "Spv Name")

	for i, _ := range orderdownload {

		rows := 2

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+rows), orderdownload[i].CreatedAt)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+rows), orderdownload[i].Category)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+rows), orderdownload[i].StatusAddress)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+rows), orderdownload[i].NoOrder)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+rows), orderdownload[i].Survey)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+rows), orderdownload[i].DownPayment)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+rows), orderdownload[i].Installment)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+rows), orderdownload[i].Need)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+rows), orderdownload[i].Tenor)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(i+rows), orderdownload[i].OtrCustom)
		f.SetCellValue("Sheet1", "K"+strconv.Itoa(i+rows), orderdownload[i].Plafond)
		f.SetCellValue("Sheet1", "L"+strconv.Itoa(i+rows), orderdownload[i].Reason)
		f.SetCellValue("Sheet1", "M"+strconv.Itoa(i+rows), orderdownload[i].OrderStatus)
		f.SetCellValue("Sheet1", "N"+strconv.Itoa(i+rows), orderdownload[i].OrderNote)
		f.SetCellValue("Sheet1", "O"+strconv.Itoa(i+rows), orderdownload[i].NomorTransaksi)
		f.SetCellValue("Sheet1", "P"+strconv.Itoa(i+rows), orderdownload[i].Nopol)
		f.SetCellValue("Sheet1", "Q"+strconv.Itoa(i+rows), orderdownload[i].OtrTransaksi)
		f.SetCellValue("Sheet1", "R"+strconv.Itoa(i+rows), orderdownload[i].Owner)
		f.SetCellValue("Sheet1", "S"+strconv.Itoa(i+rows), orderdownload[i].TaxStatus)
		f.SetCellValue("Sheet1", "T"+strconv.Itoa(i+rows), orderdownload[i].FifBranch)
		f.SetCellValue("Sheet1", "U"+strconv.Itoa(i+rows), orderdownload[i].CodePosFif)
		f.SetCellValue("Sheet1", "V"+strconv.Itoa(i+rows), orderdownload[i].PosName)
		f.SetCellValue("Sheet1", "W"+strconv.Itoa(i+rows), orderdownload[i].KodeUnit)
		f.SetCellValue("Sheet1", "X"+strconv.Itoa(i+rows), orderdownload[i].Merk)
		f.SetCellValue("Sheet1", "Y"+strconv.Itoa(i+rows), orderdownload[i].Type)
		f.SetCellValue("Sheet1", "Z"+strconv.Itoa(i+rows), orderdownload[i].Model)
		f.SetCellValue("Sheet1", "AA"+strconv.Itoa(i+rows), orderdownload[i].Year)
		f.SetCellValue("Sheet1", "AB"+strconv.Itoa(i+rows), orderdownload[i].SuretyName)
		f.SetCellValue("Sheet1", "AC"+strconv.Itoa(i+rows), orderdownload[i].SuretyBirthDate)
		f.SetCellValue("Sheet1", "AD"+strconv.Itoa(i+rows), orderdownload[i].SuretyBirthPlace)
		f.SetCellValue("Sheet1", "AE"+strconv.Itoa(i+rows), orderdownload[i].SuretyJob)
		f.SetCellValue("Sheet1", "AF"+strconv.Itoa(i+rows), orderdownload[i].SuretyCompany)
		f.SetCellValue("Sheet1", "AG"+strconv.Itoa(i+rows), orderdownload[i].SuretyPosition)
		f.SetCellValue("Sheet1", "AH"+strconv.Itoa(i+rows), orderdownload[i].SuretyWorkingTime)
		f.SetCellValue("Sheet1", "AI"+strconv.Itoa(i+rows), orderdownload[i].SuretyIncome)
		f.SetCellValue("Sheet1", "AJ"+strconv.Itoa(i+rows), orderdownload[i].SuretyOutlay)
		f.SetCellValue("Sheet1", "AK"+strconv.Itoa(i+rows), orderdownload[i].ContactFirstName)
		f.SetCellValue("Sheet1", "AL"+strconv.Itoa(i+rows), orderdownload[i].ContactLastName)
		f.SetCellValue("Sheet1", "AM"+strconv.Itoa(i+rows), orderdownload[i].ContactNik)
		f.SetCellValue("Sheet1", "AN"+strconv.Itoa(i+rows), orderdownload[i].ContactBirthDate)
		f.SetCellValue("Sheet1", "AO"+strconv.Itoa(i+rows), orderdownload[i].ContactBirthPlace)
		f.SetCellValue("Sheet1", "AP"+strconv.Itoa(i+rows), orderdownload[i].ContactReligion)
		f.SetCellValue("Sheet1", "AQ"+strconv.Itoa(i+rows), orderdownload[i].ContactMaterial)
		f.SetCellValue("Sheet1", "AR"+strconv.Itoa(i+rows), orderdownload[i].ContactJob)
		f.SetCellValue("Sheet1", "AS"+strconv.Itoa(i+rows), orderdownload[i].ContactDatasource)
		f.SetCellValue("Sheet1", "AT"+strconv.Itoa(i+rows), orderdownload[i].ContactCompany)
		f.SetCellValue("Sheet1", "AU"+strconv.Itoa(i+rows), orderdownload[i].ContactMother)
		f.SetCellValue("Sheet1", "AV"+strconv.Itoa(i+rows), orderdownload[i].ContactFamily)
		f.SetCellValue("Sheet1", "AW"+strconv.Itoa(i+rows), orderdownload[i].ContactPlaceStatus)
		f.SetCellValue("Sheet1", "AX"+strconv.Itoa(i+rows), orderdownload[i].ContactEmployeeStatus)
		f.SetCellValue("Sheet1", "AY"+strconv.Itoa(i+rows), orderdownload[i].ContactPosition)
		f.SetCellValue("Sheet1", "AZ"+strconv.Itoa(i+rows), orderdownload[i].ContactWorkingTime)
		f.SetCellValue("Sheet1", "BA"+strconv.Itoa(i+rows), orderdownload[i].ContactIncome)
		f.SetCellValue("Sheet1", "BB"+strconv.Itoa(i+rows), orderdownload[i].ContactOutlay)
		f.SetCellValue("Sheet1", "BC"+strconv.Itoa(i+rows), orderdownload[i].ContactNote)
		f.SetCellValue("Sheet1", "BD"+strconv.Itoa(i+rows), orderdownload[i].ContactPhoneSatu)
		f.SetCellValue("Sheet1", "BE"+strconv.Itoa(i+rows), orderdownload[i].ContactPhoneDua)
		f.SetCellValue("Sheet1", "BF"+strconv.Itoa(i+rows), orderdownload[i].ContactAddress)
		f.SetCellValue("Sheet1", "BG"+strconv.Itoa(i+rows), orderdownload[i].ContactRt)
		f.SetCellValue("Sheet1", "BH"+strconv.Itoa(i+rows), orderdownload[i].ContactRw)
		f.SetCellValue("Sheet1", "BI"+strconv.Itoa(i+rows), orderdownload[i].ContactKelurahan)
		f.SetCellValue("Sheet1", "BJ"+strconv.Itoa(i+rows), orderdownload[i].ContactKecamatan)
		f.SetCellValue("Sheet1", "BK"+strconv.Itoa(i+rows), orderdownload[i].ContactKabupaten)
		f.SetCellValue("Sheet1", "BL"+strconv.Itoa(i+rows), orderdownload[i].ContactProvinsi)
		f.SetCellValue("Sheet1", "BM"+strconv.Itoa(i+rows), orderdownload[i].ContactKodepos)
		f.SetCellValue("Sheet1", "BN"+strconv.Itoa(i+rows), orderdownload[i].ContactAddressCategory)
		f.SetCellValue("Sheet1", "BO"+strconv.Itoa(i+rows), orderdownload[i].UsersName)
		f.SetCellValue("Sheet1", "BP"+strconv.Itoa(i+rows), orderdownload[i].UsersNpm)
		f.SetCellValue("Sheet1", "BQ"+strconv.Itoa(i+rows), orderdownload[i].PrivilegesName)
		f.SetCellValue("Sheet1", "BR"+strconv.Itoa(i+rows), orderdownload[i].UsersOutlet)
		f.SetCellValue("Sheet1", "BS"+strconv.Itoa(i+rows), orderdownload[i].UsersBranch)
		f.SetCellValue("Sheet1", "BT"+strconv.Itoa(i+rows), orderdownload[i].OhName)
		f.SetCellValue("Sheet1", "BU"+strconv.Itoa(i+rows), orderdownload[i].SpvName)

	}

	f.SetActiveSheet(index)
	folderPath := "./files/"
	filename := "order" + time.Now().Format("20060102150405") + ".xlsx"

	os.MkdirAll(folderPath, os.ModePerm)

	// Save xlsx file by the given path.
	errs := f.SaveAs(folderPath + filename)
	if errs != nil {
		fmt.Println("err create excel", errs)
	}

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	}

	if len(orderdownload) <= 0 {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = nil

	} else {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = filename

	}

	return response
}

func DealStatus(id_mst_branch []int, created_at string) structs.JsonResponse {

	var (
		status structs.DealStatus
		t      structs.Component
	)
	response := structs.JsonResponse{}

	//approve
	approve := idb.DB.Table(order + " orders").Select("count(orders.id) as approve")
	approve = approve.Joins("left join mst_outlet on mst_outlet.id = orders.id_mst_outlet")
	approve = approve.Where("orders.id_order_mst_status = 1")
	if created_at != "" {
		approve = approve.Where("orders.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		approve = approve.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	approve = approve.Find(&status)
	errapprove := approve.Error

	//process
	process := idb.DB.Table(order + " orders").Select("count(orders.id) as process")
	process = process.Joins("left join mst_outlet on mst_outlet.id = orders.id_mst_outlet")
	process = process.Where(" orders.id_order_mst_status = 2")

	if created_at != "" {
		process = process.Where("orders.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		process = process.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	process = process.Find(&status)
	errprocess := process.Error

	//cancel
	cancel := idb.DB.Table(order + " orders").Select("count(orders.id) as cancel")
	cancel = cancel.Joins("left join mst_outlet on mst_outlet.id = orders.id_mst_outlet")
	cancel = cancel.Where(" orders.id_order_mst_status = 3")

	if created_at != "" {
		cancel = cancel.Where("orders.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		cancel = cancel.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	cancel = cancel.Find(&status)
	errcancel := cancel.Error

	//reject
	reject := idb.DB.Table(order + " orders").Select("count(orders.id) as reject")
	reject = reject.Joins(" left join mst_outlet on mst_outlet.id = orders.id_mst_outlet")
	reject = reject.Where(" orders.id_order_mst_status = 4")

	if created_at != "" {
		reject = reject.Where("orders.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		reject = reject.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	reject = reject.Find(&status)
	errreject := reject.Error

	//paid
	paid := idb.DB.Table(order + " orders").Select("count(orders.id) as paid")
	paid = paid.Joins(" left join mst_outlet on mst_outlet.id = orders.id_mst_outlet")
	paid = paid.Where(" orders.id_order_mst_status = 5")

	if created_at != "" {
		paid = paid.Where("orders.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		paid = paid.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	paid = paid.Find(&status)
	errpaid := paid.Error

	if errapprove != nil && errprocess != nil && errcancel != nil && errreject != nil && errpaid != nil {

		response.ApiStatus = 0
		response.ApiMessage = errapprove.Error()
		response.Data = nil

	} else {

		convStruct := [5]structs.Conv{}

		convStruct[0].Name = "Approve"
		convStruct[0].Value = status.Approve
		convStruct[1].Name = "Process"
		convStruct[1].Value = status.Process
		convStruct[2].Name = "Cancel"
		convStruct[2].Value = status.Cancel
		convStruct[3].Name = "Reject"
		convStruct[3].Value = status.Reject
		convStruct[4].Name = "Paid"
		convStruct[4].Value = status.Paid

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = convStruct

		fmt.Print("ajg lah", status)
	}

	return response
}

func LeadStatus(id_mst_branch []int, created_at string) structs.JsonResponse {

	var (
		status structs.LeadStatus
		t      structs.Component
	)

	response := structs.JsonResponse{}

	//new
	news := idb.DB.Table("lead").Select("count(lead.id) as new")
	news = news.Joins("join mst_outlet on mst_outlet.id = lead.id_mst_outlet")
	news = news.Where("lead.id_lead_mst_status = 1")
	if created_at != "" {
		news = news.Where("lead.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		news = news.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	news = news.Find(&status)
	errnews := news.Error

	//hot

	hots := idb.DB.Table("lead").Select("count(lead.id) as hot")
	hots = hots.Joins("join mst_outlet on mst_outlet.id = lead.id_mst_outlet")
	hots = hots.Where("lead.id_lead_mst_status = 2")
	if created_at != "" {
		hots = hots.Where("lead.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		hots = hots.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	hots = hots.Find(&status)
	errhots := hots.Error

	//working

	working := idb.DB.Table("lead").Select("count(lead.id) as working")
	working = working.Joins("join mst_outlet on mst_outlet.id = lead.id_mst_outlet")
	working = working.Where("lead.id_lead_mst_status = 3")
	if created_at != "" {
		working = working.Where("lead.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		working = working.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	working = working.Find(&status)
	errworking := working.Error

	//unqualified
	unqualified := idb.DB.Table("lead").Select("count(lead.id) as unqualified")
	unqualified = unqualified.Joins("join mst_outlet on mst_outlet.id = lead.id_mst_outlet")
	unqualified = unqualified.Where("lead.id_lead_mst_status = 4")
	if created_at != "" {
		unqualified = unqualified.Where("lead.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		unqualified = unqualified.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	unqualified = unqualified.Find(&status)
	errunqualified := unqualified.Error

	//converted
	converted := idb.DB.Table("lead").Select("count(lead.id) as converted")
	converted = converted.Joins("join mst_outlet on mst_outlet.id = lead.id_mst_outlet")
	converted = converted.Where("lead.id_lead_mst_status = 5")
	if created_at != "" {
		converted = converted.Where("lead.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		converted = converted.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	converted = converted.Find(&status)
	errconverted := converted.Error

	//deleted
	deleted := idb.DB.Table("lead").Select("count(lead.id) as deleted")
	deleted = deleted.Joins("join mst_outlet on mst_outlet.id = lead.id_mst_outlet")
	deleted = deleted.Where("lead.id_lead_mst_status = 6")
	if created_at != "" {
		deleted = deleted.Where("lead.created_at::text LIKE  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		deleted = deleted.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}
	deleted = deleted.Find(&status)
	errdeleted := deleted.Error

	if errnews != nil && errhots != nil && errworking != nil && errunqualified != nil && errconverted != nil && errdeleted != nil {

		response.ApiStatus = 0
		response.ApiMessage = errnews.Error()
		response.Data = nil

	} else {

		convStruct := [6]structs.Conv{}

		convStruct[0].Name = "New"
		convStruct[0].Value = status.New
		convStruct[1].Name = "Hot"
		convStruct[1].Value = status.Hot
		convStruct[2].Name = "Working"
		convStruct[2].Value = status.Working
		convStruct[3].Name = "Unqualified"
		convStruct[3].Value = status.Unqualified
		convStruct[4].Name = "Converted"
		convStruct[4].Value = status.Converted
		convStruct[5].Name = "Deleted"
		convStruct[5].Value = status.Deleted

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = convStruct
	}

	return response
}
