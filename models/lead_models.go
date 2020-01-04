package models

import (
	"../structs"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
	"time"
)

func LeadDownload(id_mst_outlet string, id_mst_branch []int, created_at1 string, created_at2 string, limit string, offset string) structs.JsonResponse {

	var (
		lead []structs.DownloadLead
		t    structs.Component
	)

	response := structs.JsonResponse{}
	tables := t.TblLead()

	err := idb.DB.Table(tables).Select("lead.id as id_lead,lead.created_at,lead.favorite,lead.first_name,lead.last_name," +
		"lead.id_cms_users,cms_users.name,cms_users.npm,cms_privileges.name as privileges_name," +
		"mst_outlet.outlet_name as outlet_name , mst_outlet.id_mst_branch , mst_branch.branch_name as branch_name," +
		"mst_job.job , mst_data_source.datasource , lead_mst_status.status as lead_status ," +
		"lead_address.address as alamat, mst_address.kelurahan , mst_address.kecamatan , mst_address.kabupaten," +
		"mst_address.provinsi,mst_address.kodepos,lead_log.duration,mst_log_desc.description," +
		"lead_log.recall,lead_note.note,lead_phone.number,lead_phone.status as phone_status," +
		"mst_product.nama as nama_product , mst_unit.year , mst_unit.merk , mst_unit.type ," +
		"mst_unit.model,mst_unit.otr,lead_product_detail.nopol,lead_product_detail.tax_status," +
		"lead_product_detail.owner")

	err = err.Joins("left join cms_users on lead.id_cms_users = cms_users.id")
	err = err.Joins("left join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")

	err = err.Joins("left join mst_job on mst_job.id = lead.id_mst_job ")
	err = err.Joins("left join mst_data_source on mst_data_source.id = lead.id_mst_data_source")
	err = err.Joins("left join lead_mst_status on lead_mst_status.id = lead.id_lead_mst_status")
	err = err.Joins("left join mst_outlet on mst_outlet.id = lead.id_mst_outlet")
	err = err.Joins("left join lead_address on lead_address.id_lead = lead.id and lead_address.id = (select max(a.id) from lead_address a where a.id_lead = lead_address.id_lead  )")
	err = err.Joins("left join mst_address on lead_address.id_mst_address = mst_address.id")
	err = err.Joins("left join lead_log on lead_log.id_lead = lead.id and lead_log.id = (select max(a.id) from lead_log a where a.id_lead = lead_log.id_lead)")
	err = err.Joins("left join mst_log_desc on mst_log_desc.id = lead_log.id_mst_log_desc")
	err = err.Joins("left join lead_note on lead_note.id_lead = lead.id and lead_note.id = (select max(a.id) from lead_note a where a.id_lead = lead_note.id_lead)")
	err = err.Joins("left join lead_phone on lead_phone.id_lead = lead.id and lead_phone.id = (select max(a.id) from lead_phone a where a.id_lead = lead_phone.id_lead)")
	err = err.Joins("left join lead_product on lead_product.id_lead = lead.id and lead_product.id = (select max(a.id) from lead_product a where a.id_lead = lead_product.id_lead)")
	err = err.Joins("left join mst_product on lead_product.id_mst_product = mst_product.id")
	err = err.Joins("left join lead_product_detail on lead_product_detail.id_lead_product = lead_product.id")
	err = err.Joins("left join mst_unit on lead_product_detail.id_mst_unit = mst_unit.id")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	err = err.Where("date(lead.created_at) BETWEEN ? AND ?", created_at1, created_at2)

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if id_mst_outlet != "" {
		err = err.Where("lead.id_mst_outlet = ?", id_mst_outlet)
	}
	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
	}

	err = err.Find(&lead)
	errx := err.Error

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	f.SetCellValue("Sheet1", "A1", "ID Lead")
	f.SetCellValue("Sheet1", "B1", "NPM")
	f.SetCellValue("Sheet1", "C1", "Name")
	f.SetCellValue("Sheet1", "D1", "OutletName")
	f.SetCellValue("Sheet1", "E1", "BranchName")
	f.SetCellValue("Sheet1", "F1", "Privileges")
	f.SetCellValue("Sheet1", "G1", "Tanggal")
	f.SetCellValue("Sheet1", "H1", "Lead Status")
	f.SetCellValue("Sheet1", "I1", "FirstName")
	f.SetCellValue("Sheet1", "J1", "LastName")
	f.SetCellValue("Sheet1", "K1", "Number")
	f.SetCellValue("Sheet1", "L1", "DataSource")
	f.SetCellValue("Sheet1", "M1", "Alamat")
	f.SetCellValue("Sheet1", "N1", "kelurahan")
	f.SetCellValue("Sheet1", "O1", "kecamatan")
	f.SetCellValue("Sheet1", "P1", "kabupaten")
	f.SetCellValue("Sheet1", "Q1", "provinsi")
	f.SetCellValue("Sheet1", "R1", "Job")
	f.SetCellValue("Sheet1", "S1", "Duration")
	f.SetCellValue("Sheet1", "T1", "Description")
	f.SetCellValue("Sheet1", "U1", "Recall")
	f.SetCellValue("Sheet1", "V1", "Note")
	f.SetCellValue("Sheet1", "W1", "Nama Product")
	f.SetCellValue("Sheet1", "X1", "Year")
	f.SetCellValue("Sheet1", "Y1", "Merk")
	f.SetCellValue("Sheet1", "Z1", "Type")
	f.SetCellValue("Sheet1", "AA1", "Model")
	f.SetCellValue("Sheet1", "AB1", "otr")
	f.SetCellValue("Sheet1", "AC1", "Nopol")
	f.SetCellValue("Sheet1", "AD1", "Status Pajak")
	f.SetCellValue("Sheet1", "AE1", "Owner")

	for i, _ := range lead {

		rows := 2

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+rows), lead[i].IdLead)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+rows), lead[i].Npm)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+rows), lead[i].Name)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+rows), lead[i].OutletName)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+rows), lead[i].BranchName)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+rows), lead[i].PrivilegesName)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(i+rows), lead[i].CreatedAt)
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(i+rows), lead[i].LeadStatus)
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(i+rows), lead[i].FirstName)
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(i+rows), lead[i].LastName)
		f.SetCellValue("Sheet1", "K"+strconv.Itoa(i+rows), lead[i].Number)
		f.SetCellValue("Sheet1", "L"+strconv.Itoa(i+rows), lead[i].Datasource)
		f.SetCellValue("Sheet1", "M"+strconv.Itoa(i+rows), lead[i].Alamat)
		f.SetCellValue("Sheet1", "N"+strconv.Itoa(i+rows), lead[i].Kelurahan)
		f.SetCellValue("Sheet1", "O"+strconv.Itoa(i+rows), lead[i].Kecamatan)
		f.SetCellValue("Sheet1", "P"+strconv.Itoa(i+rows), lead[i].Kabupaten)
		f.SetCellValue("Sheet1", "Q"+strconv.Itoa(i+rows), lead[i].Provinsi)
		f.SetCellValue("Sheet1", "R"+strconv.Itoa(i+rows), lead[i].Job)
		f.SetCellValue("Sheet1", "S"+strconv.Itoa(i+rows), lead[i].Duration)
		f.SetCellValue("Sheet1", "T"+strconv.Itoa(i+rows), lead[i].Description)
		f.SetCellValue("Sheet1", "U"+strconv.Itoa(i+rows), lead[i].Recall)
		f.SetCellValue("Sheet1", "V"+strconv.Itoa(i+rows), lead[i].Note)
		f.SetCellValue("Sheet1", "W"+strconv.Itoa(i+rows), lead[i].NamaProduct)
		f.SetCellValue("Sheet1", "X"+strconv.Itoa(i+rows), lead[i].Year)
		f.SetCellValue("Sheet1", "Y"+strconv.Itoa(i+rows), lead[i].Merk)
		f.SetCellValue("Sheet1", "Z"+strconv.Itoa(i+rows), lead[i].Type)
		f.SetCellValue("Sheet1", "AA"+strconv.Itoa(i+rows), lead[i].Model)
		f.SetCellValue("Sheet1", "AB"+strconv.Itoa(i+rows), lead[i].Otr)
		f.SetCellValue("Sheet1", "AC"+strconv.Itoa(i+rows), lead[i].Nopol)
		f.SetCellValue("Sheet1", "AD"+strconv.Itoa(i+rows), lead[i].TaxStatus)
		f.SetCellValue("Sheet1", "AE"+strconv.Itoa(i+rows), lead[i].Owner)

	}

	f.SetActiveSheet(index)
	folderPath := "./files/"
	filename := "Lead" + time.Now().Format("20060102150405") + ".xlsx"

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

func LeadDetail(id_lead string) structs.JsonResponse {

	var (
		status structs.LeadDetail
		t      structs.Component
	)

	response := structs.JsonResponse{}
	tables := t.TblLead()

	err := idb.DB.Table(tables).Select("distinct (lead.id ),lead.created_at,lead.favorite,lead.first_name,lead.last_name," +
		"lead.id_cms_users,cms_users.name,cms_privileges.name as privileges,mst_outlet.outlet_name as outlet_users," +
		"mst_outlet.id_mst_branch,mst_branch.branch_name,mst_job.job,mst_data_source.datasource," +
		"lead_mst_status.status as lead_status,lead_address.address,mst_address.kelurahan,mst_address.kecamatan," +
		"mst_address.kabupaten,mst_address.provinsi,mst_address.kodepos,lead_log.duration,mst_log_desc.description," +
		"lead_log.recall,lead_note.note,lead_phone.number,lead_phone.status as phone_status," +
		"mst_product.nama as nama_product,mst_unit.year,mst_unit.merk,mst_unit.type,mst_unit.model,mst_unit.otr," +
		"lead_product_detail.nopol,lead_product_detail.tax_status,lead_product_detail.owner")

	err = err.Joins("join cms_users on lead.id_cms_users = cms_users.id")
	err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")

	err = err.Joins("left join mst_job on mst_job.id = lead.id_mst_job")
	err = err.Joins("left join mst_data_source on mst_data_source.id = lead.id_mst_data_source")
	err = err.Joins("left join lead_mst_status on lead_mst_status.id = lead.id_lead_mst_status")
	err = err.Joins("left join mst_outlet on mst_outlet.id = lead.id_mst_outlet")
	err = err.Joins("left join lead_address on lead_address.id_lead = lead.id")
	err = err.Joins("left join mst_address on lead_address.id_mst_address = mst_address.id")
	err = err.Joins("left join lead_log on lead_log.id_lead = lead.id")
	err = err.Joins("left join mst_log_desc on mst_log_desc.id = lead_log.id_mst_log_desc")
	err = err.Joins("left join lead_note on lead_note.id_lead = lead.id")
	err = err.Joins("left join lead_phone on lead_phone.id_lead = lead.id")
	err = err.Joins("left join lead_product on lead_product.id_lead = lead.id")
	err = err.Joins("left join mst_product on lead_product.id_mst_product = mst_product.id")
	err = err.Joins("left join lead_product_detail on lead_product_detail.id_lead_product = lead_product.id")
	err = err.Joins("left join mst_unit on lead_product_detail.id_mst_unit = mst_unit.id")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	err = err.Where("lead.id = ?", id_lead)

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

func LeadUsers(id_mst_outlet string, id_mst_branch []int, created_at, limit string, offset string) structs.JsonResponse {

	var (
		status []structs.LeadUsers
		t      structs.Component
	)

	response := structs.JsonResponse{}
	//tables := t.TblLead()

	date := "%" + created_at + "%"

	err := idb.DB.Table("cms_users").Select("distinct (cms_users.id), cms_users.npm, cms_users.name , mst_outlet.outlet_name,"+
		"cms_users.status,cms_privileges.name as privileges_name, mst_branch.branch_name,"+
		"(select count(*) from lead where lead.id_cms_users = cms_users.id and lead.created_at::text LIKE ? ) as lead,"+
		"(select count(*) from lead where lead.id_cms_users = cms_users.id and lead.id_lead_mst_status = 1  and lead.created_at::text LIKE ? ) as new,"+
		"(select count(*) from lead where lead.id_cms_users = cms_users.id and lead.id_lead_mst_status = 2  and lead.created_at::text LIKE ? ) as hot,"+
		"(select count(*) from lead where lead.id_cms_users = cms_users.id and lead.id_lead_mst_status = 3  and lead.created_at::text LIKE ? ) as working,"+
		"(select count(*) from lead where lead.id_cms_users = cms_users.id and lead.id_lead_mst_status = 4  and lead.created_at::text LIKE ? ) as unqualified,"+
		"(select count(*) from lead where lead.id_cms_users = cms_users.id and lead.id_lead_mst_status = 5  and lead.created_at::text LIKE ? ) as converted,"+
		"(select count(*) from lead where lead.id_cms_users = cms_users.id and lead.id_lead_mst_status = 6  and lead.created_at::text LIKE ? ) as deleted", date, date, date, date, date, date, date)

	err = err.Joins("join cms_privileges on cms_users.id_cms_privileges = cms_privileges.id")
	err = err.Joins("left join mst_outlet on mst_outlet.id = cms_users.id_mst_outlet")
	err = err.Joins("left join mst_branch on mst_branch.id= mst_outlet.id_mst_branch")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if id_mst_outlet != "" {
		err = err.Where("cms_users.id_mst_outlet = ?", id_mst_outlet)
	}
	fmt.Println(id_mst_branch, len(id_mst_branch), cap(id_mst_branch))
	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch IN (?)", id_mst_branch)
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
		response.ApiMessage = t.GetMessageSucc()
		response.Data = nil

	} else {

		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = status

	}

	return response
}

func GetLeadChart(id_mst_branch []int, created_at string) structs.JsonResponse {

	var (
		orderstatus structs.GetLeadChart
		t           structs.Component
	)

	response := structs.JsonResponse{}

	err := idb.DB.Table("lead").Select("count(lead.id) as leads")

	err = err.Joins("join mst_outlet on mst_outlet.id = lead.id_mst_outlet")

	if created_at != "" {
		err = err.Where("lead.created_at::text like  ?", "%"+created_at+"%")
	}
	if len(id_mst_branch) >= 1 {
		err = err.Where("mst_outlet.id_mst_branch  IN (?)", id_mst_branch)
	}

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
