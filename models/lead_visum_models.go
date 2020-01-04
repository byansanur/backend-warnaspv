package models

import (
	"../structs"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
	"time"
)

func LeadVisumUsers(id_mst_outlet string, id_mst_branch []int, created_at, limit string, offset string) structs.JsonResponse {

	var (
		leadvisum []structs.LeadVisumUsers
		t         structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	date := "%" + created_at + "%"

	err := idb.DB.Table("cms_users").Select("distinct (cms_users.id), cms_users.npm, cms_users.name , mst_outlet.outlet_name,"+
		"cms_users.status,cms_privileges.name as privileges_name, mst_branch.branch_name,"+
		"(select count(id) from lead_visum where lead_visum.id_cms_users = cms_users.id and lead_visum.created_at::text LIKE  ? ) as visum,"+
		"(select count(id) from lead_visum where lead_visum.id_cms_users = cms_users.id and lead_visum.id_mst_visum_status = 1 and lead_visum.created_at::text LIKE ? ) as rumah_kosong,"+
		"(select count(id) from lead_visum where lead_visum.id_cms_users = cms_users.id and lead_visum.id_mst_visum_status = 2 and lead_visum.created_at::text LIKE ? ) as alamat_tidak_sesuai,"+
		"(select count(id) from lead_visum where lead_visum.id_cms_users = cms_users.id and lead_visum.id_mst_visum_status = 3 and lead_visum.created_at::text LIKE ? ) as ditolak_atau_diusir,"+
		"(select count(id) from lead_visum where lead_visum.id_cms_users = cms_users.id and lead_visum.id_mst_visum_status = 4 and lead_visum.created_at::text LIKE ? ) as butuh_waktu,"+
		"(select count(id) from lead_visum where lead_visum.id_cms_users = cms_users.id and lead_visum.id_mst_visum_status = 5 and lead_visum.created_at::text LIKE ? ) as tanya_pasangan,"+
		"(select count(id) from lead_visum where lead_visum.id_cms_users = cms_users.id and lead_visum.id_mst_visum_status = 6 and lead_visum.created_at::text LIKE ? ) as pikir_pikir,"+
		"(select count(id) from lead_visum where lead_visum.id_cms_users = cms_users.id and lead_visum.id_mst_visum_status = 7 and lead_visum.created_at::text LIKE ? ) as berminat,"+
		"(select count(id) from lead_visum where lead_visum.id_cms_users = cms_users.id and lead_visum.id_mst_visum_status = 8 and lead_visum.created_at::text LIKE ? ) as rumah_tidak_ketemu,"+
		"(select count(id) from lead_visum where lead_visum.id_cms_users = cms_users.id and lead_visum.id_mst_visum_status = 9 and lead_visum.created_at::text LIKE ? ) as tidak_berminat", date, date, date, date, date, date, date, date, date, date)

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

	err = err.Find(&leadvisum)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	}

	if len(leadvisum) <= 0 {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = nil

	} else {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = leadvisum

	}

	return response
}

func LeadVisumDetail(id_cms_users string, created_at string) structs.JsonResponse {

	var (
		leadvisumdetail []structs.LeadVisumDetail
		t               structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("lead_visum").Select("distinct (lead_visum.id) , lead_visum.created_at," +
		"mst_data_source.datasource, lead_visum.revisit, lead.first_name, lead.last_name," +
		"lead_address.address as alamat,mst_address.kabupaten,mst_address.kecamatan,mst_address.kelurahan," +
		"mst_address.provinsi,mst_address.kodepos, mst_visum_status.status,cms_users.name")

	err = err.Joins("join lead on lead.id = lead_visum.id_lead")
	err = err.Joins("join cms_users on lead_visum.id_cms_users = cms_users.id")
	err = err.Joins("left join mst_visum_status on  mst_visum_status.id = lead_visum.id_mst_visum_status")
	err = err.Joins("left join mst_data_source on mst_data_source.id = lead.id_mst_data_source")
	err = err.Joins("left join mst_outlet on mst_outlet.id = lead.id_mst_outlet")
	err = err.Joins("left join lead_address on lead_address.id_lead = lead.id")
	err = err.Joins("left join mst_address on lead_address.id_mst_address = mst_address.id")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	err = err.Where("lead_visum.id_cms_users = ?", id_cms_users)
	err = err.Where("lead_visum.created_at::text LIKE ?", "%"+created_at+"%")

	err = err.Find(&leadvisumdetail)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = leadvisumdetail
	}

	return response
}

func LeadVisumDownload(id_mst_outlet string, id_mst_branch []int, created_at string, limit string, offset string) structs.JsonResponse {

	var (
		lead []structs.LeadVisumDownload
		t    structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("lead_visum").Select("distinct (lead_visum.id),lead_visum.id_lead,lead_visum.created_at,mst_data_source.datasource," +
		"lead_visum.revisit , lead.first_name ,lead.last_name ,lead_address.address as alamat, mst_address.kabupaten ," +
		"mst_address.kecamatan , mst_address.kelurahan , mst_address.provinsi ,mst_address.kodepos ," +
		"mst_visum_status.status as visum_status, cms_users.name , cms_users.npm , mst_outlet.outlet_name as lead_outlet_name," +
		"mst_branch.branch_name as lead_branch_name , cms_privileges.name as privileges_name , lead_phone.number ," +
		"cms_users_cabang.id_cms_users_oh as id_cms_users_oh ,cms_users_cabang.id_cms_users_spv as id_cms_users_spv ," +
		"(select cms_users_cabang.id_cms_users_oh from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_oh," +
		"(select name from cms_users where cms_users.id = id_cms_users_oh )  as oh_name," +
		"(select cms_users_cabang.id_cms_users_spv from cms_users_cabang where cms_users_cabang.id_cms_users = cms_users.id limit 1) as id_cms_users_spv," +
		"(select name from cms_users where cms_users.id = id_cms_users_spv )  as spv_name")

	err = err.Joins(" left join  lead  on  lead.id  =   lead_visum.id_lead")
	err = err.Joins(" left join  mst_visum_status  on   mst_visum_status.id  =  lead_visum.id_mst_visum_status")
	err = err.Joins(" left join  lead_address  on   lead_address.id_lead  =   lead.id")
	err = err.Joins("left join  mst_address  on   mst_address.id   = lead_address.id_mst_address")
	err = err.Joins("left join  mst_data_source  on   mst_data_source.id  =   lead.id_mst_data_source")
	err = err.Joins(" left join  cms_users  on   cms_users.id  =   lead.id_cms_users")
	err = err.Joins("left join  cms_privileges  on   cms_privileges.id  =   cms_users.id_cms_privileges")
	err = err.Joins("left join  mst_outlet  on   mst_outlet.id  =   lead.id_mst_outlet")
	err = err.Joins("left join  lead_phone  on  lead_phone.id_lead  =  lead.id")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")
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
		err = err.Where("lead.id_mst_outlet = ?", id_mst_outlet)
	}
	if created_at != "" {
		err = err.Where("lead_visum.created_at::text LIKE ?", "%"+created_at+"%")
	}

	err = err.Find(&lead)
	errx := err.Error

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	f.SetCellValue("Sheet1", "A1", "ID Lead")
	f.SetCellValue("Sheet1", "B1", "Tanggal")
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
	f.SetCellValue("Sheet1", "P1", "lead_outlet_name")
	f.SetCellValue("Sheet1", "Q1", "lead_branch_name")
	f.SetCellValue("Sheet1", "R1", "number")
	f.SetCellValue("Sheet1", "S1", "oh_name")
	f.SetCellValue("Sheet1", "T1", "spv_name")

	for i, _ := range lead {

		rows := 2

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+rows), lead[i].IdLead)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+rows), lead[i].CreatedAt)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+rows), lead[i].Datasource)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+rows), lead[i].Revisit)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+rows), lead[i].FirstName)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+rows), lead[i].LastName)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+rows), lead[i].Alamat)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+rows), lead[i].Kelurahan)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+rows), lead[i].Kecamatan)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(i+rows), lead[i].Kabupaten)
		f.SetCellValue("Sheet1", "K"+strconv.Itoa(i+rows), lead[i].Provinsi)
		f.SetCellValue("Sheet1", "L"+strconv.Itoa(i+rows), lead[i].VisumStatus)
		f.SetCellValue("Sheet1", "M"+strconv.Itoa(i+rows), lead[i].Name)
		f.SetCellValue("Sheet1", "N"+strconv.Itoa(i+rows), lead[i].Npm)
		f.SetCellValue("Sheet1", "O"+strconv.Itoa(i+rows), lead[i].PrivilegesName)
		f.SetCellValue("Sheet1", "P"+strconv.Itoa(i+rows), lead[i].LeadOutletName)
		f.SetCellValue("Sheet1", "Q"+strconv.Itoa(i+rows), lead[i].LeadBranchName)
		f.SetCellValue("Sheet1", "R"+strconv.Itoa(i+rows), lead[i].Number)
		f.SetCellValue("Sheet1", "S"+strconv.Itoa(i+rows), lead[i].OhName)
		f.SetCellValue("Sheet1", "T"+strconv.Itoa(i+rows), lead[i].SpvName)

	}

	f.SetActiveSheet(index)
	folderPath := "./files/"

	filename := "lead_visum" + time.Now().Format("20060102150405") + ".xlsx"

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

	}

	if len(lead) <= 0 {

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

func LeadVisumChart(id_mst_visum_status string, id_mst_branch []int, created_at string) structs.JsonResponse {

	var (
		targetvisumdetail structs.LeadVisumChart
		t                 structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	err := idb.DB.Table("lead_visum").Select("count(lead_visum.id) as lead_visum")

	err = err.Joins("join lead on lead.id = lead_visum.id_lead")
	err = err.Joins("join mst_outlet on lead.id_mst_outlet = mst_outlet.id")

	if created_at != "" {
		err = err.Where("lead_visum.created_at::text like  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch IN  (?)", id_mst_branch)
	}
	if id_mst_visum_status != "" {
		err = err.Where("lead_visum.id_mst_visum_status =  ?", id_mst_visum_status)
	}
	fmt.Println("Created_at lea visum ", created_at)

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
