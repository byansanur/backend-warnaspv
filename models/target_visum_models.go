package models

import (
	"../structs"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
	"time"
)

func TargetVisumUsers(id_mst_outlet string, id_mst_branch []int, created_at string, limit string, offset string) structs.JsonResponse {

	var (
		targetvisum []structs.TargetVisumUsers
		t           structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	date := "%" + created_at + "%"

	err := idb.DB.Table("cms_users").Select("distinct (cms_users.id), cms_users.npm, cms_users.name , mst_outlet.outlet_name,"+
		"cms_users.status,cms_privileges.name as privileges_name, mst_branch.branch_name,"+
		"(select count(id) from target_visum where target_visum.id_cms_users = cms_users.id and target_visum.created_at::text LIKE  ? ) as visum,"+
		"(select count(id) from target_visum where target_visum.id_cms_users = cms_users.id and target_visum.id_mst_visum_status = 1 and target_visum.created_at::text LIKE ? ) as rumah_kosong,"+
		"(select count(id) from target_visum where target_visum.id_cms_users = cms_users.id and target_visum.id_mst_visum_status = 2 and target_visum.created_at::text LIKE ? ) as alamat_tidak_sesuai,"+
		"(select count(id) from target_visum where target_visum.id_cms_users = cms_users.id and target_visum.id_mst_visum_status = 3 and target_visum.created_at::text LIKE ? ) as ditolak_atau_diusir,"+
		"(select count(id) from target_visum where target_visum.id_cms_users = cms_users.id and target_visum.id_mst_visum_status = 4 and target_visum.created_at::text LIKE ? ) as butuh_waktu,"+
		"(select count(id) from target_visum where target_visum.id_cms_users = cms_users.id and target_visum.id_mst_visum_status = 5 and target_visum.created_at::text LIKE ? ) as tanya_pasangan,"+
		"(select count(id) from target_visum where target_visum.id_cms_users = cms_users.id and target_visum.id_mst_visum_status = 6 and target_visum.created_at::text LIKE ? ) as pikir_pikir,"+
		"(select count(id) from target_visum where target_visum.id_cms_users = cms_users.id and target_visum.id_mst_visum_status = 7 and target_visum.created_at::text LIKE ? ) as berminat,"+
		"(select count(id) from target_visum where target_visum.id_cms_users = cms_users.id and target_visum.id_mst_visum_status = 8 and target_visum.created_at::text LIKE ? ) as rumah_tidak_ketemu,"+
		"(select count(id) from target_visum where target_visum.id_cms_users = cms_users.id and target_visum.id_mst_visum_status = 9 and target_visum.created_at::text LIKE ? ) as tidak_berminat", date, date, date, date, date, date, date, date, date, date)

	err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")
	err = err.Joins("join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

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

	err = err.Find(&targetvisum)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	}

	if len(targetvisum) <= 0 {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = nil

	} else {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = targetvisum

	}

	return response
}

func TargetVisumDetail(id_cms_users string, created_at string) structs.JsonResponse {

	var (
		targetvisumdetail []structs.TargetVisumDetail
		t                 structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("target_visum").Select("distinct (target_visum.id) , target_visum.created_at," +
		"mst_data_source.datasource, target_visum.revisit, target.first_name, target.last_name," +
		"target.address as alamat,target.kabupaten, target.kecamatan,target.kelurahan," +
		"target.provinsi, mst_visum_status.status,cms_users.name")

	err = err.Joins("join target on target.id = target_visum.id_target")
	err = err.Joins("join cms_users on target_visum.id_cms_users = cms_users.id")
	err = err.Joins("join mst_visum_status on  mst_visum_status.id = target_visum.id_mst_visum_status")
	err = err.Joins("join mst_data_source on mst_data_source.id = target.id_mst_data_source")
	err = err.Joins("join mst_branch on mst_branch.id= target.id_mst_branch")

	err = err.Where("target_visum.id_cms_users = ?", id_cms_users)
	err = err.Where("target_visum.created_at::text LIKE ?", "%"+created_at+"%")

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

func TargetVisumDownload(id_mst_outlet string, id_mst_branch []int, created_at, limit string, offset string) structs.JsonResponse {

	var (
		targetvisum []structs.TargetVisumDownload
		t           structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("target_visum").Select("distinct target_visum.id, target_visum.id_target, target_visum.created_at," +
		"target.first_name, target.last_name, target.address as alamat, target.kelurahan, target.kecamatan," +
		"target.kabupaten, target.provinsi,mst_data_source.datasource,target_visum.revisit," +
		"mst_visum_status.status as visum_status,cms_users.name ,cms_users.npm ," +
		"cms_privileges.name as privileges_name, mst_outlet.outlet_name as users_outlet_name," +
		"mst_branch.branch_name as users_branch_name," +
		"target.hp_1 as hp_satu, target.hp_2 as hp_dua," +
		"(select cms_users_cabang.id_cms_users_oh from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_oh," +
		"(select name from cms_users where cms_users.id = id_cms_users_oh )  as oh_name," +
		"(select cms_users_cabang.id_cms_users_spv from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_spv," +
		"(select name from cms_users where cms_users.id = id_cms_users_spv )  as spv_name")

	err = err.Joins("join target on target_visum.id_target = target.id")
	err = err.Joins("left join mst_data_source on mst_data_source.id = target.id_mst_data_source")
	err = err.Joins("left join mst_visum_status on  mst_visum_status.id = target_visum.id_mst_visum_status")
	err = err.Joins("left join cms_users on cms_users.id = target_visum.id_cms_users")
	err = err.Joins("left join cms_privileges on cms_privileges.id = cms_users.id_cms_privileges")
	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id = mst_outlet.id_mst_branch")
	err = err.Joins("left join cms_users_cabang on cms_users_cabang.id_cms_users = cms_users.id")

	//err = err.Where("date (lead.created_at) BETWEEN ? AND ?", created_at1, created_at2)

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
	if created_at != "" {
		err = err.Where("target_visum.created_at::text LIKE ?", "%"+created_at+"%")
	}

	//err = err.Order("target_visum.id desc")

	err = err.Find(&targetvisum)
	errx := err.Error

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	f.SetCellValue("Sheet1", "A1", "Id Target")
	f.SetCellValue("Sheet1", "B1", "Created at")
	f.SetCellValue("Sheet1", "C1", "Datasource")
	f.SetCellValue("Sheet1", "D1", "revisit")
	f.SetCellValue("Sheet1", "E1", "first_name")
	f.SetCellValue("Sheet1", "F1", "last_name")
	f.SetCellValue("Sheet1", "G1", "alamat")
	f.SetCellValue("Sheet1", "H1", "kelurahan")
	f.SetCellValue("Sheet1", "I1", "kecamatan")
	f.SetCellValue("Sheet1", "J1", "kabupaten")
	f.SetCellValue("Sheet1", "K1", "provinsi")
	f.SetCellValue("Sheet1", "L1", "visum_status")
	f.SetCellValue("Sheet1", "M1", "name")
	f.SetCellValue("Sheet1", "N1", "NPM")
	f.SetCellValue("Sheet1", "O1", "privileges_name")
	f.SetCellValue("Sheet1", "P1", "users_outlet_name")
	f.SetCellValue("Sheet1", "Q1", "users_branch_name")
	f.SetCellValue("Sheet1", "R1", "hp_1")
	f.SetCellValue("Sheet1", "S1", "hp_2")
	f.SetCellValue("Sheet1", "T1", "oh_name")
	f.SetCellValue("Sheet1", "U1", "spv_name")

	for i, _ := range targetvisum {

		rows := 2

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+rows), targetvisum[i].IdTarget)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+rows), targetvisum[i].CreatedAt)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+rows), targetvisum[i].Datasource)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+rows), targetvisum[i].Revisit)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+rows), targetvisum[i].FirstName)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+rows), targetvisum[i].LastName)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+rows), targetvisum[i].Alamat)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+rows), targetvisum[i].Kelurahan)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+rows), targetvisum[i].Kecamatan)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(i+rows), targetvisum[i].Kabupaten)
		f.SetCellValue("Sheet1", "K"+strconv.Itoa(i+rows), targetvisum[i].Provinsi)
		f.SetCellValue("Sheet1", "L"+strconv.Itoa(i+rows), targetvisum[i].VisumStatus)
		f.SetCellValue("Sheet1", "M"+strconv.Itoa(i+rows), targetvisum[i].Name)
		f.SetCellValue("Sheet1", "N"+strconv.Itoa(i+rows), targetvisum[i].Npm)
		f.SetCellValue("Sheet1", "O"+strconv.Itoa(i+rows), targetvisum[i].PrivilegesName)
		f.SetCellValue("Sheet1", "P"+strconv.Itoa(i+rows), targetvisum[i].UsersOutletName)
		f.SetCellValue("Sheet1", "Q"+strconv.Itoa(i+rows), targetvisum[i].UsersBranchName)
		f.SetCellValue("Sheet1", "R"+strconv.Itoa(i+rows), targetvisum[i].HpSatu)
		f.SetCellValue("Sheet1", "S"+strconv.Itoa(i+rows), targetvisum[i].HpDua)
		f.SetCellValue("Sheet1", "T"+strconv.Itoa(i+rows), targetvisum[i].OhName)
		f.SetCellValue("Sheet1", "U"+strconv.Itoa(i+rows), targetvisum[i].SpvName)

	}

	f.SetActiveSheet(index)
	folderPath := "./files/"

	filename := "target_visum" + time.Now().Format("20060102150405") + ".xlsx"

	os.MkdirAll(folderPath, os.ModePerm)

	// Save xlsx file by the given path.
	errs := f.SaveAs(folderPath + filename)
	if errs != nil {
		fmt.Println(errs)
	}

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = filename
	}

	return response
}

func TargetVisumChart(id_mst_visum_status string, id_mst_branch []int, created_at string) structs.JsonResponse {

	var (
		targetvisumdetail structs.TargetVisumChart
		t                 structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("target_visum").Select("count(target_visum.id) as target_visum")

	err = err.Joins("join target on target.id = target_visum.id_target")

	if created_at != "" {
		err = err.Where("target_visum.created_at::text like  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		err = err.Where("target.id_mst_branch  IN (?)", id_mst_branch)
	}
	if id_mst_visum_status != "" {
		err = err.Where("target_visum.id_mst_visum_status =  ?", id_mst_visum_status)
	}

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
